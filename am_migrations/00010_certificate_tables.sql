-- +goose Up
-- SQL in this section is executed when the migration is applied.
INSERT INTO am.scan_address_discovered_by (discovery_id, discovered_by) values 
    (400, 'bigdata'),
    (401, 'bigdata_certificate_transparency');

CREATE TABLE am.certificate_queries (
    id serial not null primary key,
    etld varchar(512) not null unique,
    query_timestamp bigint
);

CREATE TABLE am.certificates (
    certificate_id bigserial not null primary key,
    inserted_timestamp bigint,
    server_name text not null,
    server_index bigint,
    etld varchar(512) not null,
    cert_hash varchar(256) not null unique,
    serial_number varchar(256),
    not_before bigint,
    not_after bigint,
    country varchar(256),
    organization text,
    organizational_unit text,
    common_name text,
    verified_dns_names text,
    unverified_dns_names text,
    ip_addresses text, 
    email_addresses text
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE am.certificate_queries;
DROP TABLE am.certificates;
DELETE FROM am.scan_address_discovered_by where discovery_id=400;
DELETE FROM am.scan_address_discovered_by where discovery_id=401;