package redisclient

import "fmt"

type RedisKeys struct {
	orgID     int
	groupID   int
	configFmt string
	statusFmt string
	addrFmt   string
	queueFmt  string
	portFmt   string
}

func NewRedisKeys(orgID, groupID int) *RedisKeys {
	r := &RedisKeys{orgID: orgID, groupID: groupID}
	r.configFmt = fmt.Sprintf("%d:%d:configuration", orgID, groupID)
	r.addrFmt = fmt.Sprintf("%d:%d:address", orgID, groupID)
	r.statusFmt = fmt.Sprintf("%d:%d:status", orgID, groupID)
	r.portFmt = fmt.Sprintf("%d:%d:portscan", orgID, groupID)
	return r
}

func (r *RedisKeys) Config() string {
	return r.configFmt
}

func (r *RedisKeys) Status() string {
	return r.statusFmt
}

func (r *RedisKeys) NSConfig() string {
	return r.configFmt + ":module:ns:config"
}

func (r *RedisKeys) NSZones() string {
	return r.configFmt + ":module:ns:zones"
}

// NSZone key for determining if we should do ns records or not
func (r *RedisKeys) NSZone(zone string) string {
	return r.configFmt + ":module:ns:zones:" + zone
}

func (r *RedisKeys) NSServers() string {
	return r.configFmt + ":module:ns:servers"
}

func (r *RedisKeys) BruteConfig() string {
	return r.configFmt + ":module:dnsbrute:config"
}

func (r *RedisKeys) BruteConfigHosts() string {
	return r.BruteConfig() + ":custom_hosts"
}

func (r *RedisKeys) BruteETLD(etld string) string {
	return r.configFmt + ":module:dnsbrute:zones:etld:" + etld
}

// BruteZone key for determining if we should brute force or not
func (r *RedisKeys) BruteZone(zone string) string {
	return r.configFmt + ":module:dnsbrute:zones:brute:" + zone
}

// MutateZone key for determining if we should mutate zone or not
func (r *RedisKeys) MutateZone(zone string) string {
	return r.configFmt + ":module:dnsbrute:zones:mutate:" + zone
}

// WebZone key for determining if we should mutate zone or not
func (r *RedisKeys) WebZone(zone string) string {
	return r.configFmt + ":module:web:zones:analysis:" + zone
}

// BigDataZone key for determining if we should look up domain in bigdata
func (r *RedisKeys) BigDataZone(zone string) string {
	return r.configFmt + ":module:bigdata:zones:" + zone
}

// PortZone key for determining if we should do a port scan for this hostname/ip
func (r *RedisKeys) PortZone(zone string) string {
	return r.configFmt + ":module:port:zones:" + zone
}

func (r *RedisKeys) PortConfig() string {
	return r.configFmt + ":module:port:config"
}

func (r *RedisKeys) PortConfigPorts() string {
	return r.PortConfig() + ":custom_ports"
}

func (r *RedisKeys) PortConfigTCPPorts() string {
	return r.PortConfig() + ":tcp_ports"
}

func (r *RedisKeys) PortConfigUDPPorts() string {
	return r.PortConfig() + ":udp_ports"
}

func (r *RedisKeys) PortConfigAllowedTLDs() string {
	return r.PortConfig() + ":allowed_tlds"
}

func (r *RedisKeys) PortConfigAllowedHosts() string {
	return r.PortConfig() + ":allowed_hosts"
}

func (r *RedisKeys) PortConfigDisallowedTLDs() string {
	return r.PortConfig() + ":disallowed_tlds"
}

func (r *RedisKeys) PortConfigDisallowedHosts() string {
	return r.PortConfig() + ":disallowed_hosts"
}

func (r *RedisKeys) WebConfig() string {
	return r.configFmt + ":module:web:config"
}

func (r *RedisKeys) KeywordConfig() string {
	return r.configFmt + ":module:keyword"
}

func (r *RedisKeys) AddrWorkQueue() string {
	return r.addrFmt + "_workqueue"
}

func (r *RedisKeys) AddrExistsHash() string {
	return r.addrFmt + "_hash"
}

func (r *RedisKeys) PortIP(host string) string {
	return r.portFmt + ":" + host + ":ip_address"
}
func (r *RedisKeys) PortResults(host, proto string) string {
	return r.portFmt + ":" + host + ":" + proto
}

func (r *RedisKeys) PortBannerResults(host, proto string) string {
	return r.portFmt + ":" + host + ":" + proto + ":banners"
}

// Addr returns the address key based on supplied addr id
func (r *RedisKeys) Addr(addrHash string) string {
	return fmt.Sprintf("%d:%d:address:%s", r.orgID, r.groupID, addrHash)
}

func (r *RedisKeys) AddrMatch() string {
	return r.addrFmt
}
