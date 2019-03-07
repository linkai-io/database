-- +goose Up
-- SQL in this section is executed when the migration is applied.
create table am.web_techtypes (
    techtype_id serial primary key,
    techname varchar(512) not null default '',
    category_id int not null,
    category varchar(512) not null,
    icon varchar(128) not null default '',
    website varchar(1024) not null default '',
    ignored boolean default false,
    unique(techname, category)
);

create table am.web_technologies (
    tech_id bigserial not null primary key,
    snapshot_id bigint REFERENCES am.web_snapshots (snapshot_id),
    category_id int references am.web_techtypes (techtype_id),
    matched_text text not null default '',
    match_location text not null default '',
    version varchar(512) not null default ''
);

grant select, insert, update, delete on am.web_technologies to webdataservice;
grant select on am.web_techtypes to webdataservice;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
revoke select, insert, update, delete on am.web_technologies from webdataservice;
revoke select on am.web_techtypes from webdataservice;
drop table am.web_technologies;
drop table am.web_techtypes;
