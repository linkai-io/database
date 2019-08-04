package reports

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/linkai-io/am/am"
	"github.com/rs/zerolog/log"
)

const reportHTMLTemplate = `<p> Hello {{ .FirstName }} {{ .LastName}},</p>
{{ if .GroupReports.GroupData }}
<p>The following are a collection of events between {{ time .Since .UserTimeZone }} - {{ time .Now .UserTimeZone}}</p>
<br>
  {{ range $groupID, $groups := .GroupReports.GroupData }} 
  <h3>Group: {{ $groups.Name }}</h3>
		{{ range $type_id, $events := $groups.Events}}
		{{ event $type_id $events}}
		{{ end }}
  {{ end }}
{{ else }}
<p>You have no new events during the time period between {{ time .Since .UserTimeZone}} - {{ time .Now .UserTimeZone}}</p>
{{ end }}
<p>If you would no longer like to recieve these emails, please login to the <a href="https://console.linkai.io/login/">hakken console</a> and modify your mail preferences.</p>
`

func TimePrinter(t time.Time, timeZone string) string {
	userTime, err := time.LoadLocation(timeZone)
	if err == nil {
		log.Info().Str("location", timeZone).Msg("setting location")
		t = t.In(userTime)
	}
	return t.Format(time.UnixDate)
}

