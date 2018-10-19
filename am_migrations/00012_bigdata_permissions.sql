-- +goose Up
-- SQL in this section is executed when the migration is applied.
grant linkai_user to bigdata_admin;

grant linkai_user to bigdata_reader;
grant select on am.certificate_servers to bigdata_reader;
grant select on am.certificate_queries to bigdata_reader;
grant select on am.certificates to bigdata_reader;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
revoke linkai_user from bigdata_admin;

revoke linkai_user from bigdata_reader;
revoke select on am.certificate_servers from bigdata_reader;
revoke select on am.certificate_queries from bigdata_reader;
revoke select on am.certificates from bigdata_reader;