-- +goose Up
-- SQL in this section is executed when the migration is applied.
grant linkai_user to bigdataservice;
grant usage on schema am to bigdataservice;
grant usage, select, update on all sequences in schema am to bigdataservice;
grant select, insert, update, delete on am.certificate_queries to bigdataservice;
grant select, insert, update, delete on am.certificates to bigdataservice;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
revoke linkai_user from bigdataservice;
revoke usage, select on all sequences in schema am from bigdataservice;
revoke select, insert, update, delete on am.certificate_queries from bigdataservice;
revoke select, insert, update, delete on am.certificates from bigdataservice;