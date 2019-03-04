-- +goose Up
-- SQL in this section is executed when the migration is applied.
INSERT INTO am.scan_address_discovered_by (discovery_id, discovered_by) values (104, 'ns_query_nsec_walk');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DELETE FROM am.scan_address_discovered_by where discovery_id = 104;
