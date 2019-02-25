-- +goose Up
-- SQL in this section is executed when the migration is applied.

DROP TRIGGER owner_user on am.organizations;
DROP FUNCTION update_owner_user;

-- Update user/org columns first
alter table only am.organizations add column creation_timetz timestamptz not null default '2018-05-30'; 
update am.organizations as new set creation_timetz=(select to_timestamp(new.creation_time/1000000000) from am.organizations 
    as dt where dt.organization_id=new.organization_id);
alter table only am.organizations drop column creation_time;
alter table only am.organizations rename column creation_timetz to creation_time;

alter table only am.users add column creation_timetz timestamptz not null default '2018-05-30'; 
update am.users as new set creation_timetz=(select to_timestamp(new.creation_time/1000000000) from am.users 
    as dt where dt.user_id=new.user_id);
alter table only am.users drop column creation_time;
alter table only am.users rename column creation_timetz to creation_time;

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_owner_user() RETURNS TRIGGER AS $body$
BEGIN 
    IF NEW.owner_email <> OLD.owner_email OR NEW.first_name <> OLD.first_name OR NEW.last_name <> OLD.last_name THEN
        UPDATE am.users set email=NEW.owner_email, first_name=NEW.first_name, last_name=NEW.last_name where organization_id=OLD.organization_id and email=OLD.owner_email;
    END IF;
    RETURN NEW;
END;
$body$ language plpgsql;
-- +goose StatementEnd

CREATE TRIGGER owner_user AFTER UPDATE ON am.organizations FOR EACH ROW execute procedure update_owner_user();

-- scangroup.creation_time
alter table only am.scan_group add column creation_timetz timestamptz not null default '2018-05-30';  
update am.scan_group as new set creation_timetz=(select to_timestamp(new.creation_time/1000000000) from am.scan_group 
    as dt where dt.scan_group_id=new.scan_group_id);
alter table only am.scan_group drop column creation_time;
alter table only am.scan_group rename column creation_timetz to creation_time;

-- scangroup.modified_time
alter table only am.scan_group add column modified_timetz timestamptz not null default '2018-05-30';  
update am.scan_group as new set modified_timetz=(select to_timestamp(new.modified_time/1000000000) from am.scan_group 
    as dt where dt.scan_group_id=new.scan_group_id);
alter table only am.scan_group drop column modified_time;
alter table only am.scan_group rename column modified_timetz to modified_time;

-- scan_group_events:
drop table am.scan_group_events;
CREATE TABLE am.scan_group_events (
    event_id bigserial not null primary key,
    organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_id bigint REFERENCES am.scan_group (scan_group_id),
    event_user_id integer references am.users (user_id), 
    event_time timestamptz not null default '2018-05-30',
    event_description text,
    event_from required_text
);

drop table am.scan_group_findings;
CREATE TABLE am.scan_group_findings (
    organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_id integer REFERENCES am.scan_group (scan_group_id),
    address_id integer REFERENCES am.scan_group_addresses (address_id) on delete cascade,
    finding_id integer REFERENCES am.scan_finding_types (finding_id),
    discovered_timestamp timestamptz not null default '2018-05-30',
    last_seen_timestamp timestamptz not null default '2018-05-30',
    custom_description text,
    data jsonb
);

-- scan_group_address:
alter table only am.scan_group_addresses add column discovered_timestamptz timestamptz not null default '2018-05-30'; 
update am.scan_group_addresses as sga set discovered_timestamptz=(
    select to_timestamp(sga.discovered_timestamp/1000000000) from am.scan_group_addresses as dt where dt.address_id=sga.address_id);
alter table only am.scan_group_addresses drop column discovered_timestamp;
alter table only am.scan_group_addresses rename column discovered_timestamptz to discovered_timestamp;

-- last_scanned_timestamp:
alter table only am.scan_group_addresses add column last_scanned_timestamptz timestamptz not null default '2018-05-30';  
update am.scan_group_addresses as sga set last_scanned_timestamptz=(select to_timestamp(sga.last_scanned_timestamp/1000000000) from am.scan_group_addresses as dt where dt.address_id=sga.address_id);
alter table only am.scan_group_addresses drop column last_scanned_timestamp;
alter table only am.scan_group_addresses rename column last_scanned_timestamptz to last_scanned_timestamp;

-- last_seen_timestamp:
alter table only am.scan_group_addresses add column last_seen_timestamptz timestamptz not null default '2018-05-30'; 
update am.scan_group_addresses as sga set last_seen_timestamptz=(select to_timestamp(sga.last_seen_timestamp/1000000000) from am.scan_group_addresses as dt where dt.address_id=sga.address_id);
alter table only am.scan_group_addresses drop column last_seen_timestamp;
alter table only am.scan_group_addresses rename column last_seen_timestamptz to last_seen_timestamp;


