-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE SCHEMA am;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP SCHEMA am;