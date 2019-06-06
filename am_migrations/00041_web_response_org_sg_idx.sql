-- +goose Up
-- SQL in this section is executed when the migration is applied.
create index url_org_sg_responses_idx on am.web_responses (url_request_timestamp, organization_id, scan_group_id);
create index url_org_sg_snapshots_idx on am.web_snapshots (url_request_timestamp, organization_id, scan_group_id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop index am.url_org_sg_responses_idx;
drop index am.url_org_sg_snapshots_idx;
