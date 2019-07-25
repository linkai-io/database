-- +goose Up
-- SQL in this section is executed when the migration is applied.

alter table only am.organizations add column payment_required_timestamp timestamptz not null default 'epoch';
alter table only am.organizations add column stripe_plan_id text not null default '';
alter table only am.organizations add column stripe_subscription_id text not null default '';
alter table only am.organizations add column is_beta_plan boolean default true;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
alter table only am.organizations drop column payment_required_timestamp;
alter table only am.organizations drop column stripe_plan_id;
alter table only am.organizations drop column stripe_subscription_id;
alter table only am.organizations drop column is_beta_plan;