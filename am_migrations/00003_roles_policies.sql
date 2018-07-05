-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS am.ladon_policy (
    id           varchar(255) NOT NULL PRIMARY KEY,
    description  text NOT NULL,
    effect       text NOT NULL CHECK (effect='allow' OR effect='deny'),
    conditions	 text NOT NULL
);

CREATE TABLE IF NOT EXISTS am.ladon_policy_subject (
    compiled text NOT NULL,
    template varchar(1023) NOT NULL,
    policy   varchar(255) NOT NULL,
    FOREIGN KEY (policy) REFERENCES am.ladon_policy(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS am.ladon_policy_permission (
    compiled text NOT NULL,
    template varchar(1023) NOT NULL,
    policy   varchar(255) NOT NULL,
    FOREIGN KEY (policy) REFERENCES am.ladon_policy(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS am.ladon_policy_resource (
    compiled text NOT NULL,
    template varchar(1023) NOT NULL,
    policy   varchar(255) NOT NULL,
    FOREIGN KEY (policy) REFERENCES am.ladon_policy(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS am.ladon_subject (
    id          varchar(64) NOT NULL PRIMARY KEY,
    has_regex   bool NOT NULL,
    compiled    varchar(511) NOT NULL UNIQUE,
    template    varchar(511) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS am.ladon_action (
    id          varchar(64) NOT NULL PRIMARY KEY,
    has_regex   bool NOT NULL,
    compiled    varchar(511) NOT NULL UNIQUE,
    template    varchar(511) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS am.ladon_resource (
    id          varchar(64) NOT NULL PRIMARY KEY,
    has_regex   bool NOT NULL,
    compiled    varchar(511) NOT NULL UNIQUE,
    template    varchar(511) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS am.ladon_policy_subject_rel (
    policy   varchar(255) NOT NULL,
    subject  varchar(64) NOT NULL,
    PRIMARY KEY (policy, subject),
    FOREIGN KEY (policy) REFERENCES am.ladon_policy(id) ON DELETE CASCADE,
    FOREIGN KEY (subject) REFERENCES am.ladon_subject(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS am.ladon_policy_action_rel (
    policy  varchar(255) NOT NULL,
    action  varchar(64) NOT NULL,
    PRIMARY KEY (policy, action),
    FOREIGN KEY (policy) REFERENCES am.ladon_policy(id) ON DELETE CASCADE,
    FOREIGN KEY (action) REFERENCES am.ladon_action(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS am.ladon_policy_resource_rel (
    policy    varchar(255) NOT NULL,
    resource  varchar(64) NOT NULL,
    PRIMARY KEY (policy, resource),
    FOREIGN KEY (policy) REFERENCES am.ladon_policy(id) ON DELETE CASCADE,
    FOREIGN KEY (resource) REFERENCES am.ladon_resource(id) ON DELETE CASCADE
);

CREATE INDEX ladon_subject_compiled_idx ON am.ladon_subject (compiled text_pattern_ops);

CREATE INDEX ladon_permission_compiled_idx ON am.ladon_action (compiled text_pattern_ops);

CREATE INDEX ladon_resource_compiled_idx ON am.ladon_resource (compiled text_pattern_ops);

ALTER TABLE am.ladon_policy ADD COLUMN meta json;

CREATE TABLE IF NOT EXISTS am.ladon_role (
	role_id varchar(255) NOT NULL PRIMARY KEY,
    organization_id integer REFERENCES am.organizations (organization_id) not null
);

CREATE TABLE IF NOT EXISTS am.ladon_role_member (
    organization_id integer REFERENCES am.organizations (organization_id),
    member_id integer REFERENCES am.users (user_id) not null,
	role_id		varchar(255) NOT NULL,
	FOREIGN KEY (role_id) REFERENCES am.ladon_role(role_id) ON DELETE CASCADE,
	PRIMARY KEY (member_id, role_id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE am.ladon_role_member;
DROP TABLE am.ladon_role;
ALTER TABLE am.ladon_policy DROP COLUMN IF EXISTS meta;
DROP INDEX am.ladon_resource_compiled_idx;
DROP INDEX am.ladon_permission_compiled_idx;
DROP INDEX am.ladon_subject_compiled_idx;
DROP TABLE am.ladon_policy_resource_rel;
DROP TABLE am.ladon_policy_action_rel;
DROP TABLE am.ladon_policy_subject_rel;
DROP TABLE am.ladon_resource;
DROP TABLE am.ladon_action;
DROP TABLE am.ladon_subject;
DROP TABLE am.ladon_policy_resource;
DROP TABLE am.ladon_policy_permission;
DROP TABLE am.ladon_policy_subject;
DROP TABLE am.ladon_policy;
