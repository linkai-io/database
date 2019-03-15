-- +goose Up
-- SQL in this section is executed when the migration is applied.
alter table am.scan_group_addresses drop constraint scan_group_addresses_scan_group_id_host_address_ip_address_key;
alter table am.scan_group_addresses add constraint unique_address_hash unique (scan_group_id, address_hash);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
alter table am.scan_group_addresses drop constraint unique_address_hash;
alter table am.scan_group_addresses add constraint scan_group_addresses_scan_group_id_host_address_ip_address_key unique (scan_group_id, host_address, ip_address);