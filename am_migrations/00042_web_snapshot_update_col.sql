-- +goose Up
-- SQL in this section is executed when the migration is applied.
alter table only am.web_snapshots add column previous_url_timestamp timestamptz not null default 'epoch';
alter table only am.web_snapshots_archive add column previous_url_timestamp timestamptz not null default 'epoch';

alter table only am.web_snapshots add column updated boolean default true;
alter table only am.web_snapshots_archive add column updated boolean default true;

alter table only am.web_technologies add column updated boolean default true;
alter table only am.web_technologies_archive add column updated boolean default true;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
alter table only am.web_technologies drop column updated;
alter table only am.web_technologies_archive drop column updated;

alter table only am.web_snapshots drop column updated;
alter table only am.web_snapshots_archive drop column updated;

alter table only am.web_snapshots drop column previous_url_timestamp;
alter table only am.web_snapshots_archive drop column previous_url_timestamp;

