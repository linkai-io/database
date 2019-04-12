package reports

import (
	"time"

	sq "github.com/Masterminds/squirrel"
)

func buildGetReportQuery(orgID, userID int, weekly bool, since time.Time) (string, []interface{}, error) {
	sub := sq.Select().Columns(
		"organization_id",
		"user_id",
		"type_id",
		"subscribed_since",
		"subscribed").From("am.user_notification_subscriptions").Where(sq.Eq{"organization_id": orgID}).
		Where(sq.Eq{"user_id": userID}).
		Where(sq.Eq{"subscribed": true})

	p := sq.Select().Columns("subs.organization_id",
		"sg.scan_group_name",
		"subs.user_id",
		"subs.type_id",
		"events.event_timestamp",
		"events.event_data",
		"settings.user_timezone",
		"settings.should_daily_email",
		"settings.should_weekly_email").FromSelect(sub, "subs").
		Join("am.event_notifications as events on subs.type_id=events.type_id and events.organization_id=subs.organization_id").
		Join("am.scan_group as sg on events.scan_group_id=sg.scan_group_id and events.organization_id=sg.organization_id").
		Join("am.user_notification_settings as settings on subs.user_id=settings.user_id").
		Where(sq.Eq{"sg.deleted": false}).
		Where(sq.Eq{"settings.user_id": userID}).
		Where(sq.GtOrEq{"events.event_timestamp": since})

	if weekly {
		p = p.Where(sq.Eq{"should_weekly_email": true})
	} else {
		p = p.Where(sq.Eq{"should_daily_email": true})
	}

	return p.PlaceholderFormat(sq.Dollar).ToSql()
}
