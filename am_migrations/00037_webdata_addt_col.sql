-- +goose Up
-- SQL in this section is executed when the migration is applied.
alter table only am.web_snapshots add column load_url bytea not null default '';
alter table only am.web_snapshots add column url_request_timestamp timestamptz not null default 'epoch';
create index web_snapshots_url_request_timestamp_index on am.web_snapshots (url_request_timestamp desc);
update am.web_snapshots set url_request_timestamp=response_timestamp;

-- add index to web responses as well
create index web_responses_url_request_timestamp_index on am.web_responses (url_request_timestamp desc);
create index web_responses_response_id_desc_index on am.web_responses (response_id desc);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
alter table only am.web_snapshots drop column load_url;
drop index am.web_snapshots_url_request_timestamp_index;
alter table only am.web_snapshots drop column url_request_timestamp;
drop index am.web_responses_url_request_timestamp_index;