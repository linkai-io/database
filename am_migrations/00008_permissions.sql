-- +goose Up
-- SQL in this section is executed when the migration is applied.
revoke all on SCHEMA am FROM PUBLIC;
revoke all on schema public from public;

-- allow all services to access the roles/perm/policy tables.
-- allows usage on am schema, but disable creating tables.
create role linkai_user with inherit;
grant usage on schema am to linkai_user;
grant select,update on all sequences in schema am to linkai_user;
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


-- jobservice permissions
grant linkai_user to jobservice;
revoke all on schema am from jobservice;
grant select, insert, update, delete on am.jobs to jobservice; 
grant select, insert, update, delete on am.job_events to jobservice;
grant references (organization_id) on table am.organizations to jobservice;
grant references (user_id) on table am.users to jobservice;
grant references (scan_group_id) on table am.scan_group to jobservice;


-- orgservice permissions
grant linkai_user to orgservice;
grant select, insert, update, delete on am.organizations to orgservice; 
grant select, insert, update, delete on am.users to orgservice; 
grant select, insert, update, delete on am.ladon_role to orgservice; 
grant select, insert, update, delete on am.ladon_role_member to orgservice; 
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
grant select, insert, update, delete on am.scan_address_configuration to scangroupservice;
grant select, insert, update, delete on am.scan_group_addresses to scangroupservice;
grant select, insert, update, delete on am.scan_group_address_map to scangroupservice;
grant select on am.scan_address_added_by to scangroupservice;
grant references (organization_id) on table am.organizations to scangroupservice;
grant references (user_id) on table am.users to scangroupservice;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
revoke linkai_user from scangroupservice;
revoke select, insert, update, delete on am.scan_group from scangroupservice;
revoke select, insert, update, delete on am.scan_address_configuration from scangroupservice;
revoke select, insert, update, delete on am.scan_group_addresses from scangroupservice;
revoke select, insert, update, delete on am.scan_group_address_map from scangroupservice;
revoke select on am.scan_address_added_by from scangroupservice;
revoke references (organization_id) on table am.organizations from scangroupservice;
revoke references (user_id) on table am.users from scangroupservice;

revoke linkai_user from userservice;
revoke select, insert, update, delete on am.users from userservice;
revoke select on am.user_status from userservice;
revoke references (organization_id) on table am.organizations from userservice;

revoke linkai_user from orgservice;
revoke select, insert, update, delete on am.organizations from orgservice; 
revoke select, insert, update, delete on am.users from orgservice; 
revoke select, insert, update, delete on am.ladon_role from orgservice; 
revoke select, insert, update, delete on am.ladon_role_member from orgservice; 
revoke select on am.organization_status from orgservice;
revoke select on am.subscription_types from orgservice;

revoke linkai_user from jobservice;
revoke select, insert, update, delete on am.jobs from jobservice; 
revoke select, insert, update, delete on am.job_events from jobservice;
revoke references (organization_id) on table am.organizations from jobservice;
revoke references (user_id) on table am.users from jobservice;
revoke references (scan_group_id) on table am.scan_group from jobservice;

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
