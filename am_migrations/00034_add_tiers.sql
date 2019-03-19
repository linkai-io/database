-- +goose Up
-- SQL in this section is executed when the migration is applied.
alter table am.organizations add column limit_tld integer not null default 0;
alter table am.organizations add column limit_tld_reached boolean not null default false;
alter table am.organizations add column limit_hosts integer not null default 0;
alter table am.organizations add column limit_hosts_reached boolean not null default false;
alter table am.organizations add column limit_custom_web_flows integer not null default 0;
alter table am.organizations add column limit_custom_web_flows_reached boolean not null default false;

alter table am.scan_group_addresses add column deleted boolean not null default false;

insert into am.subscription_types (subscription_id, title) values
    (101, 'monthly_small'),
    (102, 'monthly_medium');

create table am.scan_group_addresses_overflow (
    overflow_id bigserial not null primary key,
    organization_id integer REFERENCES am.organizations (organization_id),
    scan_group_id integer REFERENCES am.scan_group (scan_group_id),
    discovery_id integer REFERENCES am.scan_address_discovered_by (discovery_id),
    host_address varchar(512) not null,
    UNIQUE (organization_id, scan_group_id, host_address)
);

grant select (organization_id, limit_tld, limit_tld_reached, limit_hosts, limit_hosts_reached), update (limit_tld_reached, limit_hosts_reached) on am.organizations to addressservice;
grant select, insert, update, delete on am.scan_group_addresses_overflow to addressservice;


-- +goose StatementBegin
CREATE OR REPLACE FUNCTION am.handle_address_overflow() RETURNS TRIGGER AS $$
DECLARE
	cnt integer := 0;
	max_limit integer := 0;
	limit_reached boolean := false;
BEGIN
    -- we have various non important records (mx, ns etc)
    if NEW.confidence_score = 0 then 
        return new; 
    end if;

    -- get limits first
    select limit_hosts_reached, limit_hosts into limit_reached, max_limit from am.organizations where organization_id=NEW.organization_id;
	
    -- we reached the limit for this org, throw it in overflow
	if limit_reached then
		insert into am.scan_group_addresses_overflow (organization_id, scan_group_id, discovery_id, host_address) values 
			(NEW.organization_id, NEW.scan_group_id, NEW.discovery_id, NEW.host_address);
		return NULL;
	end if;
	
    -- ugh we need to count the number of 'good' hosts
    select count(*) into cnt from (SELECT host_address, (min(discovered_timestamp)) AS discovered_timestamp FROM am.scan_group_addresses 
        WHERE ignored = false 
        AND deleted = false
        AND organization_id = NEW.organization_id 
        AND (confidence_score = 100 OR user_confidence_score = 100)
        GROUP BY host_address ORDER BY discovered_timestamp asc) as total;
	
    -- check if we reached the limit if we did, that means we need to update limit_hosts_reached for this org and move this record into overflow
	if cnt >= max_limit then
		update am.organizations set limit_hosts_reached=true where organization_id=NEW.organization_id;
		insert into am.scan_group_addresses_overflow (organization_id, scan_group_id, discovery_id, host_address) values 
			(NEW.organization_id, NEW.scan_group_id, NEW.discovery_id, NEW.host_address);
		return NULL;
	end if;
	-- or, we are good, no limits have been reached yet
	return new;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

CREATE TRIGGER check_address_overflow BEFORE INSERT ON am.scan_group_addresses FOR EACH ROW EXECUTE PROCEDURE am.handle_address_overflow();

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION am.delete_org(org_id integer) RETURNS void as 
$BODY$
BEGIN 
    delete from am.ladon_role_member where organization_id=org_id;
    delete from am.ladon_role where organization_id=org_id;
    delete from am.scan_group_activity where organization_id=org_id;
    delete from am.scan_group_addresses_overflow where organization_id=org_id;
    delete from am.web_technologies where organization_id=org_id;
    delete from am.web_snapshots where organization_id=org_id;
    delete from am.web_responses where organization_id=org_id;
    delete from am.web_certificates where organization_id=org_id;
    delete from am.scan_group_findings where organization_id=org_id;
    delete from am.scan_group_addresses where organization_id=org_id;
    delete from am.scan_group_events where organization_id=org_id;
    delete from am.scan_group where organization_id=org_id;
    delete from am.users where organization_id=org_id; 
    delete from am.organizations where organization_id=org_id;
END
$BODY$ LANGUAGE plpgsql SECURITY DEFINER;
-- +goose StatementEnd

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION am.delete_org(org_id integer) RETURNS void as 
$BODY$
BEGIN 
    delete from am.ladon_role_member where organization_id=org_id;
    delete from am.ladon_role where organization_id=org_id;
    delete from am.scan_group_activity where organization_id=org_id;
    delete from am.web_technologies where organization_id=org_id;
    delete from am.web_snapshots where organization_id=org_id;
    delete from am.web_responses where organization_id=org_id;
    delete from am.web_certificates where organization_id=org_id;
    delete from am.scan_group_findings where organization_id=org_id;
    delete from am.scan_group_addresses where organization_id=org_id;
    delete from am.scan_group_events where organization_id=org_id;
    delete from am.scan_group where organization_id=org_id;
    delete from am.users where organization_id=org_id; 
    delete from am.organizations where organization_id=org_id;
END
$BODY$ LANGUAGE plpgsql SECURITY DEFINER;
-- +goose StatementEnd

drop trigger check_address_overflow on am.scan_group_addresses;
drop function am.handle_address_overflow;

revoke select (organization_id, limit_tld, limit_tld_reached, limit_hosts, limit_hosts_reached), update (limit_tld_reached, limit_hosts_reached) on am.organizations from addressservice;
revoke select, insert, update, delete on am.scan_group_addresses_overflow from addressservice;

drop table am.scan_group_addresses_overflow;
delete from am.subscription_types where subscription_id=101 or subscription_id=102;

alter table am.scan_group_addresses drop column deleted;

alter table am.organizations drop column limit_custom_web_flows_reached;
alter table am.organizations drop column limit_custom_web_flows;
alter table am.organizations drop column limit_hosts_reached;
alter table am.organizations drop column limit_hosts;
alter table am.organizations drop column limit_tld_reached;
alter table am.organizations drop column limit_tld;