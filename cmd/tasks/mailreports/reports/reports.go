package reports

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"time"

	"github.com/jackc/pgx"
	"github.com/linkai-io/am/am"
	"github.com/linkai-io/am/pkg/initializers"
	"github.com/linkai-io/am/pkg/mail"
	"github.com/rs/zerolog/log"
)

const (
	dailySubject  = "Your daily hakken service email report"
	weeklySubject = "Your weekly hakken service email report"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("email").Funcs(template.FuncMap{
		"event": EventPrinter,
		"time":  TimePrinter,
	}).Parse(reportHTMLTemplate))
}

type Report struct {
	OrgID            int
	UserID           int
	UserTimeZone     string
	ReportType       string
	OrganizationName string
	FirstName        string
	LastName         string
	Email            string
	GroupReports     map[string]map[int32][]*ScanGroupReport // sg name (key), map (type_id) key
	Since            time.Time
	Now              time.Time
}

type ScanGroupReport struct {
	Timestamp time.Time
	Data      []string
}

func GenerateReports(appConfig initializers.AppConfig, weekly bool, mailer mail.Mailer) {
	since := time.Now().Add(time.Hour * (-24))

	_, pool := initializers.DB(&appConfig)

	query := `select org.organization_id, org.organization_name, u.user_id, u.first_name, u.last_name, u.email from am.users as u 
		join am.organizations as org on org.organization_id=u.organization_id 
		where org.deleted=false and u.deleted=false	and status_id=1000 and subscription_id<>9999`

	rows, err := pool.Query(query)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get list of users for reporting")
	}
	defer rows.Close()

	if weekly {
		since = time.Now().Add(time.Hour * (-24 * 7))
	}

	for rows.Next() {
		report := &Report{ReportType: "daily"}
		report.Since = since
		if weekly {
			report.ReportType = "weekly"
		}
		if err := rows.Scan(&report.OrgID, &report.OrganizationName, &report.UserID, &report.FirstName, &report.LastName, &report.Email); err != nil {
			log.Warn().Err(err).Msg("unable to get oid/uid, continuing")
			continue
		}

		query, args, err := buildGetReportQuery(report.OrgID, report.UserID, weekly, report.Since)
		if err != nil {
			log.Fatal().Err(err).Int("OrgID", report.OrgID).Int("UserID", report.UserID).Str("query", query).Msg("something wrong with report query")
		}

		report.GroupReports = make(map[string]map[int32][]*ScanGroupReport)
		eventRows, err := pool.Query(query, args...)
		if err != nil {
			log.Warn().Err(err).Int("OrgID", report.OrgID).Int("UserID", report.UserID).Str("query", query).Msg("failed to query events")
			continue
		}
		if err := GetReportEvents(mailer, pool, eventRows, report); err != nil {
			log.Warn().Err(err).Int("OrgID", report.OrgID).Int("UserID", report.UserID).Msg("error generating/sending report")
		}
	}
}

func GetReportEvents(mailer mail.Mailer, pool *pgx.ConnPool, rows *pgx.Rows, report *Report) error {
	defer rows.Close()

	for rows.Next() {
		var orgID int
		var userID int
		var sgName string
		var typeID int32
		var data []string
		var ts time.Time
		var shouldEmailWeekly bool
		var shouldEmailDaily bool

		if err := rows.Scan(&orgID, &sgName, &userID, &typeID, &ts, &data, &report.UserTimeZone, &shouldEmailWeekly, &shouldEmailDaily); err != nil {
			log.Warn().Err(err).Msg("failed to read event data")
			continue
		}

		if report.UserID != userID {
			log.Error().Msg("user id did not match")
			return errors.New("user id did not match from events")
		}

		if report.OrgID != orgID {
			log.Error().Msg("org id did not match")
			return errors.New("org id did not match from events")
		}

		if report.ReportType == "weekly" && shouldEmailWeekly == false {
			log.Info().Str("Org", report.OrganizationName).Int("User", report.UserID).Msg("user not subscribed for weekly emails")
			return nil
		}

		if report.ReportType == "daily" && shouldEmailDaily == false {
			log.Info().Str("Org", report.OrganizationName).Int("User", report.UserID).Msg("user not subscribed for daily emails")
			return nil
		}

		if _, ok := report.GroupReports[sgName]; !ok {
			report.GroupReports[sgName] = make(map[int32][]*ScanGroupReport)
		}

		if _, ok := report.GroupReports[sgName][typeID]; !ok {
			report.GroupReports[sgName][typeID] = make([]*ScanGroupReport, 0)
		}

		// for certificates, there's the chance it already expired, so we should just grab it directly from DB instead of events
		if typeID == am.EventCertExpiring {
			evt, err := CheckCertificates(pool, sgName, report)
			if err != nil {
				log.Warn().Err(err).Msg("failed to get certificates")
				continue
			}

			if evt == nil || len(evt) == 0 {
				continue
			}
			report.GroupReports[sgName][typeID] = append(report.GroupReports[sgName][typeID], &ScanGroupReport{Timestamp: time.Now(), Data: evt})
			continue
		}
		report.GroupReports[sgName][typeID] = append(report.GroupReports[sgName][typeID], &ScanGroupReport{Timestamp: ts, Data: data})
	}
	report.Now = time.Now()
	return SendReport(mailer, report)
}

func CheckCertificates(pool *pgx.ConnPool, sgName string, report *Report) ([]string, error) {
	rows, err := pool.Query(`select subject_name, port, valid_to from am.web_certificates 
		where (TIMESTAMPTZ 'epoch' + valid_to * '1 second'::interval) 
		between now() and now() + interval '30 days'
		and organization_id=$1
		and scan_group_id=(select scan_group_id from am.scan_group where scan_group_name=$2)`, report.OrgID, sgName)
	if err != nil {
		return nil, err
	}
	certs := make([]string, 0)
	for i := 0; rows.Next(); i++ {
		var subjectName string
		var port int
		var validTo int64
		if err := rows.Scan(&subjectName, &port, &validTo); err != nil {
			log.Error().Err(err).Msg("failed to scan certificate expired event")
			continue
		}
		validTime := time.Unix(validTo, 0)
		ts := validTime.Sub(time.Now()).Hours()
		expires := ""
		if ts <= float64(24) {
			expires = fmt.Sprintf("%.01f hours", ts)
		} else {
			expires = fmt.Sprintf("%.0f days", ts/float64(24))
		}

		certs = append(certs, subjectName)
		certs = append(certs, fmt.Sprintf("%d", port))
		certs = append(certs, expires)
	}
	// no new certs this round
	if len(certs) == 0 {
		return nil, nil
	}

	return certs, nil
}

func SendReport(mailer mail.Mailer, report *Report) error {
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, report); err != nil {
		return err
	}
	subject := dailySubject
	if report.ReportType == "weekly" {
		subject = weeklySubject
	}
	log.Info().Str("organization_name", report.OrganizationName).Msg("Sending report")
	return mailer.SendMail(subject, report.Email, buf.String(), buf.String())
}
