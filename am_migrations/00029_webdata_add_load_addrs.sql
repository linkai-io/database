-- +goose Up
-- SQL in this section is executed when the migration is applied.
alter table ONLY am.web_responses 
    add column load_host_address varchar(512) not null default '',
    add column load_ip_address varchar(256) not null default '';


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
alter table only am.web_responses drop column load_host_address;
alter table only am.web_responses drop column load_ip_address;
