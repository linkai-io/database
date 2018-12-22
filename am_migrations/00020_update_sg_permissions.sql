-- +goose Up
-- SQL in this section is executed when the migration is applied.
grant select, references (user_id, email) on table am.users to scangroupservice;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
revoke select, references (user_id, email) on table am.users from scangroupservice;