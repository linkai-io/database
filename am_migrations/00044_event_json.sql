-- +goose Up
-- SQL in this section is executed when the migration is applied.
alter table only am.event_notifications add column event_data_json json default '{}';

alter table only am.user_notification_subscriptions add column webhook_version text default '';
alter table only am.user_notification_subscriptions add column webhook_enabled boolean default false;
alter table only am.user_notification_subscriptions add column webhook_url text default '';
alter table only am.user_notification_subscriptions add column webhook_type text default '';

alter table only am.user_notification_settings add column webhook_current_key text default '';
alter table only am.user_notification_settings add column webhook_previous_key text default '';

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
alter table only am.event_notifications drop column event_data_json;

alter table only am.user_notification_subscriptions drop column webhook_version;
alter table only am.user_notification_subscriptions drop column webhook_enabled;
alter table only am.user_notification_subscriptions drop column webhook_url;
alter table only am.user_notification_subscriptions drop column webhook_type;

alter table only am.user_notification_settings drop column webhook_current_key;
alter table only am.user_notification_settings drop column webhook_previous_key;

