-- +goose Up
-- SQL in this section is executed when the migration is applied.
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION am.delete_org(org_id integer) RETURNS void as 
$BODY$
BEGIN 
    delete from am.ladon_role_member where organization_id=org_id;
    delete from am.ladon_role where organization_id=org_id;
    delete from am.scan_group_activity where organization_id=org_id;
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

create table am.web_techtypes (
    techtype_id serial primary key,
    techname varchar(512) not null default '',
    category_id int not null,
    category varchar(512) not null,
    icon varchar(128) not null default '',
    website varchar(1024) not null default '',
    ignored boolean default false,
    unique(techname, category)
);

create table am.web_technologies (
    tech_id bigserial not null primary key,
    organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_id bigint REFERENCES am.scan_group (scan_group_id),
    snapshot_id bigint REFERENCES am.web_snapshots (snapshot_id),
    techtype_id int references am.web_techtypes (techtype_id),
    matched_text text not null default '',
    match_location text not null default '',
    version varchar(512) not null default '',
    UNIQUE(snapshot_id,techtype_id,match_location)
);

grant select, insert, update, delete on am.web_technologies to webdataservice;
grant select on am.web_techtypes to webdataservice;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION am.delete_org(org_id integer) RETURNS void as 
$BODY$
BEGIN 
    delete from am.ladon_role_member where organization_id=org_id;
    delete from am.ladon_role where organization_id=org_id;
    delete from am.scan_group_activity where organization_id=org_id;
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

revoke select, insert, update, delete on am.web_technologies from webdataservice;
revoke select on am.web_techtypes from webdataservice;
drop table am.web_technologies;
drop table am.web_techtypes;