-- webdata

-- web_snapshots
alter table only am.web_snapshots add column response_timestamptz timestamptz not null default '2018-05-30';
update am.web_snapshots as new set response_timestamptz=(select to_timestamp(new.response_timestamp/1000000000) from am.web_snapshots 
    as dt where dt.snapshot_id=new.snapshot_id);
alter table only am.web_snapshots drop column response_timestamp;
alter table only am.web_snapshots rename column response_timestamptz to response_timestamp;

-- web_responses
alter table only am.web_responses add column response_timestamptz timestamptz not null default '2018-05-30';
update am.web_responses as new set response_timestamptz=(select to_timestamp(new.response_timestamp/1000000000) from am.web_responses 
    as dt where dt.response_id=new.response_id);
alter table only am.web_responses drop column response_timestamp;
alter table only am.web_responses rename column response_timestamptz to response_timestamp;

alter table only am.web_responses add column url_request_timestamptz timestamptz not null default '2018-05-30';
update am.web_responses as new set url_request_timestamptz=(select to_timestamp(new.url_request_timestamp/1000000000) from am.web_responses 
    as dt where dt.response_id=new.response_id);
alter table only am.web_responses drop column url_request_timestamp;
alter table only am.web_responses rename column url_request_timestamptz to url_request_timestamp;

-- web certificates
alter table only am.web_certificates add column response_timestamptz timestamptz not null default '2018-05-30';
update am.web_certificates as new set response_timestamptz=(select to_timestamp(new.response_timestamp/1000000000) from am.web_certificates 
    as dt where dt.certificate_id=new.certificate_id);
alter table only am.web_certificates drop column response_timestamp;
alter table only am.web_certificates rename column response_timestamptz to response_timestamp;


-- big data
-- certificiate_queries
alter table only am.certificate_queries add column query_timestamptz timestamptz not null default '2018-05-30';
update am.certificate_queries as new set query_timestamptz=(select to_timestamp(new.query_timestamp/1000000000) from am.certificate_queries 
    as dt where dt.id=new.id);
alter table only am.certificate_queries drop column query_timestamp;
alter table only am.certificate_queries rename column query_timestamptz to query_timestamp;

-- certificates
alter table only am.certificates add column inserted_timestamptz timestamptz not null default '2018-05-30';
update am.certificates as new set inserted_timestamptz=(select to_timestamp(new.inserted_timestamp/1000000000) from am.certificates 
    as dt where dt.certificate_id=new.certificate_id);
alter table only am.certificates drop column inserted_timestamp;
alter table only am.certificates rename column inserted_timestamptz to inserted_timestamp;

-- certificate_queries_subdomains
alter table only am.certificate_queries_subdomains add column query_timestamptz timestamptz not null default '2018-05-30';
update am.certificate_queries_subdomains as new set query_timestamptz=(select to_timestamp(new.query_timestamp/1000000000) from am.certificate_queries_subdomains 
    as dt where dt.etld_id=new.etld_id);
alter table only am.certificate_queries_subdomains drop column query_timestamp;
alter table only am.certificate_queries_subdomains rename column query_timestamptz to query_timestamp;

-- certificate_subdomains
alter table only am.certificate_subdomains add column inserted_timestamptz timestamptz not null default '2018-05-30';
update am.certificate_subdomains as new set inserted_timestamptz=(select to_timestamp(new.inserted_timestamp/1000000000) from am.certificate_subdomains 
    as dt where dt.subdomain_id=new.subdomain_id);
alter table only am.certificate_subdomains drop column inserted_timestamp;
alter table only am.certificate_subdomains rename column inserted_timestamptz to inserted_timestamp;

-- scan group activity
alter table only am.scan_group_activity add column batch_starttz timestamptz not null default '2018-05-30';
update am.scan_group_activity as new set batch_starttz=(select to_timestamp(new.batch_start/1000000000) from am.scan_group_activity 
    as dt where dt.organization_id=new.organization_id and dt.scan_group_id=new.scan_group_id);
alter table only am.scan_group_activity drop column batch_start;
alter table only am.scan_group_activity rename column batch_starttz to batch_start;

alter table only am.scan_group_activity add column batch_endtz timestamptz not null default '2018-05-30';
update am.scan_group_activity as new set batch_endtz=(select to_timestamp(new.batch_end/1000000000) from am.scan_group_activity 
    as dt where dt.organization_id=new.organization_id and dt.scan_group_id=new.scan_group_id);
alter table only am.scan_group_activity drop column batch_end;
alter table only am.scan_group_activity rename column batch_endtz to batch_end;

