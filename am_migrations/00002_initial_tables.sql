-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE DOMAIN required_text as varchar(256) not null check (length(value) > 0);

CREATE TABLE am.subscription_types (
    subscription_id integer not null primary key,
    title varchar(128) not null unique
);

INSERT INTO am.subscription_types (subscription_id, title) values 
    (1, 'pending'),
    (10, 'one time'),
    (100, 'monthly'),
    (1000, 'enterprise'),
    (9999, 'system');

CREATE TABLE am.organization_status (
    status_id integer not null primary key,
    description required_text unique
);

INSERT INTO am.organization_status (status_id, description) values 
    -- Disabled reasons 1-99
    (1, 'Disabled - Pending Payment'),
    (2, 'Disabled - Closed'),
    (3, 'Disabled - Locked'),
    -- Not Enabled reasons 100 - 999
    (100, 'Awaiting Activation'),
    -- Enabled reasons 1000 - ...
    (1000, 'Active');

CREATE TABLE am.organizations (
    organization_id serial not null primary key,
    organization_name required_text unique,
    organization_custom_id required_text unique,
    user_pool_id required_text,
    identity_pool_id required_text,
    owner_email required_text,
    first_name required_text,
    last_name required_text,
    phone required_text,
    country required_text,
    state_prefecture required_text,
    street required_text, 
    address1 varchar(256) not null,
    address2 varchar(256) not null,
    city required_text,
    postal_code required_text,
    creation_time bigint not null,
    status_id integer REFERENCES am.organization_status (status_id),
    deleted boolean not null,
    subscription_id integer REFERENCES am.subscription_types (subscription_id)
);

CREATE UNIQUE INDEX idx_lower_organizations_organization_name ON am.organizations (lower(organization_name));
CREATE UNIQUE INDEX idx_lower_organizations_owner_email ON am.organizations (lower(owner_email));

CREATE TABLE am.user_status (
    status_id integer not null primary key,
    description required_text unique
);

INSERT INTO am.user_status (status_id, description) values 
    -- Disabled reasons 1-99
    (1, 'Disabled - Locked'),
    -- Not Enabled reasons 100 - 999
    (100, 'Awaiting Activation'),
    -- Enabled reasons 1000 - ...
    (1000, 'Active');

CREATE TABLE am.users (
    user_id serial not null primary key,
    organization_id integer REFERENCES am.organizations (organization_id),
    user_custom_id required_text unique,
    email required_text,
    first_name required_text,
    last_name required_text,
    user_status_id integer REFERENCES am.user_status (status_id),
    creation_time bigint not null,
    deleted boolean not null,
    UNIQUE (organization_id, email)
);

CREATE UNIQUE INDEX idx_lower_users_email ON am.users (lower(email));

CREATE TABLE am.scan_group (
    scan_group_id serial not null primary key,
    organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_name required_text,
    creation_time bigint not null,
    created_by integer REFERENCES am.users (user_id),
    modified_time bigint not null,
    modified_by integer REFERENCES am.users (user_id),
    original_input_s3_url required_text,
    configuration jsonb,
    deleted boolean not null,
    paused boolean not null,
    UNIQUE (organization_id, scan_group_name)
);

CREATE TABLE am.scan_address_discovered_by (
    discovery_id integer not null primary key,
    discovered_by required_text
);

INSERT INTO am.scan_address_discovered_by (discovery_id, discovered_by) values 
    -- manual addition of addresses
    (1, 'input_list'),
    (2, 'manual'),
    (3, 'other'),
    -- ns analyzer module 100-200
    (100, 'ns_query_other'),
    (101, 'ns_query_ip_to_name'),
    (102, 'ns_query_name_to_ip'),
    -- dns brute module 200-300
    (200, 'dns_brute_forcer'),
    (201, 'dns_axfr'),
    -- web modules 300 - 999
    (300, 'web_crawler'),
    -- other, feature modules
    (1000, 'git_hooks');

CREATE TABLE am.scan_group_events (
    event_id bigserial not null primary key,
    organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_id bigint REFERENCES am.scan_group (scan_group_id),
    event_user_id integer references am.users (user_id), 
    event_time bigint,
    event_description text,
    event_from required_text
);

CREATE TABLE am.scan_group_addresses (
    address_id bigserial not null primary key,
    organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_id integer REFERENCES am.scan_group (scan_group_id),
    host_address varchar(512),
    ip_address varchar(256),
    discovered_timestamp bigint,
    discovery_id integer REFERENCES am.scan_address_discovered_by (discovery_id),
    last_scanned_timestamp bigint,
    last_seen_timestamp bigint,
    confidence_score float,
    user_confidence_score float,
    is_soa boolean not null,
    is_wildcard_zone boolean not null,
    is_hosted_service boolean not null,
    ignored boolean not null,
    check (host_address is not null or ip_address is not null),
    UNIQUE(scan_group_id, host_address, ip_address)
);

CREATE INDEX idx_scan_group_addresses_address_id ON am.scan_group_addresses (organization_id,scan_group_id,address_id);

CREATE TABLE am.scan_finding_types (
    finding_id integer not null primary key,
    description required_text
);

INSERT INTO am.scan_finding_types (finding_id, description) values 
    (1, 'all_ns_records'),
    (2, 'axfr_results'),
    (3, 'other');

CREATE TABLE am.scan_group_findings (
    organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_id integer REFERENCES am.scan_group (scan_group_id),
    address_id integer REFERENCES am.scan_group_addresses (address_id) on delete cascade,
    finding_id integer REFERENCES am.scan_finding_types (finding_id),
    discovered_timestamp bigint not null,
    last_seen_timestamp bigint not null,
    custom_description text,
    data jsonb
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE am.scan_group_findings;
DROP TABLE am.scan_finding_types;
DROP INDEX am.idx_scan_group_addresses_address_id;
DROP TABLE am.scan_group_addresses;
DROP TABLE am.scan_group_events;
DROP TABLE am.scan_address_discovered_by;
DROP TABLE am.scan_group;
DROP INDEX am.idx_lower_users_email;
DROP TABLE am.users;
DROP TABLE am.user_status;
DROP INDEX am.idx_lower_organizations_organization_name;
DROP INDEX am.idx_lower_organizations_owner_email;
DROP TABLE am.organizations;
DROP TABLE am.organization_status;
DROP TABLE am.subscription_types;
DROP DOMAIN required_text;