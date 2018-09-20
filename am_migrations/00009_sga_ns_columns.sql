-- +goose Up
-- SQL in this section is executed when the migration is applied.
alter table am.scan_group_addresses
    add column found_from bigint,
    add column ns_record int,
    add column address_hash varchar(128);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
alter table am.scan_group_addresses 
    drop column found_from,
    drop column ns_record,
    drop column address_hash;