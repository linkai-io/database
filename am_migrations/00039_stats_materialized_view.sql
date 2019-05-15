-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE MATERIALIZED VIEW IF NOT EXISTS am.webdata_server_counts_mv as 
select wb.organization_id, wb.scan_group_id, agg.server, count(1) as cnt from 
	(select host_address,headers->>'server' as server,
			max(web_responses.url_request_timestamp) as url_request_timestamp,
			max(web_responses.response_timestamp) as response_timestamp from am.web_responses 
			where load_host_address=host_address 
			and load_ip_address=ip_address 
			and url_request_timestamp > now()-'7 days'::interval
			group by scan_group_id,server,host_address) as agg 
join am.web_responses as wb  on wb.url_request_timestamp=agg.url_request_timestamp and wb.response_timestamp=agg.response_timestamp 
join am.scan_group as sg on sg.scan_group_id=wb.scan_group_id where sg.deleted=false 
group by wb.organization_id,wb.scan_group_id,agg.server order by cnt desc;

CREATE UNIQUE INDEX IF NOT EXISTS unique_webdata_server_counts_mv_idx ON am.webdata_server_counts_mv (organization_id, scan_group_id, server);

grant select on am.webdata_server_counts_mv to webdataservice;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
revoke select on am.webdata_server_counts_mv from webdataservice;
drop index am.unique_webdata_server_counts_mv_idx;
drop materialized view am.webdata_server_counts_mv;
