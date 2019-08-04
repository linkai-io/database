package reports

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"text/template"
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
	OrgID             int
	UserID            int
	UserTimeZone      string
	ReportType        string
	OrganizationName  string
	FirstName         string
	LastName          string
	Email             string
	ShouldWeeklyEmail bool
	ShouldDailyEmail  bool
	GroupReports      *ScanGroupReport // sg name (key), map (type_id) key
	Since             time.Time
	Now               time.Time
}

func GenerateReports(appConfig initializers.AppConfig, weekly bool, mailer mail.Mailer) {
	since := time.Now().Add(time.Hour * (-24))

	_, pool := initializers.DB(&appConfig)

	activeQuery := `select sga.organization_id,sga.scan_group_id from am.scan_group as sg 
		join am.scan_group_addresses as sga on sga.scan_group_id=sg.scan_group_id 
		where sga.deleted=false group by sga.organization_id,sga.scan_group_id`

	activeGroups := make(map[int][]int)
	activeRows, err := pool.Query(activeQuery)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get list of active groups")
	}
	defer activeRows.Close()

	for activeRows.Next() {
		var orgID int
		var groupID int
		if err := activeRows.Scan(&orgID, &groupID); err != nil {
			log.Fatal().Err(err).Msg("failed to read group")
		}
		if _, ok := activeGroups[orgID]; !ok {
			activeGroups[orgID] = make([]int, 0)
		}
		activeGroups[orgID] = append(activeGroups[orgID], groupID)
	}

	// only return orgs/users that are actually subscribed to receive emails
	query := `select org.organization_id, org.organization_name, u.user_id, u.first_name, u.last_name, u.email, uns.user_timezone, uns.should_weekly_email, uns.should_daily_email from am.users as u 
	join am.organizations as org on org.organization_id=u.organization_id 
	join am.user_notification_settings as uns on u.user_id=uns.user_id where org.deleted=false and u.deleted=false and status_id=1000 and subscription_id<>9999`

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

		if err := rows.Scan(&report.OrgID, &report.OrganizationName, &report.UserID,
			&report.FirstName, &report.LastName, &report.Email, &report.UserTimeZone,
			&report.ShouldWeeklyEmail, &report.ShouldDailyEmail); err != nil {
			log.Warn().Err(err).Msg("unable to get oid/uid, continuing")
			continue
		}

		if _, ok := activeGroups[report.OrgID]; !ok {
			log.Info().Int("OrgID", report.OrgID).Msg("this org has no groups, skipping")
			continue
		}

		if report.ReportType == "weekly" && report.ShouldWeeklyEmail == false {
			log.Info().Str("Org", report.OrganizationName).Int("User", report.UserID).Msg("user not subscribed for weekly emails")
			continue
		}

		if report.ReportType == "daily" && report.ShouldDailyEmail == false {
			log.Info().Str("Org", report.OrganizationName).Int("User", report.UserID).Msg("user not subscribed for daily emails")
			continue
		}

		query, args, err := buildGetReportQuery(report.OrgID, report.UserID, report.Since)
		if err != nil {
			log.Fatal().Err(err).Int("OrgID", report.OrgID).Int("UserID", report.UserID).Str("query", query).Msg("something wrong with report query")
		}
		log.Info().Int("OrgID", report.OrgID).Int("UserID", report.UserID).Str("query", query).Msgf("query args %#v", args)

		report.GroupReports = NewScanGroupReport()
		eventRows, err := pool.Query(query, args...)
		if err != nil {
			log.Warn().Err(err).Int("OrgID", report.OrgID).Int("UserID", report.UserID).Str("query", query).Msg("failed to query events")
			continue
		}

		if err = GetReportEvents(mailer, pool, eventRows, report); err != nil {
			log.Warn().Err(err).Int("OrgID", report.OrgID).Int("UserID", report.UserID).Msg("error generating report")
			continue
		}

		report.Now = time.Now()
		SendReport(mailer, report)
	}
}

func GetReportEvents(mailer mail.Mailer, pool *pgx.ConnPool, rows *pgx.Rows, report *Report) error {
	defer rows.Close()

	for rows.Next() {
		var orgID int
		var userID int
		var groupID int
		var sgName string
		var typeID int32
		var data []string
		var jsonData string
		var ts time.Time

		if err := rows.Scan(&orgID, &groupID, &sgName, &userID, &typeID, &ts, &data, &jsonData); err != nil {
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

		// for certificates, there's the chance it already expired, so skip this and we'll check all after
		if typeID == am.EventCertExpiringID {
			continue
		}
		report.GroupReports.Add(groupID, sgName, typeID, &ScanGroupReportEvent{Timestamp: ts, Data: data, JSONData: jsonData})
	}

	if err := CheckCertificates(pool, report); err != nil {
		log.Warn().Err(err).Msg("failed to get certificates")
	}
	return nil
}

func CheckCertificates(pool *pgx.ConnPool, report *Report) error {
	for groupID, groupReport := range report.GroupReports.GroupData {
		// our list of scan_group_id's for this org is already filtered on 'non-deleted' groups.
		rows, err := pool.Query(`select subject_name, port, valid_to from am.web_certificates 
		where (TIMESTAMPTZ 'epoch' + valid_to * '1 second'::interval) 
		between now() and now() + interval '30 days'
		and organization_id=$1 and scan_group_id=$2
		`, report.OrgID, groupID)
		if err != nil {
			return err
		}
		certs := make([]*am.EventCertExpiring, 0)
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
			certs = append(certs, &am.EventCertExpiring{
				SubjectName:   subjectName,
				Port:          port,
				ValidTo:       validTo,
				TimeRemaining: expires,
			})
		}
		// no new certs this round
		if len(certs) == 0 {
			return nil
		}
		m, err := json.Marshal(certs)
		if err != nil {
			return err
		}
		groupReport.Add(am.EventCertExpiringID, &ScanGroupReportEvent{Timestamp: time.Now(), JSONData: string(m), Data: nil})
	}

	return nil
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
	log.Info().Str("organization_name", report.OrganizationName).Int("UserID", report.UserID).Msg("Sending report")
	return mailer.SendMail(subject, report.Email, buf.String(), buf.String())
}
