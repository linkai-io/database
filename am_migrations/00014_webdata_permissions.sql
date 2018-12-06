-- +goose Up
-- SQL in this section is executed when the migration is applied.
grant linkai_user to webdataservice;
grant usage on schema am to webdataservice;
grant usage, select, update on all sequences in schema am to webdataservice;
grant references (organization_id) on table am.organizations to webdataservice;
grant references (user_id) on table am.users to webdataservice;
grant references (scan_group_id) on table am.scan_group to webdataservice;
grant select, insert, update, delete on am.web_snapshots to webdataservice;
grant select, insert, update, delete on am.web_status_text to webdataservice;
grant select, insert, update, delete on am.web_mime_type to webdataservice;
grant select, insert, update, delete on am.web_responses to webdataservice;
grant select, insert, update, delete on am.web_certificates to webdataservice;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
revoke linkai_user from webdataservice;
revoke usage in schema am from webdataservice;
revoke usage, select on all sequences in schema am from webdataservice;
revoke references (organization_id) on table am.organizations from webdataservice;
revoke references (user_id) on table am.users from webdataservice;
revoke references (scan_group_id) on table am.scan_group from webdataservice;
revoke select, insert, update, delete on am.web_snapshots from webdataservice;
revoke select, insert, update, delete on am.web_status_text from webdataservice;
revoke select, insert, update, delete on am.web_mime_type from webdataservice;
revoke select, insert, update, delete on am.web_responses from webdataservice;
revoke select, insert, update, delete on am.web_certificates from webdataservice;