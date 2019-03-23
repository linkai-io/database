-- +goose Up
-- SQL in this section is executed when the migration is applied.
create table am.event_notification_types (
    type_id int primary key not null,
    notify_description text not null default ''
);

insert into am.event_notification_types (type_id, notify_description) values 
    (1, 'initial scan group analysis completed'),
    (2, 'maximum number of hostnames reached for pricing plan'),
    (10, 'new hostname'),
    (11, 'new record'),
    (100, 'new website detected'),
    (101, 'website''s html updated'),
    (102, 'website''s technology changed'),
    (103, 'website''s javascript changed'),
    (150, 'certificate expiring in 30 days'),
    (151, 'certificate expiring in 15 days'),
    (152, 'certificate expiring in 5 days'),
    (153, 'certificate expiring in 1 day'),
    (154, 'certificate expired'),
    (200, 'dns server exposing records via zone transfer'),
    (201, 'dns server exposing records via NSEC walking');

create table am.event_notifications (
    notification_id bigserial primary key not null,
    organization_id int references am.organizations (organization_id),
    scan_group_id int references am.scan_group (scan_group_id),
    type_id int references am.event_notification_types (type_id),
    event_timestamp timestamptz not null,
    event_data jsonb
);

create table am.event_notifications_archive (
    notification_id serial primary key not null,
    organization_id int references am.organizations (organization_id),
    scan_group_id int references am.scan_group (scan_group_id),
    event_timestamp timestamptz not null,
    type_id int references am.event_notification_types (type_id),
    event_data jsonb
);

create table am.user_notification_settings (
    organization_id int references am.organizations (organization_id),
    user_id int references am.users (user_id),
    weekly_report_send_day int not null default 0,
    daily_report_send_hour int not null default 0,
    user_timezone varchar(128) not null default '',
    should_daily_email boolean default false,
    should_weekly_email boolean default false,
    UNIQUE(organization_id,user_id)
);

create table am.user_notification_subscriptions (
    organization_id int references am.organizations (organization_id),
    user_id int references am.users (user_id),
    type_id int references am.event_notification_types (type_id),
    subscribed_since timestamptz not null,
    subscribed boolean not null default false,
    UNIQUE(organization_id, user_id, type_id)
);

create table am.user_notifications_read (
    organization_id int references am.organizations (organization_id),
    user_id int references am.users (user_id),
    notification_id bigint references am.event_notifications (notification_id),
    UNIQUE(organization_id,user_id,notification_id)
);


grant select on am.event_notification_types to eventservice;
grant select, insert, update, delete on am.event_notifications to eventservice;
grant select, insert, update, delete on am.event_notifications_archive to eventservice;
grant select, insert, update, delete on am.user_notification_settings to eventservice;
grant select, insert, update, delete on am.user_notifications_read to eventservice;
grant select, insert, update, delete on am.user_notification_subscriptions to eventservice;

-- permissions for building reports
grant select on am.web_certificates to eventservice;
grant select on am.web_responses to eventservice;
grant select on am.web_snapshots to eventservice;
grant select on am.web_technologies to eventservice;
grant select on am.scan_group_addresses to eventservice;

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
    delete from am.scan_group_activity where organization_id=org_id;
    delete from am.scan_group_addresses_overflow where organization_id=org_id;
    delete from am.scan_group_findings where organization_id=org_id;
    delete from am.scan_group_addresses where organization_id=org_id;
    delete from am.scan_group_events where organization_id=org_id;
    delete from am.scan_group where organization_id=org_id;
    -- org
    delete from am.users where organization_id=org_id; 
    delete from am.organizations where organization_id=org_id;
END
$BODY$ LANGUAGE plpgsql SECURITY DEFINER;
-- +goose StatementEnd

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION am.delete_org(org_id integer) RETURNS void as 
$BODY$
BEGIN 
    delete from am.ladon_role_member where organization_id=org_id;
    delete from am.ladon_role where organization_id=org_id;
    delete from am.scan_group_activity where organization_id=org_id;
    delete from am.scan_group_addresses_overflow where organization_id=org_id;
    delete from am.web_technologies where organization_id=org_id;
    delete from am.web_snapshots where organization_id=org_id;
    delete from am.web_responses where organization_id=org_id;
    delete from am.web_certificates where organization_id=org_id;
    delete from am.scan_group_findings where organization_id=org_id;
    delete from am.scan_group_addresses where organization_id=org_id;
    delete from am.scan_group_events where organization_id=org_id;
    delete from am.scan_group where organization_id=org_id;
    delete from am.users where organization_id=org_id; 
    delete from am.organizations where organization_id=org_id;
END
$BODY$ LANGUAGE plpgsql SECURITY DEFINER;
-- +goose StatementEnd

revoke select on am.web_certificates from eventservice;
revoke select on am.web_responses from eventservice;
revoke select on am.web_snapshots from eventservice;
revoke select on am.web_technologies from eventservice;
revoke select on am.scan_group_addresses from eventservice;

revoke select on am.event_notification_types from eventservice;
revoke select, insert, update, delete on am.event_notifications from eventservice;
revoke select, insert, update, delete on am.event_notifications_archive from eventservice;
revoke select, insert, update, delete on am.user_notification_settings from eventservice;
revoke select, insert, update, delete on am.user_notifications_read from eventservice;
revoke select, insert, update, delete on am.user_notification_subscriptions from eventservice;

drop table am.user_notifications_read;
drop table am.user_notification_subscriptions;
drop table am.user_notification_settings;
drop table am.event_notifications_archive;
drop table am.event_notifications;
drop table am.event_notification_types;
