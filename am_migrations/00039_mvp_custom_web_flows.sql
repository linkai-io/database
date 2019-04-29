-- +goose Up
-- SQL in this section is executed when the migration is applied.
create table am.custom_web_flows (
    web_flow_id serial primary key not null,
    organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_id integer REFERENCES am.scan_group (scan_group_id),
    web_flow_name varchar(128) not null default '',
    configuration jsonb,
    created_timestamp timestamptz not null,
    modified_timestamp timestamptz not null,
    deleted boolean default false,
    UNIQUE(organization_id, web_flow_name)
);

create table am.custom_web_flow_status (
    status_id serial primary key not null,
    web_flow_id integer REFERENCES am.custom_web_flows (web_flow_id),
    organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_id integer REFERENCES am.scan_group (scan_group_id),
    last_updated_timestamp timestamptz not null,
    started_timestamp timestamptz not null,
    finished_timestamp timestamptz not null,
    web_flow_status integer not null default 0,
    total integer not null default 0,
    in_progress integer not null default 0,
    completed integer not null default 0
);

create table am.custom_web_flow_results (
    result_id bigserial primary key not null,
    web_flow_id integer REFERENCES am.custom_web_flows (web_flow_id),
    organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_id integer REFERENCES am.scan_group (scan_group_id),
    run_timestamp timestamptz not null,
    url bytea not null default '',
    load_url bytea not null default '',
    load_host_address varchar(512) not null default '',
    load_ip_address varchar(256) not null default '',
    requested_port int not null default 0,
    response_port int not null default 0,
    response_timestamp timestamptz not null,
    result jsonb,
    response_body_hash varchar(512) not null default '',
    response_body_link text not null default ''
);

create index custom_web_flow_results_response_timestamp_index on am.custom_web_flow_results (response_timestamp desc);
create index custom_web_flow_results_run_timestamp_index on am.custom_web_flow_results (run_timestamp desc);

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION am.delete_org(org_id integer) RETURNS void as 
$BODY$
BEGIN 
    delete from am.ladon_role_member where organization_id=org_id;
    delete from am.ladon_role where organization_id=org_id;
    delete from am.custom_web_flow_results where organization_id=org_id;
    delete from am.custom_web_flow_status where organization_id=org_id;
    delete from am.custom_web_flows where organization_id=org_id;
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

-- generic permissions
grant linkai_user to customwebflowservice;
grant usage on schema am to customwebflowservice;
grant usage, select, update on all sequences in schema am to customwebflowservice;
grant references (organization_id) on table am.organizations to customwebflowservice;
grant references (user_id) on table am.users to customwebflowservice;
grant references (scan_group_id) on table am.scan_group to customwebflowservice;

-- service specific
grant select, insert, update, delete on am.custom_web_flow_results to customwebflowservice;
grant select, insert, update, delete on am.custom_web_flow_status to customwebflowservice;
grant select, insert, update, delete on am.custom_web_flows to customwebflowservice;

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


revoke select, insert, update, delete on am.custom_web_flow_results from customwebflowservice;
revoke select, insert, update, delete on am.custom_web_flow_status from customwebflowservice;
revoke select, insert, update, delete on am.custom_web_flows from customwebflowservice;


revoke references (organization_id) on table am.organizations from customwebflowservice;
revoke references (user_id) on table am.users from customwebflowservice;
revoke references (scan_group_id) on table am.scan_group from customwebflowservice;
revoke usage, select on all sequences in schema am from customwebflowservice;
revoke all on all sequences in schema am from customwebflowservice;
revoke usage on schema am from customwebflowservice;
revoke linkai_user from customwebflowservice;

drop index am.custom_web_flow_results_response_timestamp_index;
drop index am.custom_web_flow_results_run_timestamp_index;
drop table am.custom_web_flow_results;
drop table am.custom_web_flow_status;
drop table am.custom_web_flows;