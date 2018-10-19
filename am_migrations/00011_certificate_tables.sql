-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE am.certificate_servers (
    id serial not null primary key,
    url text not null unique,
    index bigint,
    index_updated bigint,
    step integer,
    tree_size bigint,
    tree_size_updated bigint
);

ALTER TABLE am.certificate_servers owner to bigdata_admin;

CREATE TABLE am.certificate_queries (
    id serial not null primary key,
    etld varchar(512),
    query_timestamp bigint
);

ALTER TABLE am.certificate_queries owner to bigdata_admin;

CREATE TABLE am.certificates (
    certificate_id bigserial not null primary key,
    inserted_timestamp bigint,
    etld varchar(512),
    cert_hash varchar(256) not null unique,
    serial_number varchar(256),
    not_before timestamptz,
    not_after timestamptz,
    country varchar(256),
    organization text,
    organizational_unit text,
    common_name text,
    verified_dns_names jsonb,
    unverified_dns_names jsonb,
    ip_addresses jsonb, 
    email_addresses jsonb
);

ALTER TABLE am.certificates owner to bigdata_admin;


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE am.certificate_queries;
DROP TABLE am.certificate_servers;
DROP TABLE am.certificates;