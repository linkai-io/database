package reports

import (
	"time"

	sq "github.com/Masterminds/squirrel"
)

func buildGetReportQuery(orgID, userID int, since time.Time) (string, []interface{}, error) {
	sub := sq.Select().Columns(
		"organization_id",
		"user_id",
		"type_id",
		"subscribed").From("am.user_notification_subscriptions").Where(sq.Eq{"organization_id": orgID}).
		Where(sq.Eq{"user_id": userID}).
		Where(sq.Eq{"subscribed": true})

	p := sq.Select().Columns("subs.organization_id",
		"sg.scan_group_id",
		"sg.scan_group_name",
		"subs.user_id",
		"subs.type_id",
		"events.event_timestamp",
		"events.event_data").FromSelect(sub, "subs").
		Join("am.event_notifications as events on subs.type_id=events.type_id and events.organization_id=subs.organization_id").
		Join("am.scan_group as sg on events.scan_group_id=sg.scan_group_id and events.organization_id=sg.organization_id").
		Where(sq.Eq{"sg.deleted": false}).
		Where(sq.GtOrEq{"events.event_timestamp": since})

	return p.PlaceholderFormat(sq.Dollar).ToSql()
}
