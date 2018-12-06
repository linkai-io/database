-- +goose Up
-- SQL in this section is executed when the migration is applied.
revoke all on schema am from public;
revoke all on schema public from public;
revoke all on all functions in schema public from public;
revoke all on all functions in schema am from public;

-- allow all services to access the roles/perm/policy tables.
-- allows usage on am schema, but disable creating tables.
create role linkai_user with inherit;
grant usage on schema am to linkai_user;
grant select, update on all sequences in schema am to linkai_user;
alter default privileges in schema am grant usage on sequences to linkai_user;
alter default privileges in schema am grant select on sequences to linkai_user; 
grant select on am.ladon_role to linkai_user;
grant select on am.ladon_role_member to linkai_user;
grant select on am.ladon_policy to linkai_user;
grant select on am.ladon_policy_subject to linkai_user;
grant select on am.ladon_policy_permission to linkai_user;
grant select on am.ladon_policy_resource to linkai_user;
grant select on am.ladon_subject to linkai_user;
grant select on am.ladon_action to linkai_user;
grant select on am.ladon_resource to linkai_user;
grant select on am.ladon_policy_subject_rel to linkai_user;
grant select on am.ladon_policy_action_rel to linkai_user;
grant select on am.ladon_policy_resource_rel to linkai_user;


-- eventservice permissions
grant linkai_user to eventservice;
grant select, insert, update, delete on am.scan_group_events to eventservice;
grant references (organization_id) on table am.organizations to eventservice;
grant references (user_id) on table am.users to eventservice;
grant references (scan_group_id) on table am.scan_group to eventservice;


-- orgservice permissions
grant linkai_user to orgservice;
grant select, insert, update, delete on am.organizations to orgservice; 
grant select, insert, update, delete on am.users to orgservice; 
grant select, insert, update, delete on am.ladon_role to orgservice; 
grant select, insert, update, delete on am.ladon_role_member to orgservice; 
grant execute on function am.delete_org to orgservice;
grant select on am.organization_status to orgservice;
grant select on am.subscription_types to orgservice;

-- userservice permissions
grant linkai_user to userservice;
grant select, insert, update, delete on am.users to userservice;
grant select on am.user_status to userservice;
grant references (organization_id) on table am.organizations to userservice;

-- scangroupservice permissions
grant linkai_user to scangroupservice;
grant select, insert, update, delete on am.scan_group to scangroupservice;
grant references (organization_id) on table am.organizations to scangroupservice;
grant references (user_id) on table am.users to scangroupservice;

-- addressservice permissions
grant linkai_user to addressservice;
grant select, insert, update, delete on am.scan_group_addresses to addressservice;
grant select on am.scan_address_discovered_by to addressservice;
grant references (organization_id) on table am.organizations to addressservice;
grant references (user_id) on table am.users to addressservice;
grant references (scan_group_id) on table am.scan_group to addressservice;

-- findingsservice permissions
grant linkai_user to findingsservice;
grant select, insert, update, delete on am.scan_group_findings to findingsservice;
grant select on am.scan_finding_types to findingsservice;
grant references (organization_id) on table am.organizations to findingsservice;
grant references (user_id) on table am.users to findingsservice;
grant references (scan_group_id) on table am.scan_group to findingsservice;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

-- findingsservice permissions
revoke linkai_user from findingsservice;
revoke select, insert, update, delete on am.scan_group_findings from findingsservice;
revoke select on am.scan_finding_types from findingsservice;
revoke references (organization_id) on table am.organizations from findingsservice;
revoke references (user_id) on table am.users from findingsservice;
revoke references (scan_group_id) on table am.scan_group from findingsservice;

-- addressservice permissions
revoke linkai_user from addressservice;
revoke select, insert, update, delete on am.scan_group_addresses from addressservice;
revoke select on am.scan_address_discovered_by from addressservice;
revoke references (organization_id) on table am.organizations from addressservice;
revoke references (user_id) on table am.users from addressservice;
revoke references (scan_group_id) on table am.scan_group from addressservice;

-- scangroupservice permissions
revoke linkai_user from scangroupservice;
revoke select, insert, update, delete on am.scan_group from scangroupservice;
revoke references (organization_id) on table am.organizations from scangroupservice;
revoke references (user_id) on table am.users from scangroupservice;

-- userservice permissions
revoke linkai_user from userservice;
revoke select, insert, update, delete on am.users from userservice;
revoke select on am.user_status from userservice;
revoke references (organization_id) on table am.organizations from userservice;

-- orgservice permissions
revoke linkai_user from orgservice;
revoke select, insert, update, delete on am.organizations from orgservice; 
revoke select, insert, update, delete on am.users from orgservice; 
revoke select, insert, update, delete on am.ladon_role from orgservice; 
revoke select, insert, update, delete on am.ladon_role_member from orgservice; 
revoke execute on function am.delete_org from orgservice;
revoke select on am.organization_status from orgservice;
revoke select on am.subscription_types from orgservice;

-- eventservice permissions
revoke linkai_user from eventservice;
revoke select, insert, update, delete on am.scan_group_events from eventservice;
revoke references (organization_id) on table am.organizations from eventservice;
revoke references (user_id) on table am.users from eventservice;
revoke references (scan_group_id) on table am.scan_group from eventservice;

-- linkai_user role
revoke select on am.ladon_role from linkai_user;
revoke select on am.ladon_role_member from linkai_user;
revoke select on am.ladon_policy from linkai_user;
revoke select on am.ladon_policy_subject from linkai_user;
revoke select on am.ladon_policy_permission from linkai_user;
revoke select on am.ladon_policy_resource from linkai_user;
revoke select on am.ladon_subject from linkai_user;
revoke select on am.ladon_action from linkai_user;
revoke select on am.ladon_resource from linkai_user;
revoke select on am.ladon_policy_subject_rel from linkai_user;
revoke select on am.ladon_policy_action_rel from linkai_user;
revoke select on am.ladon_policy_resource_rel from linkai_user;
revoke usage on schema am from linkai_user;
revoke select,update on all sequences in schema am from linkai_user;
drop role linkai_user;
