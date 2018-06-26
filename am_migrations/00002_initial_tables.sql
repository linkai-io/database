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

CREATE TABLE am.roles (
    role_id serial not null primary key,
    role_name varchar(128) not null unique 
);

INSERT INTO am.roles (role_name) values 
    ('Owner'),
    ('Administrator'),
    ('Auditor'),
    ('Editor'),
    ('Reviewer');

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
    creation_time integer not null,
    subscription_id integer REFERENCES am.subscription_types (subscription_id)
);

CREATE UNIQUE INDEX idx_lower_organizations_owner_email ON am.organizations (lower(owner_email));

CREATE TABLE am.users (
    user_id serial not null primary key,
    organization_id integer REFERENCES am.organizations (organization_id),
    email varchar(256) not null,
    UNIQUE (organization_id, email),
    role_id integer REFERENCES am.roles (role_id)
);

CREATE UNIQUE INDEX idx_lower_users_email ON am.users (lower(email));

CREATE TABLE am.user_groups (
    user_group_id serial not null primary key,
    organization_id integer REFERENCES am.organizations (organization_id),
    user_group_name varchar(256) not null,
    UNIQUE (organization_id, user_group_name)
);

CREATE TABLE am.user_group_members (
    user_group_id integer REFERENCES am.user_groups (user_group_id),
    organization_id integer REFERENCES am.organizations (organization_id),
    member_id integer REFERENCES am.users (user_id)
);

CREATE TABLE am.scan_group (
    scan_group_id serial not null primary key,
    organization_id integer REFERENCES am.organizations (organization_id),
    creation_time integer not null,
    created_by integer REFERENCES am.users (user_id),
    raw_input bytea not null,
    parsed_input bytea not null,
    configuration jsonb,
    config_version integer not null,
    version_name varchar(128) not null
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
    event_time integer,
    event_description text,
    event_from varchar(256) 
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE am.job_events;
DROP TABLE am.jobs;
DROP TABLE am.scan_group;
DROP TABLE am.user_group_members;
DROP TABLE am.user_groups;
DROP INDEX am.idx_lower_users_email;
DROP TABLE am.users;
DROP INDEX am.idx_lower_organizations_owner_email;
DROP TABLE am.organizations;
DROP TABLE am.roles;
DROP TABLE am.subscription_types;