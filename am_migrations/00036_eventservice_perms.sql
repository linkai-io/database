-- +goose Up
-- SQL in this section is executed when the migration is applied.
grant select on am.scan_group to eventservice;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
revoke select on am.scan_group from eventservice;
