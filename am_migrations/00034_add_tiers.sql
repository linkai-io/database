-- +goose Up
-- SQL in this section is executed when the migration is applied.
insert into am.subscription_types (subscription_id, title) values
    (101, 'monthly_smb'),
    (102, 'monthly_medium');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
delete from am.subscription_types where subscription_id=101 or subscription_id=102;