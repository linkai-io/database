-- +goose Up
-- SQL in this section is executed when the migration is applied.
alter table only am.user_notification_subscriptions drop column webhook_version;
alter table only am.user_notification_subscriptions drop column webhook_enabled;
alter table only am.user_notification_subscriptions drop column webhook_url;
alter table only am.user_notification_subscriptions drop column webhook_type;

alter table only am.user_notification_settings drop column webhook_current_key;
alter table only am.user_notification_settings drop column webhook_previous_key;


create table am.webhook_event_settings (
    webhook_id serial primary key not null,
    organization_id integer REFERENCES am.organizations(organization_id),
    scan_group_id integer REFERENCES am.scan_group (scan_group_id),
    name text default '',
    events integer[] not null,
    enabled boolean default false,
    version text default '',
    url text default '',
    type text default '',
    current_key text default '',
    previous_key text default '',
    deleted boolean default false,
    UNIQUE(organization_id, scan_group_id, name)
);

create table am.webhook_events (
    webhook_event_id bigserial primary key not null,
    organization_id integer REFERENCES am.organizations(organization_id),
    scan_group_id integer REFERENCES am.scan_group (scan_group_id),
    notification_id bigint REFERENCES am.event_notifications (notification_id),
    webhook_id integer REFERENCES am.webhook_event_settings (webhook_id),
    type_id integer REFERENCES am.event_notification_types (type_id),
    last_attempt_timestamp timestamptz not null default 'epoch',
    last_attempt_status integer default 0,
    UNIQUE(organization_id, scan_group_id, webhook_id, notification_id)
);

create index webhook_events_timestamp_idx on am.webhook_events (organization_id, scan_group_id, last_attempt_timestamp);

grant select,insert,update,delete on am.webhook_events to eventservice;
grant select,insert,update,delete on am.webhook_event_settings to eventservice;

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION am.delete_org(org_id integer) RETURNS void as 
$BODY$
BEGIN 
    -- permissions
    delete from am.ladon_role_member where organization_id=org_id;
    delete from am.ladon_role where organization_id=org_id;
    -- events
    delete from am.webhook_events where organization_id=org_id;
    delete from am.webhook_event_settings where organization_id=org_id;
    delete from am.user_notifications_read where organization_id=org_id;
    delete from am.user_notification_subscriptions where organization_id=org_id;
    delete from am.user_notification_settings where organization_id=org_id;
    delete from am.event_notifications_archive where organization_id=org_id;
    delete from am.event_notifications where organization_id=org_id;
    -- web
    delete from am.web_technologies where organization_id=org_id;
    delete from am.web_snapshots where organization_id=org_id;
    delete from am.web_responses where organization_id=org_id;
    delete from am.web_certificates where organization_id=org_id;
    -- scangroups
	delete from am.scan_group_addresses_ports where organization_id=org_id;
    delete from am.scan_group_activity where organization_id=org_id;
    delete from am.scan_group_addresses_overflow where organization_id=org_id;
    delete from am.scan_group_findings where organization_id=org_id;
    delete from am.scan_group_addresses where organization_id=org_id;
    delete from am.scan_group_events where organization_id=org_id;
    delete from am.scan_group where organization_id=org_id;
    -- org
    delete from am.users where organization_id=org_id; 
    delete from am.organizations where organization_id=org_id;
	-- archives
	delete from am.scan_group_addresses_ports_archive where organization_id=org_id;
	delete from am.scan_group_addresses_archive where organization_id=org_id;
	delete from am.web_responses_archive where organization_id=org_id;
	delete from am.web_technologies_archive where organization_id=org_id;
	delete from am.web_snapshots_archive where organization_id=org_id;
END
$BODY$ LANGUAGE plpgsql SECURITY DEFINER;
-- +goose StatementEnd

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION am.delete_org(org_id integer) RETURNS void as 
$BODY$
BEGIN 
    -- permissions
    delete from am.ladon_role_member where organization_id=org_id;
    delete from am.ladon_role where organization_id=org_id;
    -- events
    delete from am.user_notifications_read where organization_id=org_id;
    delete from am.user_notification_subscriptions where organization_id=org_id;
    delete from am.user_notification_settings where organization_id=org_id;
    delete from am.event_notifications_archive where organization_id=org_id;
    delete from am.event_notifications where organization_id=org_id;
    -- web
    delete from am.web_technologies where organization_id=org_id;
    delete from am.web_snapshots where organization_id=org_id;
    delete from am.web_responses where organization_id=org_id;
    delete from am.web_certificates where organization_id=org_id;
    -- scangroups
	delete from am.scan_group_addresses_ports where organization_id=org_id;
    delete from am.scan_group_activity where organization_id=org_id;
    delete from am.scan_group_addresses_overflow where organization_id=org_id;
    delete from am.scan_group_findings where organization_id=org_id;
    delete from am.scan_group_addresses where organization_id=org_id;
    delete from am.scan_group_events where organization_id=org_id;
    delete from am.scan_group where organization_id=org_id;
    -- org
    delete from am.users where organization_id=org_id; 
    delete from am.organizations where organization_id=org_id;
	-- archives
	delete from am.scan_group_addresses_ports_archive where organization_id=org_id;
	delete from am.scan_group_addresses_archive where organization_id=org_id;
	delete from am.web_responses_archive where organization_id=org_id;
	delete from am.web_technologies_archive where organization_id=org_id;
	delete from am.web_snapshots_archive where organization_id=org_id;
END
$BODY$ LANGUAGE plpgsql SECURITY DEFINER;
-- +goose StatementEnd

revoke select,insert,update,delete on am.webhook_events from eventservice;
revoke select,insert,update,delete on am.webhook_event_settings from eventservice;

drop index am.webhook_events_timestamp_idx;

drop table am.webhook_events;
drop table am.webhook_event_settings;

alter table only am.user_notification_subscriptions add column webhook_version text default '';
alter table only am.user_notification_subscriptions add column webhook_enabled boolean default false;
alter table only am.user_notification_subscriptions add column webhook_url text default '';
alter table only am.user_notification_subscriptions add column webhook_type text default '';

alter table only am.user_notification_settings add column webhook_current_key text default '';
alter table only am.user_notification_settings add column webhook_previous_key text default '';