func EventPrinter(typeID int32, events []*ScanGroupReportEvent) string {
	var buf bytes.Buffer

	switch typeID {
	case am.EventCertExpiredID:
	case am.EventCertExpiringID:
		buf.WriteString("The following certificates will expire soon:")
		buf.WriteString("<br>\n<ul>")
		for _, e := range events {
			if e.JSONData != "" {
				// handle new json type
				var expireCerts []*am.EventCertExpiring
				if err := json.Unmarshal([]byte(e.JSONData), &expireCerts); err != nil {
					continue
				}
				for _, expired := range expireCerts {
					event := fmt.Sprintf("\n\t<li>%s on port %d expires in %s</li>\n", expired.SubjectName, expired.Port, expired.TimeRemaining)
					buf.WriteString(event)
				}
			}
			if e.Data != nil && len(e.Data) > 3 {
				event := fmt.Sprintf("\n\t<li>%s on port %s expires in %s</li>\n", e.Data[0], e.Data[1], e.Data[2])
				buf.WriteString(event)
			}

		}
		buf.WriteString("</ul>")
	case am.EventNewOpenPortID:
		buf.WriteString("The following ports were opened:")
		buf.WriteString("<br>\n<ul>")
		for _, e := range events {
			// handle new json type
			if e.JSONData != "" {

				var openPorts []*am.EventNewOpenPort
				if err := json.Unmarshal([]byte(e.JSONData), &openPorts); err != nil {
					continue
				}
				for _, open := range openPorts {
					ips := open.CurrentIP
					if open.CurrentIP != open.PreviousIP {
						ips += ") previously (" + open.PreviousIP
					}
					buf.WriteString(fmt.Sprintf("\n\t<li>Host %s (%s) ports: %s</li>\n", open.Host, ips, IntToString(open.OpenPorts)))
				}
			}

			if e.Data != nil {
				if len(e.Data)%4 != 0 {
					continue
				}
				for i := 0; i < len(e.Data); i += 4 {
					ips := e.Data[i+1] + ") previously (" + e.Data[i+2]
					if e.Data[i+1] == e.Data[i+2] {
						ips = e.Data[i+1]
					}
					buf.WriteString(fmt.Sprintf("\n\t<li>Host %s (%s) ports: %s</li>\n", e.Data[i], ips, e.Data[i+3]))
				}
			}

		}
		buf.WriteString("</ul>")
	case am.EventClosedPortID:
		buf.WriteString("The following ports were recently closed:")
		buf.WriteString("<br>\n<ul>")

		for _, e := range events {
			if e.JSONData != "" {
				var closedPorts []*am.EventClosedPort
				if err := json.Unmarshal([]byte(e.JSONData), &closedPorts); err != nil {
					continue
				}
				for _, closed := range closedPorts {
					ips := closed.CurrentIP
					if closed.CurrentIP != closed.PreviousIP {
						ips += ") previously (" + closed.PreviousIP
					}
					buf.WriteString(fmt.Sprintf("\n\t<li>Host %s (%s) ports: %s</li>\n", closed.Host, ips, IntToString(closed.ClosedPorts)))
				}
			}

			if e.Data != nil {
				if len(e.Data)%4 != 0 {
					continue
				}
				for i := 0; i < len(e.Data); i += 4 {
					ips := e.Data[i+1] + ") previously (" + e.Data[i+2]
					if e.Data[i+1] == e.Data[i+2] {
						ips = e.Data[i+1]
					}
					buf.WriteString(fmt.Sprintf("\n\t<li>Host %s (%s) ports: %s</li>\n", e.Data[i], ips, e.Data[i+3]))
				}
			}

		}
		buf.WriteString("</ul>")
	case am.EventInitialGroupCompleteID:
	case am.EventMaxHostPricingID:
	case am.EventNewHostID:
		buf.WriteString("The following new hosts were found:")
		buf.WriteString("\n<ul>")
		for _, e := range events {

			if e.JSONData != "" {
				var newHosts []*am.EventNewHost
				if err := json.Unmarshal([]byte(e.JSONData), &newHosts); err != nil {
					continue
				}
				for _, newHost := range newHosts {
					buf.WriteString("\n\t<li>" + newHost.Host + "</li>\n")
				}
			}

			if e.Data != nil {

				for _, host := range e.Data {
					buf.WriteString("\n\t<li>" + host + "</li>\n")
				}
			}

		}
		buf.WriteString("</ul>")
	case am.EventAXFRID:
		buf.WriteString("The following name servers were found leaking hostnames via Zone Transfer")
		buf.WriteString("<br>\n<ul>")
		for _, e := range events {
			if e.JSONData != "" {
				var axfrServers []*am.EventAXFR
				if err := json.Unmarshal([]byte(e.JSONData), &axfrServers); err != nil {
					continue
				}
				for _, axfr := range axfrServers {
					buf.WriteString("\n\t<li>" + strings.Join(axfr.Servers, ",") + "</li>\n")
				}
			}
			if e.Data != nil {
				for _, nsServer := range e.Data {
					buf.WriteString("\n\t<li>" + nsServer + "</li>\n")
				}
			}
		}
		buf.WriteString("</ul>")
	case am.EventNSECID:
		buf.WriteString("The following name servers are leaking hostnames via NSEC records")
		buf.WriteString("<br>\n<ul>")
		for _, e := range events {
			if e.JSONData != "" {
				var nsecServers []*am.EventNSEC
				if err := json.Unmarshal([]byte(e.JSONData), &nsecServers); err != nil {
					continue
				}
				for _, nsec := range nsecServers {
					buf.WriteString("\n\t<li>" + strings.Join(nsec.Servers, ",") + "</li>\n")
				}
			}

			if e.Data != nil {
				for _, nsServer := range e.Data {
					buf.WriteString("\n\t<li>" + nsServer + "</li>\n")
				}
			}
		}
		buf.WriteString("</ul>")
	case am.EventNewWebsiteID:
		buf.WriteString("The following new web sites were found:")
		buf.WriteString("<br>\n<ul>")
		for _, e := range events {
			if e.JSONData != "" {
				var newSites []*am.EventNewWebsite
				if err := json.Unmarshal([]byte(e.JSONData), &newSites); err != nil {
					continue
				}
				for _, site := range newSites {
					if wasRedirected(site.LoadURL, site.URL) {
						buf.WriteString(fmt.Sprintf("\n\t<li><a href=\"%s\">%s</a> (was redirected to <a href=\"%s\">%s</a>) on port %d</li>\n", site.LoadURL, site.LoadURL, site.URL, site.URL, site.Port))
					} else {
						buf.WriteString(fmt.Sprintf("\n\t<li><a href=\"%s\">%s</a> on port %d</li>\n", site.LoadURL, site.LoadURL, site.Port))
					}
				}
			}

			if e.Data != nil {
				for i := 0; i < len(e.Data); i += 2 {
					buf.WriteString(fmt.Sprintf("\n\t<li><a href=\"%s\">%s</a> on port %s</li>\n", e.Data[i], e.Data[i], e.Data[i+1]))
				}
			}
		}
		buf.WriteString("</ul>")
	case am.EventWebHTMLUpdatedID:
	case am.EventWebJSChangedID:
	case am.EventNewWebTechID:
		buf.WriteString("The following new or updated technologies were found:")
		buf.WriteString("<br>\n<ul>")
		for _, e := range events {
			if e.JSONData != "" {
				var newTech []*am.EventNewWebTech
				if err := json.Unmarshal([]byte(e.JSONData), &newTech); err != nil {
					continue
				}
				for _, tech := range newTech {
					if wasRedirected(tech.LoadURL, tech.URL) {
						buf.WriteString(fmt.Sprintf("\n\t<li><a href=\"%s\">%s</a> (was redirected to <a href=\"%s\">%s</a>) is running %s %s</li>\n", tech.LoadURL, tech.LoadURL, tech.URL, tech.URL, tech.TechName, tech.Version))
					} else {
						buf.WriteString(fmt.Sprintf("\n\t<li><a href=\"%s\">%s</a> is running %s %s</li>\n", tech.LoadURL, tech.LoadURL, tech.TechName, tech.Version))
					}
				}
			}

			if e.Data != nil {
				for i := 0; i < len(e.Data); i += 4 {
					buf.WriteString(fmt.Sprintf("\n\t<li><a href=\"%s\">%s</a> on port %s is running %s %s</li>\n", e.Data[i], e.Data[i], e.Data[i+1], e.Data[i+2], e.Data[i+3]))
				}
			}

		}
		buf.WriteString("</ul>")
	}
	return buf.String()
}

func IntToString(in []int32) string {
	b := make([]string, len(in))
	for i, v := range in {
		b[i] = strconv.FormatInt(int64(v), 10)
	}

	return strings.Join(b, ",")
}

// wasRedirected adds a '/' to the end of the load url and returns true if the load url got redirected to a different url, false otherwise
func wasRedirected(loadURL, url string) bool {
	load := loadURL

	if !strings.HasSuffix(loadURL, "/") {
		load = loadURL + "/"
	}

	return strings.Compare(load, url) != 0
}
