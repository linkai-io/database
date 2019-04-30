-- +goose Up
-- SQL in this section is executed when the migration is applied.
alter table only am.users add column agreement_accepted boolean default false;
alter table only am.users add column agreement_accepted_timestamp timestamptz not null default 'epoch';

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
alter table only am.users drop column agreement_accepted;
alter table only am.users drop column agreement_accepted_timestamp;
