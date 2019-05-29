-- +goose Up
-- SQL in this section is executed when the migration is applied.
alter table only am.scan_group add column last_paused_timestamp timestamptz not null default 'epoch';

create table am.scan_group_addresses_ports (
	organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_id integer REFERENCES am.scan_group (scan_group_id),
    host_address varchar(512) not null default '',
    ip_address varchar(256) not null default '',
	address_hash varchar(128) not null default '',
	port integer not null default 0,
	protocol integer not null default 0, -- 0 tcp, 1 udp
	banner text,
	previous_banner text,
	port_open boolean default false,
	previous_port_open boolean default false,
	scanned_timestamp timestamptz not null default 'epoch',
	previous_scanned_timestamp timestamptz not null default 'epoch',
	check (host_address is not null or ip_address is not null),
    UNIQUE(scan_group_id, host_address, ip_address, port, protocol)
);


create table am.scan_group_addresses_ports_archive (
	like am.scan_group_addresses_ports,
	archived_timestamp timestamptz not null default 'epoch' 
);

create table am.scan_group_addresses_archive (
	like am.scan_group_addresses, 
	archived_timestamp timestamptz not null default 'epoch'
);

create table am.web_responses_archive (
	like am.web_responses, 
	archived_timestamp timestamptz not null default 'epoch'
);

create table am.web_technologies_archive (
	like am.web_technologies, 
	archived_timestamp timestamptz not null default 'epoch'
);

create table am.web_snapshots_archive (
	like am.web_snapshots, 
	archived_timestamp timestamptz not null default 'epoch'
);

grant select,insert,update,delete on am.scan_group_addresses_ports to addressservice;
grant select,insert,update,delete on am.scan_group_addresses_ports_archive to addressservice;
grant select,insert,update,delete on am.scan_group_addresses_archive to addressservice;
grant select,insert,update,delete on am.web_responses_archive to webdataservice;
grant select,insert,update,delete on am.web_technologies_archive to webdataservice;
grant select,insert,update,delete on am.web_snapshots_archive to webdataservice;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
revoke select,insert,update,delete on am.scan_group_addresses_ports from addressservice;
revoke select,insert,update,delete on am.scan_group_addresses_ports_archive from addressservice;
revoke select,insert,update,delete on am.scan_group_addresses_archive from addressservice;
revoke select,insert,update,delete on am.web_responses_archive from webdataservice;
revoke select,insert,update,delete on am.web_technologies_archive from webdataservice;
revoke select,insert,update,delete on am.web_snapshots_archive from webdataservice;

drop table am.web_snapshots_archive;
drop table am.web_technologies_archive;
drop table am.web_responses_archive;
drop table am.scan_group_addresses_archive;
drop table am.scan_group_addresses_ports_archive;
drop table am.scan_group_addresses_ports;