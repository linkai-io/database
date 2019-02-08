-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE am.certificate_queries_subdomains (
    etld_id serial not null primary key,
    etld varchar(512) not null unique,
    query_timestamp bigint
);

CREATE TABLE am.certificate_subdomains (
    subdomain_id bigserial not null primary key,
    etld_id int references am.certificate_queries_subdomains (etld_id) on delete cascade,
    inserted_timestamp bigint,
    common_name text
);

grant select, insert, update, delete on am.certificate_queries_subdomains to bigdataservice;
grant select, insert, update, delete on am.certificate_subdomains to bigdataservice;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
revoke select, insert, update, delete on am.certificate_queries_subdomains from bigdataservice;
revoke select, insert, update, delete on am.certificate_subdomains from bigdataservice;
DROP TABLE am.certificate_subdomains;
DROP TABLE am.certificate_queries_subdomains;
