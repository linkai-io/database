-- +goose Up
-- SQL in this section is executed when the migration is applied.
alter table only am.scan_group add column last_paused_timestamp timestamptz not null default 'epoch';
alter table only am.scan_group add column archive_after_days integer default 7;

alter table only am.users add column last_login_timestamp timestamptz not null default 'epoch';

insert into am.event_notification_types (type_id, notify_description) values 
	(12, 'new port open'),
	(13, 'port closed');

create table am.scan_group_addresses_ports (
	port_id bigserial primary key not null,
	organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_id integer REFERENCES am.scan_group (scan_group_id),
    host_address varchar(512) not null default '',
    ip_address varchar(256) not null default '',
	address_hash varchar(128) not null default '',
	port_data jsonb not null default '{}'::jsonb,
	scanned_timestamp timestamptz not null default 'epoch',
	previous_scanned_timestamp timestamptz not null default 'epoch',
	check (host_address is not null or ip_address is not null),
    UNIQUE(scan_group_id, host_address, ip_address)
);

create index scan_group_addresses_ports_scanned_timestamp_idx on am.scan_group_addresses_ports (scanned_timestamp desc);
create index scan_group_addresses_ports_address_hash_idx on am.scan_group_addresses_ports (address_hash);

create table am.scan_group_addresses_ports_archive (
	like am.scan_group_addresses_ports,
	archived_timestamp timestamptz not null default 'epoch' 
);

create table am.scan_group_addresses_archive (
	like am.scan_group_addresses, 
	archived_timestamp timestamptz not null default 'epoch'
);

create table am.web_responses_archive (
	like am.web_responses, 
	archived_timestamp timestamptz not null default 'epoch'
);

create table am.web_technologies_archive (
	like am.web_technologies, 
	archived_timestamp timestamptz not null default 'epoch'
);

create table am.web_snapshots_archive (
	like am.web_snapshots, 
	archived_timestamp timestamptz not null default 'epoch'
);

grant select,insert,update,delete on am.scan_group_addresses_ports to addressservice;
grant select,insert,update,delete on am.scan_group_addresses_ports_archive to addressservice;
grant select,insert,update,delete on am.scan_group_addresses_archive to addressservice;
grant select,insert,update,delete on am.web_responses_archive to webdataservice;
grant select,insert,update,delete on am.web_technologies_archive to webdataservice;
grant select,insert,update,delete on am.web_snapshots_archive to webdataservice;

grant select on am.web_techtypes to eventservice;
grant select on am.scan_group_addresses_ports to eventservice;

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

revoke select on am.web_techtypes from eventservice;
revoke select on am.scan_group_addresses_ports from eventservice;

revoke select,insert,update,delete on am.scan_group_addresses_ports from addressservice;
revoke select,insert,update,delete on am.scan_group_addresses_ports_archive from addressservice;
revoke select,insert,update,delete on am.scan_group_addresses_archive from addressservice;
revoke select,insert,update,delete on am.web_responses_archive from webdataservice;
revoke select,insert,update,delete on am.web_technologies_archive from webdataservice;
revoke select,insert,update,delete on am.web_snapshots_archive from webdataservice;

drop table am.web_snapshots_archive;
drop table am.web_technologies_archive;
drop table am.web_responses_archive;
drop table am.scan_group_addresses_archive;
drop table am.scan_group_addresses_ports_archive;

drop index am.scan_group_addresses_ports_address_hash_idx;
drop index am.scan_group_addresses_ports_scanned_timestamp_idx;
drop table am.scan_group_addresses_ports;

delete from am.event_notification_types where type_id=12;
delete from am.event_notification_types where type_id=13;

alter table only am.users drop column last_login_timestamp;
alter table only am.scan_group drop column archive_after_days;
alter table only am.scan_group drop column last_paused_timestamp;