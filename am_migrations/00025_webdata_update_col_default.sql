-- +goose Up
-- SQL in this section is executed when the migration is applied.
alter table ONLY am.web_responses alter column url_request_timestamp set default 0;
update am.web_responses set url_request_timestamp = 0 where url_request_timestamp is null;
alter table ONLY am.web_responses alter column url_request_timestamp set not null;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
alter table am.web_responses alter column url_request_timestamp drop default;
