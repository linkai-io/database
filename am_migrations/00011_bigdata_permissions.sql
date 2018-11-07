-- +goose Up
-- SQL in this section is executed when the migration is applied.
grant linkai_user to bigdataservice;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
revoke linkai_user from bigdataservice;