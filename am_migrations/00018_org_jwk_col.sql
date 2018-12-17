-- +goose Up
-- SQL in this section is executed when the migration is applied.
alter table am.organizations add column user_pool_jwk text default '' not null;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
alter table am.organizations drop column user_pool_jwk;