alter table only am.scan_group_activity add column last_updatedtz timestamptz not null default '2018-05-30';
update am.scan_group_activity as new set last_updatedtz=(select to_timestamp(new.last_updated/1000000000) from am.scan_group_activity 
    as dt where dt.organization_id=new.organization_id and dt.scan_group_id=new.scan_group_id);
alter table only am.scan_group_activity drop column last_updated;
alter table only am.scan_group_activity rename column last_updatedtz to last_updated;


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TRIGGER owner_user on am.organizations;
DROP FUNCTION update_owner_user;

-- Update user/org columns first
alter table only am.organizations add column creation_timetz bigint not null default 0; 
update am.organizations as new set creation_timetz=(select extract(epoch FROM new.creation_time) from am.organizations 
    as dt where dt.organization_id=new.organization_id);
alter table only am.organizations drop column creation_time;
alter table only am.organizations rename column creation_timetz to creation_time;

alter table only am.users add column creation_timetz bigint not null default 0; 
update am.users as new set creation_timetz=(select extract(epoch FROM new.creation_time) from am.users 
    as dt where dt.user_id=new.user_id);
alter table only am.users drop column creation_time;
alter table only am.users rename column creation_timetz to creation_time;

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_owner_user() RETURNS TRIGGER AS $body$
BEGIN 
    IF NEW.owner_email <> OLD.owner_email OR NEW.first_name <> OLD.first_name OR NEW.last_name <> OLD.last_name THEN
        UPDATE am.users set email=NEW.owner_email, first_name=NEW.first_name, last_name=NEW.last_name where organization_id=OLD.organization_id and email=OLD.owner_email;
    END IF;
    RETURN NEW;
END;
$body$ language plpgsql;
-- +goose StatementEnd

CREATE TRIGGER owner_user AFTER UPDATE ON am.organizations FOR EACH ROW execute procedure update_owner_user();

-- scangroup.creation_time
alter table only am.scan_group add column creation_timetz bigint not null default 0;  
update am.scan_group as new set creation_timetz=(select extract(epoch FROM new.creation_time) from am.scan_group 
    as dt where dt.scan_group_id=new.scan_group_id);
alter table only am.scan_group drop column creation_time;
alter table only am.scan_group rename column creation_timetz to creation_time;

-- scangroup.modified_time
alter table only am.scan_group add column modified_timetz bigint not null default 0;  
update am.scan_group as new set modified_timetz=(select extract(epoch FROM new.modified_time) from am.scan_group 
    as dt where dt.scan_group_id=new.scan_group_id);
alter table only am.scan_group drop column modified_time;
alter table only am.scan_group rename column modified_timetz to modified_time;

-- scan_group_events:
drop table am.scan_group_events;
CREATE TABLE am.scan_group_events (
    event_id bigserial not null primary key,
    organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_id bigint REFERENCES am.scan_group (scan_group_id),
    event_user_id integer references am.users (user_id), 
    event_time bigint not null default 0,
    event_description text,
    event_from required_text
);

drop table am.scan_group_findings;
CREATE TABLE am.scan_group_findings (
    organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_id integer REFERENCES am.scan_group (scan_group_id),
    address_id integer REFERENCES am.scan_group_addresses (address_id) on delete cascade,
    finding_id integer REFERENCES am.scan_finding_types (finding_id),
    discovered_timestamp bigint not null default 0,
    last_seen_timestamp bigint not null default 0,
    custom_description text,
    data jsonb
);

-- scan_group_address:
alter table only am.scan_group_addresses add column discovered_timestamptz bigint not null default 0; 
update am.scan_group_addresses as sga set discovered_timestamptz=(select extract(epoch FROM sga.discovered_timestamp) from am.scan_group_addresses as dt where dt.address_id=sga.address_id);
alter table only am.scan_group_addresses drop column discovered_timestamp;
alter table only am.scan_group_addresses rename column discovered_timestamptz to discovered_timestamp;

-- last_scanned_timestamp:
alter table only am.scan_group_addresses add column last_scanned_timestamptz bigint not null default 0;  
update am.scan_group_addresses as sga set last_scanned_timestamptz=(select extract(epoch FROM sga.last_scanned_timestamp) from am.scan_group_addresses as dt where dt.address_id=sga.address_id);
alter table only am.scan_group_addresses drop column last_scanned_timestamp;
alter table only am.scan_group_addresses rename column last_scanned_timestamptz to last_scanned_timestamp;

-- last_seen_timestamp:
alter table only am.scan_group_addresses add column last_seen_timestamptz bigint not null default 0; 
update am.scan_group_addresses as sga set last_seen_timestamptz=(select extract(epoch FROM sga.last_seen_timestamp) from am.scan_group_addresses as dt where dt.address_id=sga.address_id);
alter table only am.scan_group_addresses drop column last_seen_timestamp;
alter table only am.scan_group_addresses rename column last_seen_timestamptz to last_seen_timestamp;

