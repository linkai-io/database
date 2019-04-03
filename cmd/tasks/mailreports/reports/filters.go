package reports

import (
	"time"

	sq "github.com/Masterminds/squirrel"
)

func buildGetReportQuery(orgID, userID int, weekly bool, since time.Time) (string, []interface{}, error) {
	p := sq.Select().Columns("sg.scan_group_name",
		"en.type_id",
		"en.event_timestamp",
		"settings.user_timezone",
		"en.event_data").From("am.event_notifications as en").
		Join("lateral (select user_id, type_id, event_timestamp from am.user_notification_subscriptions as uns where uns.subscribed=true and en.type_id=uns.type_id and en.event_timestamp >= uns.subscribed_since) as uns on true").
		Join("am.scan_group as sg on sg.scan_group_id=sg.scan_group_id").
		Join("am.user_notification_settings as settings on settings.user_id=uns.user_id").
		Where(sq.Eq{"uns.user_id": userID}).
		Where(sq.Eq{"sg.deleted": false}).
		Where(sq.Eq{"en.organization_id": orgID}).
		Where(sq.GtOrEq{"en.event_timestamp": since})

	if weekly {
		p = p.Where(sq.Eq{"should_weekly_email": true})
	} else {
		p = p.Where(sq.Eq{"should_daily_email": true})
	}

	return p.PlaceholderFormat(sq.Dollar).ToSql()
}
