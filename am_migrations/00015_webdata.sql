-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE am.web_snapshots (
    snapshot_id bigserial not null primary key,
    organization_id int references am.organizations (organization_id),
    scan_group_id int references am.scan_group (scan_group_id),
    address_id bigint references am.scan_group_addresses (address_id),
    response_timestamp bigint,
    serialized_dom_hash varchar(512),
    serialized_dom_link text,
    snapshot_link text,
    is_deleted boolean
);

CREATE TABLE am.web_status_text (
    status_text_id serial not null primary key,
    status_text text not null
);

CREATE TABLE am.web_mime_type (
    mime_type_id serial not null primary key,
    mime_type text not null
);

CREATE TABLE am.web_responses (
    response_id bigserial not null primary key,
    organization_id int references am.organizations (organization_id),
    scan_group_id int references am.scan_group (scan_group_id),
    address_id bigint references am.scan_group_addresses (address_id),
    response_timestamp bigint,
    is_document boolean,
    scheme varchar(12),
    ip_address varchar(256),
    host_address varchar(512),
    response_port int,
    requested_port int,
    url bytea not null,
    headers jsonb,
    status int, 
    status_text_id int references am.web_status_text (status_text_id),
    mime_type_id int references am.web_mime_type (mime_type_id),
    raw_body_hash varchar(512),
    raw_body_link text,
    is_deleted boolean
);

CREATE TABLE am.web_certificates (
    certificate_id bigserial not null primary key,
    organization_id int references am.organizations (organization_id),
    scan_group_id int references am.scan_group (scan_group_id),
    response_timestamp bigint,
    host varchar(512),
    port int,
    protocol text,
    key_exchange text,
    key_exchange_group text,
    cipher text,
    mac text,
    certificate_value int,
    subject_name text,
    san_list jsonb,
    issuer text,
    valid_from bigint,
    valid_to bigint,
    ct_compliance text,
    is_deleted boolean
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE am.web_certificates;
DROP TABLE am.web_responses;
DROP TABLE am.web_mime_type;
DROP TABLE am.web_status_text;
DROP TABLE am.web_snapshots;