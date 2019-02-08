-- +goose Up
-- SQL in this section is executed when the migration is applied.
grant select, references (user_id, email) on table am.users to scangroupservice;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
-- 00008_permissions.sql has references allowed already so don't break during roll back
revoke select (user_id) on table am.users from scangroupservice;
revoke select, references (email) on table am.users from scangroupservice;