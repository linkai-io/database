-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE SCHEMA linkai;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP SCHEMA linkai;