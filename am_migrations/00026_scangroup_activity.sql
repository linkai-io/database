-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE am.scan_group_activity (
    organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_id bigint REFERENCES am.scan_group (scan_group_id),
    active_addresses int not null default 0,
    batch_size int not null default 0,
    last_updated bigint not null default 0,
    batch_start bigint not null default 0,
    batch_end bigint not null default 0,
    UNIQUE(organization_id, scan_group_id)
);

grant select, insert, update, delete on am.scan_group_activity to scangroupservice;

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

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION am.delete_org(org_id integer) RETURNS void as 
$BODY$
BEGIN 
    delete from am.ladon_role_member where organization_id=org_id;
    delete from am.ladon_role where organization_id=org_id;
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

revoke select, insert, update, delete on am.scan_group_activity from scangroupservice;
DROP TABLE am.scan_group_activity;