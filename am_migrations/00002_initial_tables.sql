-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE am.subscription_types (
    subscription_id serial not null primary key,
    title varchar(128) not null unique
);

INSERT INTO am.subscription_types (title) values 
    ('one time'),
    ('monthly'),
    ('enterprise');

CREATE TABLE am.organizations (
    organization_id serial not null primary key,
    organization_name varchar(256) not null unique,
    owner_email varchar(256) not null,
    first_name varchar(256) not null,
    last_name varchar(256) not null,
    phone varchar(32) not null,
    country varchar(128) not null,
    state_prefecture varchar(256) not null,
    street varchar(256) not null, 
    address1 varchar(256),
    address2 varchar(256),
    city varchar(256) not null,
    postal_code varchar(32) not null,
    creation_time bigint not null,
    deleted boolean,
    subscription_id integer REFERENCES am.subscription_types (subscription_id)
);

CREATE UNIQUE INDEX idx_lower_organizations_organization_name ON am.organizations (lower(organization_name));
CREATE UNIQUE INDEX idx_lower_organizations_owner_email ON am.organizations (lower(owner_email));

CREATE TABLE am.users (
    user_id serial not null primary key,
    organization_id integer REFERENCES am.organizations (organization_id),
    email varchar(256) not null,
    first_name varchar(256) not null,
    last_name varchar(256) not null,
    deleted boolean,
    UNIQUE (organization_id, email)
);

CREATE UNIQUE INDEX idx_lower_users_email ON am.users (lower(email));

CREATE TABLE am.scan_group (
    scan_group_id serial not null primary key,
    organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_name varchar(256) not null,
    creation_time bigint not null,
    created_by integer REFERENCES am.users (user_id),
    modified_time bigint not null,
    modified_by integer REFERENCES am.users (user_id),
    original_input bytea not null,
    configuration jsonb,
    deleted boolean,
    UNIQUE (organization_id, scan_group_name)
);

CREATE TABLE am.scan_address_configuration (
    scan_address_configuration_id serial not null primary key,
    organization_id integer REFERENCES am.organizations (organization_id),
    configuration_name varchar(256) not null,
    configuration jsonb,
    UNIQUE (organization_id, configuration_name)
);

CREATE TABLE am.scan_group_addresses (
    address_id bigserial not null primary key,
    organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_id integer REFERENCES am.scan_group (scan_group_id),
    address varchar(512) not null,
    configuration_id integer REFERENCES am.scan_address_configuration (scan_address_configuration_id),
    added_timestamp bigint,
    added_by varchar(128) not null,
    deleted boolean,
    ignored boolean
);

CREATE TABLE am.scan_group_address_map (
    address_map_id bigserial not null primary key,
    organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_id integer REFERENCES am.scan_group (scan_group_id),
    hostname varchar(512),
    ipv4 varchar(64),
    ipv6 varchar(128),
    deleted boolean
);

CREATE TABLE am.jobs (
    job_id bigserial not null primary key,
    organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_id integer REFERENCES am.scan_group (scan_group_id)
);

CREATE TABLE am.job_events (
    event_id bigserial not null primary key,
    organization_id integer REFERENCES am.organizations (organization_id),
    job_id bigint REFERENCES am.jobs (job_id),
    event_time bigint,
    event_description text,
    event_from varchar(256) 
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE am.job_events;
DROP TABLE am.jobs;
DROP TABLE am.scan_group_address_map;
DROP TABLE am.scan_address_configuration;
DROP TABLE am.scan_group_addresses;
DROP TABLE am.scan_group;
DROP INDEX am.idx_lower_users_email;
DROP TABLE am.users;
DROP INDEX am.idx_lower_organizations_organization_name;
DROP INDEX am.idx_lower_organizations_owner_email;
DROP TABLE am.organizations;
DROP TABLE am.subscription_types;