-- webdata

-- web_snapshots
alter table only am.web_snapshots add column response_timestamptz bigint not null default 0;
update am.web_snapshots as new set response_timestamptz=(select extract(epoch FROM new.response_timestamp) from am.web_snapshots 
    as dt where dt.snapshot_id=new.snapshot_id);
alter table only am.web_snapshots drop column response_timestamp;
alter table only am.web_snapshots rename column response_timestamptz to response_timestamp;

-- web_responses
alter table only am.web_responses add column response_timestamptz bigint not null default 0;
update am.web_responses as new set response_timestamptz=(select extract(epoch FROM new.response_timestamp) from am.web_responses 
    as dt where dt.response_id=new.response_id);
alter table only am.web_responses drop column response_timestamp;
alter table only am.web_responses rename column response_timestamptz to response_timestamp;

alter table only am.web_responses add column url_request_timestamptz bigint not null default 0;
update am.web_responses as new set url_request_timestamptz=(select extract(epoch FROM new.url_request_timestamp) from am.web_responses 
    as dt where dt.response_id=new.response_id);
alter table only am.web_responses drop column url_request_timestamp;
alter table only am.web_responses rename column url_request_timestamptz to url_request_timestamp;

-- web certificates
alter table only am.web_certificates add column response_timestamptz bigint not null default 0;
update am.web_certificates as new set response_timestamptz=(select extract(epoch FROM new.response_timestamp) from am.web_certificates 
    as dt where dt.certificate_id=new.certificate_id);
alter table only am.web_certificates drop column response_timestamp;
alter table only am.web_certificates rename column response_timestamptz to response_timestamp;


-- big data
-- certificiate_queries
alter table only am.certificate_queries add column query_timestamptz bigint not null default 0;
update am.certificate_queries as new set query_timestamptz=(select extract(epoch FROM new.query_timestamp) from am.certificate_queries 
    as dt where dt.id=new.id);
alter table only am.certificate_queries drop column query_timestamp;
alter table only am.certificate_queries rename column query_timestamptz to query_timestamp;

-- certificates
alter table only am.certificates add column inserted_timestamptz bigint not null default 0;
update am.certificates as new set inserted_timestamptz=(select extract(epoch FROM new.inserted_timestamp) from am.certificates 
    as dt where dt.certificate_id=new.certificate_id);
alter table only am.certificates drop column inserted_timestamp;
alter table only am.certificates rename column inserted_timestamptz to inserted_timestamp;

-- certificate_queries_subdomains
alter table only am.certificate_queries_subdomains add column query_timestamptz bigint not null default 0;
update am.certificate_queries_subdomains as new set query_timestamptz=(select extract(epoch FROM new.query_timestamp) from am.certificate_queries_subdomains 
    as dt where dt.etld_id=new.etld_id);
alter table only am.certificate_queries_subdomains drop column query_timestamp;
alter table only am.certificate_queries_subdomains rename column query_timestamptz to query_timestamp;

-- certificate_subdomains
alter table only am.certificate_subdomains add column inserted_timestamptz bigint not null default 0;
update am.certificate_subdomains as new set inserted_timestamptz=(select extract(epoch FROM new.inserted_timestamp) from am.certificate_subdomains 
    as dt where dt.subdomain_id=new.subdomain_id);
alter table only am.certificate_subdomains drop column inserted_timestamp;
alter table only am.certificate_subdomains rename column inserted_timestamptz to inserted_timestamp;

-- scan group activity
alter table only am.scan_group_activity add column batch_starttz bigint not null default 0;
update am.scan_group_activity as new set batch_starttz=(select extract(epoch FROM new.batch_start) from am.scan_group_activity 
    as dt where dt.organization_id=new.organization_id and dt.scan_group_id=new.scan_group_id);
alter table only am.scan_group_activity drop column batch_start;
alter table only am.scan_group_activity rename column batch_starttz to batch_start;

alter table only am.scan_group_activity add column batch_endtz bigint not null default 0;
update am.scan_group_activity as new set batch_endtz=(select extract(epoch FROM new.batch_end) from am.scan_group_activity 
    as dt where dt.organization_id=new.organization_id and dt.scan_group_id=new.scan_group_id);
alter table only am.scan_group_activity drop column batch_end;
alter table only am.scan_group_activity rename column batch_endtz to batch_end;

alter table only am.scan_group_activity add column last_updatedtz bigint not null default 0;
update am.scan_group_activity as new set last_updatedtz=(select extract(epoch FROM new.last_updated) from am.scan_group_activity 
    as dt where dt.organization_id=new.organization_id and dt.scan_group_id=new.scan_group_id);
alter table only am.scan_group_activity drop column last_updated;
alter table only am.scan_group_activity rename column last_updatedtz to last_updated;