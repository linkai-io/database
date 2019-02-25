-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- update all web snapshot columns/constraints to not use address_id's anymore
alter table am.web_snapshots drop constraint web_snapshots_organization_id_scan_group_id_serialized_dom__key;
alter table am.web_snapshots drop constraint web_snapshots_address_id_fkey;

alter table am.web_snapshots 
    add column url bytea not null,
    add column address_hash varchar(128) not null default '',
    add column host_address varchar(512) not null default '',
    add column ip_address varchar(256) not null default '',
    add column scheme varchar(12) not null default '',
    add column response_port int not null default 0;

alter table am.web_snapshots rename is_deleted to deleted;
update am.web_snapshots
    set (address_hash, host_address, ip_address) =
        (select old.address_hash, old.host_address, old.ip_address from am.scan_group_addresses as old 
            where old.address_id=address_id);
                
alter table am.web_snapshots drop column address_id;
alter table am.web_snapshots add constraint unique_address_dom unique (organization_id, scan_group_id, address_hash, serialized_dom_hash, response_port);

-- update web_responses
alter table am.web_responses rename is_deleted to deleted;
alter table am.web_responses add column address_hash varchar(128) not null default '';

update am.web_responses 
    set (address_hash) = (select old.address_hash from am.scan_group_addresses as old 
        where old.address_id=address_id);

alter table am.web_responses drop column address_id;

-- update web_certificates
alter table am.web_certificates rename is_deleted to deleted;
alter table am.web_certificates add column address_hash varchar(128) not null default '',
    add column ip_address varchar(256) not null default '';

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
-- revert web_certificates
alter table am.web_certificates drop column address_hash;
alter table am.web_certificates drop column ip_address;
alter table am.web_certificates rename deleted to is_deleted;

-- revert web_responses
alter table am.web_responses add column address_id bigint default 0;
update am.web_responses 
    set address_id = (select old.address_id from am.scan_group_addresses as old 
        where old.address_hash=address_hash and old.organization_id=organization_id and 
        old.scan_group_id=scan_group_id);

alter table am.web_responses drop column address_hash;
alter table am.web_responses rename deleted to is_deleted;

-- revert snapshots
alter table am.web_snapshots drop constraint unique_address_dom;
alter table am.web_snapshots add column address_id bigint default 0;

alter table am.web_snapshots rename deleted to is_deleted;
update am.web_snapshots 
    set ip_address=(select old.address_id from am.scan_group_addresses as old 
            where old.address_hash=address_hash and old.organization_id=organization_id and old.scan_group_id=scan_group_id);

alter table am.web_snapshots drop column url;
alter table am.web_snapshots drop column address_hash;
alter table am.web_snapshots drop column host_address;
alter table am.web_snapshots drop column ip_address;
alter table am.web_snapshots drop column scheme;
alter table am.web_snapshots drop column response_port;

alter table am.web_snapshots add constraint web_snapshots_address_id_fkey foreign key (address_id) references am.scan_group_addresses (address_id);
alter table am.web_snapshots add constraint web_snapshots_organization_id_scan_group_id_serialized_dom__key unique (organization_id, scan_group_id, serialized_dom_hash);
