-- +goose Up
-- SQL in this section is executed when the migration is applied.
drop index am.scan_group_addresses_ports_address_hash_idx;
drop index am.scan_group_addresses_ports_scanned_timestamp_idx;


-- we need to recreate the entire ports/archive tables again because of constraints
drop table am.scan_group_addresses_ports;
drop table am.scan_group_addresses_ports_archive;

create table am.scan_group_addresses_ports (
	port_id bigserial primary key not null,
	organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_id integer REFERENCES am.scan_group (scan_group_id),
    host_address varchar(512) not null default '',
	port_data jsonb not null default '{}'::jsonb,
	scanned_timestamp timestamptz not null default 'epoch',
	previous_scanned_timestamp timestamptz not null default 'epoch',
	check (host_address is not null),
    UNIQUE(scan_group_id, host_address)
);

create table am.scan_group_addresses_ports_archive (
	like am.scan_group_addresses_ports,
	archived_timestamp timestamptz not null default 'epoch' 
);

create index scan_group_addresses_ports_scanned_timestamp_idx on am.scan_group_addresses_ports (organization_id, scan_group_id, scanned_timestamp);
create index scan_group_addresses_ports_host_address_idx on am.scan_group_addresses_ports (organization_id, scan_group_id, host_address);

alter table only am.scan_group_addresses_archive add column port_scan_override_tld boolean default false;
alter table only am.scan_group_addresses_archive add column port_scan_enabled boolean default false;

alter table only am.scan_group_addresses add column port_scan_override_tld boolean default false;
alter table only am.scan_group_addresses add column port_scan_enabled boolean default false;

alter table only am.organizations add column port_scan_feature_enabled boolean default false;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
alter table only am.organizations drop column port_scan_feature_enabled;

alter table only am.scan_group_addresses drop column port_scan_override_tld;
alter table only am.scan_group_addresses drop column port_scan_enabled;

alter table only am.scan_group_addresses_archive drop port_scan_override_tld;
alter table only am.scan_group_addresses_archive drop column port_scan_enabled;

drop index am.scan_group_addresses_ports_host_address_idx;
drop index am.scan_group_addresses_ports_scanned_timestamp_idx;

drop table am.scan_group_addresses_ports;
drop table am.scan_group_addresses_ports_archive; 

create table am.scan_group_addresses_ports (
	port_id bigserial primary key not null,
	organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_id integer REFERENCES am.scan_group (scan_group_id),
    host_address varchar(512) not null default '',
    ip_address varchar(256) not null default '',
	address_hash varchar(128) not null default '',
	port_data jsonb not null default '{}'::jsonb,
	scanned_timestamp timestamptz not null default 'epoch',
	previous_scanned_timestamp timestamptz not null default 'epoch',
	check (host_address is not null or ip_address is not null),
    UNIQUE(scan_group_id, host_address, ip_address)
);

create table am.scan_group_addresses_ports_archive (
	like am.scan_group_addresses_ports,
	archived_timestamp timestamptz not null default 'epoch' 
);

create index scan_group_addresses_ports_scanned_timestamp_idx on am.scan_group_addresses_ports (scanned_timestamp desc);
create index scan_group_addresses_ports_address_hash_idx on am.scan_group_addresses_ports (address_hash);


