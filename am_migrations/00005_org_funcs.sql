-- +goose Up
-- +goose StatementBegin
CREATE FUNCTION am.create_test_org(orgname varchar(256)) RETURNS void AS 
$BODY$
BEGIN
    with org as (
                insert into am.organizations (
                    organization_name, organization_custom_id, user_pool_id, identity_pool_id, 
                    owner_email, first_name, last_name, phone, country, state_prefecture, street, 
                    address1, address2, city, postal_code, creation_time, status_id, deleted, subscription_id
                )
                values 
                    (orgname, 'deadbeef', 'user.pool','identity.pool', orgname||'email@admin.org', 'first','last','1-111-111-1111','usa','ca','1 fake lane','','','bevery hills', '90219', 123, 1000, false, 1000)
                returning organization_id
            ) 
            insert into am.users (
                    organization_id, user_custom_id, email, first_name, last_name, user_status_id, creation_time, deleted
                ) 
                values
                    ( (select org.organization_id from org), 'd34db33f',  orgname||'email@admin.org', 'first','last', 1000, 111, false);
END 
$BODY$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE FUNCTION am.delete_org(org_id integer) RETURNS void as 
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

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP FUNCTION am.delete_org(org_id integer);
DROP FUNCTION am.create_test_org(orgname varchar(256));