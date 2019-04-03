package reports

import (
	"bytes"
	"fmt"
	"time"

	"github.com/linkai-io/am/am"
	"github.com/rs/zerolog/log"
)

const reportHTMLTemplate = `<p> Hello {{ .FirstName }} {{ .LastName}},</p>
{{ if .GroupReports }}
<p>The following are a collection of events between {{ time .Since .UserTimeZone }} - {{ time .Now .UserTimeZone}}</p>
<br>
  {{ range $sg, $types := .GroupReports }} 
  <h3>Group: {{ $sg }}</h3>
		{{ range $type_id, $events := $types}}
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

func EventPrinter(typeID int32, events []*ScanGroupReport) string {
	var buf bytes.Buffer

	switch typeID {
	// TODO: join certexpired/expiring format the date field to a string and filter out where subjectname [0] and port [1] are the same
	// so we don't have duplicate reports for the same certs
	case am.EventCertExpired:
	case am.EventCertExpiring:
		buf.WriteString("The following certificates will expire soon:")
		buf.WriteString("<br>\n<ul>")
		for _, e := range events {
			event := fmt.Sprintf("\n\t<li>%s on port %s expires in %s</li>\n", e.Data[0], e.Data[1], e.Data[2])
			buf.WriteString(event)
		}
		buf.WriteString("</ul>")
	case am.EventInitialGroupComplete:
	case am.EventMaxHostPricing:
	case am.EventNewHost:
		buf.WriteString("The following new hosts were found:")

		for _, e := range events {
			buf.WriteString("<br>\n<ul>")
			for _, host := range e.Data {
				buf.WriteString("\n\t<li>" + host + "</li>\n")
			}
			buf.WriteString("</ul>")
		}
	case am.EventAXFR:
		buf.WriteString("The following name servers were found leaking hostnames via Zone Transfer")

		for _, e := range events {
			buf.WriteString("<br>\n<ul>")
			for _, nsServer := range e.Data {
				buf.WriteString("\n\t<li>" + nsServer + "</li>\n")
			}
			buf.WriteString("</ul>")
		}
	case am.EventNSEC:
		buf.WriteString("The following name servers are leaking hostnames via NSEC records")

		for _, e := range events {
			buf.WriteString("<br>\n<ul>")
			for _, nsServer := range e.Data {
				buf.WriteString("\n\t<li>" + nsServer + "</li>\n")
			}
			buf.WriteString("</ul>")
		}
	case am.EventNewWebsite:
		buf.WriteString("The following new web sites were found:")

		for _, e := range events {
			buf.WriteString("<br>\n<ul>")
			for i := 0; i < len(e.Data); i += 2 {
				buf.WriteString(fmt.Sprintf("\n\t<li><a href=\"%s\">%s</a> on port %s</li>\n", e.Data[i], e.Data[i], e.Data[i+1]))
			}
			buf.WriteString("</ul>")
		}
	case am.EventWebHTMLUpdated:
	case am.EventWebJSChanged:
	case am.EventWebTechChanged:
	}
	return buf.String()
}
