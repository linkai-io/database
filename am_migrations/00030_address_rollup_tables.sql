-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE am.rollups (
    name text primary key,
    event_table_name text not null,
    event_id_sequence_name text not null,
    last_aggregated_id bigint default 0
);

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION am.incremental_rollup_window(rollup_name text, OUT window_start bigint, OUT window_end bigint)
RETURNS record
LANGUAGE plpgsql
AS $function$
DECLARE
    table_to_lock regclass;
BEGIN
    /*
     * Perform aggregation from the last aggregated ID + 1 up to the last committed ID.
     * We do a SELECT .. FOR UPDATE on the row in the rollup table to prevent
     * aggregations from running concurrently.
     */
    SELECT event_table_name, last_aggregated_id+1, pg_sequence_last_value(event_id_sequence_name)
    INTO table_to_lock, window_start, window_end
    FROM am.rollups
    WHERE name = rollup_name FOR UPDATE;

    IF NOT FOUND THEN
        RAISE 'rollup ''%'' is not in the rollups table', rollup_name;
    END IF;

    IF window_end IS NULL THEN
        /* sequence was never used */
        window_end := 0;
        RETURN;
    END IF;

    /*
     * Play a little trick: We very briefly lock the table for writes in order to
     * wait for all pending writes to finish. That way, we are sure that there are
     * no more uncommitted writes with a identifier lower or equal to window_end.
     * By throwing an exception, we release the lock immediately after obtaining it
     * such that writes can resume.
     */
    BEGIN
        EXECUTE format('LOCK %s IN EXCLUSIVE MODE', table_to_lock);
        RAISE 'release table lock';
    EXCEPTION WHEN OTHERS THEN
    END;

    /*
     * Remember the end of the window to continue from there next time.
     */
    UPDATE am.rollups SET last_aggregated_id = window_end WHERE name = rollup_name;
END;
$function$;
-- +goose StatementEnd


CREATE TABLE am.discoveries_1day (
    address_id bigint,
    organization_id int,
    scan_group_id int,
    period_start timestamptz,
    discovered_count bigint,
    primary key (address_id, organization_id, scan_group_id, period_start)
);

CREATE TABLE am.discoveries_3hour (
    address_id bigint,
    organization_id int,
    scan_group_id int,
    period_start timestamptz,
    discovered_count bigint,
    primary key (address_id, organization_id, scan_group_id, period_start)
);

CREATE TABLE am.seen_1day (
    address_id bigint,
    organization_id int,
    scan_group_id int,
    period_start timestamptz,
    seen_count bigint,
    primary key (address_id, organization_id, scan_group_id, period_start)
);

CREATE TABLE am.seen_3hour (
    address_id bigint,
    organization_id int,
    scan_group_id int,
    period_start timestamptz,
    seen_count bigint,
    primary key (address_id, organization_id, scan_group_id, period_start)
);

CREATE TABLE am.scanned_1day (
    address_id bigint,
    organization_id int,
    scan_group_id int,
    period_start timestamptz,
    scanned_count bigint,
    primary key (address_id, organization_id, scan_group_id, period_start)
);

CREATE TABLE am.scanned_3hour (
    address_id bigint,
    organization_id int,
    scan_group_id int,
    period_start timestamptz,
    scanned_count bigint,
    primary key (address_id, organization_id, scan_group_id, period_start)
);

-- Add the 1-day rollup to the rollups table
INSERT INTO am.rollups (name, event_table_name, event_id_sequence_name)
VALUES ('am.discoveries_1day_rollup', 'am.scan_group_addresses', 'am.scan_group_addresses_address_id_seq');

INSERT INTO am.rollups (name, event_table_name, event_id_sequence_name)
VALUES ('am.discoveries_3hour_rollup', 'am.scan_group_addresses', 'am.scan_group_addresses_address_id_seq');

INSERT INTO am.rollups (name, event_table_name, event_id_sequence_name)
VALUES ('am.seen_1day_rollup', 'am.scan_group_addresses', 'am.scan_group_addresses_address_id_seq');

INSERT INTO am.rollups (name, event_table_name, event_id_sequence_name)
VALUES ('am.seen_3hour_rollup', 'am.scan_group_addresses', 'am.scan_group_addresses_address_id_seq');

INSERT INTO am.rollups (name, event_table_name, event_id_sequence_name)
VALUES ('am.scanned_1day_rollup', 'am.scan_group_addresses', 'am.scan_group_addresses_address_id_seq');

INSERT INTO am.rollups (name, event_table_name, event_id_sequence_name)
VALUES ('am.scanned_3hour_rollup', 'am.scan_group_addresses', 'am.scan_group_addresses_address_id_seq');

