-- +goose Up
-- SQL in this section is executed when the migration is applied.

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

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION am.delete_org(org_id integer) RETURNS void as 
$BODY$
BEGIN 
    delete from am.ladon_role_member where organization_id=org_id;
    delete from am.ladon_role where organization_id=org_id;
    delete from am.scan_group_findings where organization_id=org_id;
    delete from am.scan_group_addresses where organization_id=org_id;
    delete from am.scan_group_events where organization_id=org_id;
    delete from am.scan_group where organization_id=org_id;
    delete from am.users where organization_id=org_id; 
    delete from am.organizations where organization_id=org_id;

END
$BODY$ LANGUAGE plpgsql SECURITY DEFINER;
-- +goose StatementEnd
