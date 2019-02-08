-- +goose Up
-- SQL in this section is executed when the migration is applied.
grant select, references (address_id, host_address, ip_address) on table am.scan_group_addresses to webdataservice;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
-- 000014_webdata_permissions.sql has references allowed already so don't break during roll back
revoke select (address_id) on table am.scan_group_addresses from webdataservice;
revoke select, references (host_address, ip_address) on table am.scan_group_addresses from webdataservice;