-- discovered aggregates
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION am.do_daily_discovered_aggregation(OUT start_id bigint, OUT end_id bigint)
RETURNS record
LANGUAGE plpgsql
AS $function$
BEGIN
    /* determine which discoveries we can safely aggregate */
    SELECT window_start, window_end INTO start_id, end_id
    FROM am.incremental_rollup_window('am.discoveries_1day_rollup');

    /* exit early if there are no new page views to aggregate */
    IF start_id > end_id THEN RETURN; END IF;

    /* aggregate the page views, merge results if the entry already exists */
    INSERT INTO am.discoveries_1day (address_id, organization_id, scan_group_id, period_start, discovered_count)
      SELECT address_id, organization_id, scan_group_id, date_trunc('day', discovered_timestamp), count(*) AS discovered_count
      FROM am.scan_group_addresses
      WHERE address_id BETWEEN start_id AND end_id
      AND ignored=false
      GROUP BY address_id, organization_id, scan_group_id, date_trunc('day', discovered_timestamp)
      ON CONFLICT (address_id, organization_id, scan_group_id,  period_start) DO UPDATE
      SET discovered_count = am.discoveries_1day.discovered_count + EXCLUDED.discovered_count;
END;
$function$;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION am.do_trihourly_discovered_aggregation(OUT start_id bigint, OUT end_id bigint)
RETURNS record
LANGUAGE plpgsql
AS $function$
BEGIN
    /* determine which discoveries we can safely aggregate */
    SELECT window_start, window_end INTO start_id, end_id
    FROM am.incremental_rollup_window('am.discoveries_3hour_rollup');

    /* exit early if there are no new page views to aggregate */
    IF start_id > end_id THEN RETURN; END IF;

    /* aggregate the page views, merge results if the entry already exists */
    INSERT INTO am.discoveries_3hour (address_id, organization_id, scan_group_id, period_start, discovered_count)
      SELECT address_id, organization_id, scan_group_id, date_trunc('day', discovered_timestamp) + date_part('hour', discovered_timestamp)::int / 3 * interval '3 hour' as trihourly, count(*) AS discovered_count
      FROM am.scan_group_addresses
      WHERE address_id BETWEEN start_id AND end_id
      AND ignored=false
      GROUP BY address_id, organization_id, scan_group_id, trihourly
      ON CONFLICT (address_id, organization_id, scan_group_id, period_start) DO UPDATE
      SET discovered_count = am.discoveries_3hour.discovered_count + EXCLUDED.discovered_count;
END;
$function$;
-- +goose StatementEnd

-- last seen time aggregations
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION am.do_daily_seen_aggregation(OUT start_id bigint, OUT end_id bigint)
RETURNS record
LANGUAGE plpgsql
AS $function$
BEGIN
    /* determine which seen hosts we can safely aggregate */
    SELECT window_start, window_end INTO start_id, end_id
    FROM am.incremental_rollup_window('am.seen_1day_rollup');

    /* exit early if there are no new page views to aggregate */
    IF start_id > end_id THEN RETURN; END IF;

    /* aggregate the page views, merge results if the entry already exists */
    INSERT INTO am.seen_1day (address_id, organization_id, scan_group_id, period_start, seen_count)
      SELECT address_id, organization_id, scan_group_id, date_trunc('day', last_seen_timestamp), count(*) AS seen_count
      FROM am.scan_group_addresses
      WHERE address_id BETWEEN start_id AND end_id
      AND ignored=false
      GROUP BY address_id, organization_id, scan_group_id, date_trunc('day', last_seen_timestamp)
      ON CONFLICT (address_id, organization_id, scan_group_id,  period_start) DO UPDATE
      SET seen_count = am.seen_1day.seen_count + EXCLUDED.seen_count;
END;
$function$;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION am.do_trihourly_seen_aggregation(OUT start_id bigint, OUT end_id bigint)
RETURNS record
LANGUAGE plpgsql
AS $function$
BEGIN
    /* determine which seen hosts we can safely aggregate */
    SELECT window_start, window_end INTO start_id, end_id
    FROM am.incremental_rollup_window('am.seen_3hour_rollup');

    /* exit early if there are no new page views to aggregate */
    IF start_id > end_id THEN RETURN; END IF;

    /* aggregate the seen updates divide the day by 3 hour intervals, merge results if the entry already exists */
    INSERT INTO am.seen_3hour (address_id, organization_id, scan_group_id, period_start, seen_count)
      SELECT address_id, organization_id, scan_group_id,  date_trunc('day', last_seen_timestamp) + date_part('hour', last_seen_timestamp)::int / 3 * interval '3 hour' as trihourly, count(*) AS seen_count
      FROM am.scan_group_addresses
      WHERE address_id BETWEEN start_id AND end_id
      AND ignored=false
      GROUP BY address_id, organization_id, scan_group_id, trihourly
      ON CONFLICT (address_id, organization_id, scan_group_id,  period_start) DO UPDATE
      SET seen_count = am.seen_3hour.seen_count + EXCLUDED.seen_count;
