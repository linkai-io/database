-- +goose Up
-- SQL in this section is executed when the migration is applied.
alter table am.web_responses add column url_request_timestamp bigint;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
alter table am.web_responses drop column url_request_timestamp;