END;
$function$;
-- +goose StatementEnd

-- last scanned time aggregations
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION am.do_daily_scanned_aggregation(OUT start_id bigint, OUT end_id bigint)
RETURNS record
LANGUAGE plpgsql
AS $function$
BEGIN
    /* determine which seen hosts we can safely aggregate */
    SELECT window_start, window_end INTO start_id, end_id
    FROM am.incremental_rollup_window('am.scanned_1day_rollup');

    /* exit early if there are no new page views to aggregate */
    IF start_id > end_id THEN RETURN; END IF;

    /* aggregate the page views, merge results if the entry already exists */
    INSERT INTO am.scanned_1day (address_id, organization_id, scan_group_id, period_start, scanned_count)
      SELECT address_id, organization_id, scan_group_id, date_trunc('day', last_scanned_timestamp), count(*) AS scanned_count
      FROM am.scan_group_addresses
      WHERE address_id BETWEEN start_id AND end_id
      AND ignored=false
      GROUP BY address_id, organization_id, scan_group_id, date_trunc('day', last_scanned_timestamp)
      ON CONFLICT (address_id, organization_id, scan_group_id,  period_start) DO UPDATE
      SET scanned_count = am.scanned_1day.scanned_count + EXCLUDED.scanned_count;
END;
$function$;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION am.do_trihourly_scanned_aggregation(OUT start_id bigint, OUT end_id bigint)
RETURNS record
LANGUAGE plpgsql
AS $function$
BEGIN
    /* determine which seen hosts we can safely aggregate */
    SELECT window_start, window_end INTO start_id, end_id
    FROM am.incremental_rollup_window('am.scanned_3hour_rollup');

    /* exit early if there are no new page views to aggregate */
    IF start_id > end_id THEN RETURN; END IF;

    /* aggregate the seen updates divide the day by 3 hour intervals, merge results if the entry already exists */
    INSERT INTO am.scanned_3hour (address_id, organization_id, scan_group_id, period_start, scanned_count)
      SELECT address_id, organization_id, scan_group_id,  date_trunc('day', last_scanned_timestamp) + date_part('hour', last_scanned_timestamp)::int / 3 * interval '3 hour' as trihourly, count(*) AS scanned_count
      FROM am.scan_group_addresses
      WHERE address_id BETWEEN start_id AND end_id
      AND ignored=false
      GROUP BY address_id, organization_id, scan_group_id, trihourly
      ON CONFLICT (address_id, organization_id, scan_group_id,  period_start) DO UPDATE
      SET scanned_count = am.scanned_3hour.scanned_count + EXCLUDED.scanned_count;
END;
$function$;
-- +goose StatementEnd

grant select on am.discoveries_1day to addressservice;
grant select on am.discoveries_3hour to addressservice;
grant select on am.seen_1day to addressservice;
grant select on am.seen_3hour to addressservice;
grant select on am.scanned_1day to addressservice;
grant select on am.scanned_3hour to addressservice;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
revoke select on am.discoveries_1day from addressservice;
revoke select on am.discoveries_3hour from addressservice;
revoke select on am.seen_1day from addressservice;
revoke select on am.seen_3hour from addressservice;
revoke select on am.scanned_1day from addressservice;
revoke select on am.scanned_3hour from addressservice;

drop function am.do_trihourly_scanned_aggregation(OUT start_id bigint, OUT end_id bigint);
drop function am.do_daily_scanned_aggregation(OUT start_id bigint, OUT end_id bigint);
drop function am.do_trihourly_seen_aggregation(OUT start_id bigint, OUT end_id bigint);
drop function am.do_daily_seen_aggregation(OUT start_id bigint, OUT end_id bigint);
drop function am.do_trihourly_discovered_aggregation(OUT start_id bigint, OUT end_id bigint);
drop function am.do_daily_discovered_aggregation(OUT start_id bigint, OUT end_id bigint);

drop table am.scanned_3hour;
drop table am.scanned_1day;
drop table am.seen_3hour;
drop table am.seen_1day;
drop table am.discoveries_3hour;
drop table am.discoveries_1day;

drop function am.incremental_rollup_window(rollup_name text, OUT window_start bigint, OUT window_end bigint);
drop table am.rollups;