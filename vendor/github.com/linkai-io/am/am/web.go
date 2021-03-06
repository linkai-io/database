package am

import (
	"context"
	"time"
)

const (
	RNWebData             = "lrn:service:webdata:feature:"
	RNWebDataResponses    = "lrn:service:webdata:feature:responses"
	RNWebDataCertificates = "lrn:service:webdata:feature:certificates"
	RNWebDataSnapshots    = "lrn:service:webdata:feature:snapshots"
	WebDataServiceKey     = "webdataservice"
)

const (
	FilterDeleted                  = "deleted"
	FilterWebTechType              = "tech_type"
	FilterWebTechTypeVersion       = "tech_type_version"
	FilterWebDependentHostAddress  = "dependent_host_address"
	FilterWebEqualsHostAddress     = "host_address"
	FilterWebAfterResponseTime     = "after_response_time"
	FilterWebBeforeResponseTime    = "before_response_time"
	FilterWebLatestOnly            = "latest_only"
	FilterWebMimeType              = "mime_type"
	FilterWebHeaderNames           = "header_names"
	FilterWebNotHeaderNames        = "not_header_names"
	FilterWebHeaderPairNames       = "header_pair_names"
	FilterWebHeaderPairValues      = "header_pair_values"
	FilterWebAfterURLRequestTime   = "after_request_time"
	FilterWebBeforeURLRequestTime  = "before_request_time"
	FilterWebEqualsURLRequestTime  = "url_request_timestamp"
	FilterWebEqualsResponseTime    = "response_timestamp"
	FilterWebEqualsIPAddress       = "ip_address"
	FilterWebEndsHostAddress       = "ends_host_address"
	FilterWebStartsHostAddress     = "starts_host_address"
	FilterWebEqualsLoadIPAddress   = "load_ip_address"
	FilterWebEqualsLoadHostAddress = "load_host_address"
	FilterWebEndsLoadHostAddress   = "ends_load_host_address"
	FilterWebStartsLoadHostAddress = "starts_load_host_address"
	FilterWebEqualsServerType      = "server_type"
	FilterWebEqualsURL             = "url"
	FilterWebAfterValidTo          = "after_valid_to"
	FilterWebBeforeValidTo         = "before_valid_to"
	FilterWebAfterValidFrom        = "after_valid_from"
	FilterWebBeforeValidFrom       = "before_valid_from"
)

type WebTech struct {
	Matched  string `json:"matched"`
	Version  string `json:"version"`
	Location string `json:"location"`
}

type WebTechResults struct {
	WebTech
	CategoryName string `json:"category"`
	Icon         string `json:"icon"`
	Website      string `json:"website"`
}

// WebData is the primary container of a scangroup address's http response data.
type WebData struct {
	Address             *ScanGroupAddress   `json:"address"`
	Responses           []*HTTPResponse     `json:"responses"`
	Snapshot            string              `json:"snapshot,omitempty"` // only used internally by browser package
	SnapshotLink        string              `json:"snapshot_link"`
	URL                 string              `json:"url"`
	Scheme              string              `json:"scheme"`
	AddressHash         string              `json:"address_hash"`
	HostAddress         string              `json:"host_address"`
	IPAddress           string              `json:"ip_address"`
	ResponsePort        int                 `json:"response_port"`
	RequestedPort       int                 `json:"requested_port"`
	SerializedDOM       string              `json:"serialized_dom,omitempty"` // only used internally by browser package
	SerializedDOMHash   string              `json:"serialized_dom_hash"`
	SerializedDOMLink   string              `json:"serialized_dom_link"`
	ResponseTimestamp   int64               `json:"response_timestamp"`
	URLRequestTimestamp int64               `json:"url_request_timestamp"`
	DetectedTech        map[string]*WebTech `json:"detected_tech"`
	LoadURL             string              `json:"load_url"`
}

type URLData struct {
	ResponseID  int64  `json:"response_id"`
	URL         string `json:"url"`
	RawBodyLink string `json:"raw_body_link"`
	MimeType    string `json:"mime_type"`
}

// URLListResponse is for holding a list of urls from an intial request
type URLListResponse struct {
	OrgID               int        `json:"organization_id,omitempty"`
	GroupID             int        `json:"group_id,omitempty"`
	URLRequestTimestamp int64      `json:"url_request_timestamp"`
	HostAddress         string     `json:"host_address,omitempty"` // used for returning data to user
	IPAddress           string     `json:"ip_address,omitempty"`   // used for returning data to user
	URLs                []*URLData `json:"urls"`
}

type WebDomainNode struct {
	ID     string `json:"id"`
	Origin int32  `json:"origin"`
}

type WebDomainLink struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type WebDomainDependency struct {
	Status    string           `json:"status,omitempty"`
	OrgID     int              `json:"organization_id,omitempty"`
	GroupID   int              `json:"group_id,omitempty"`
	LastIndex int64            `json:"last_index"`
	Nodes     []*WebDomainNode `json:"nodes"`
	Links     []*WebDomainLink `json:"links"`
}

// HTTPResponse represents a captured network response
type HTTPResponse struct {
	ResponseID          int64             `json:"response_id,omitempty"`
	OrgID               int               `json:"organization_id,omitempty"`
	GroupID             int               `json:"group_id,omitempty"`
	Scheme              string            `json:"scheme"`
	AddressHash         string            `json:"address_hash"`
	HostAddress         string            `json:"host_address"`
	IPAddress           string            `json:"ip_address"`
	ResponsePort        string            `json:"response_port"`
	RequestedPort       string            `json:"requested_port"`
	RequestID           string            `json:"request_id,omitempty"` // only used internally by browser package
	Status              int               `json:"status"`
	StatusText          string            `json:"status_text"`
	URL                 string            `json:"url"`
	Headers             map[string]string `json:"headers"`
	MimeType            string            `json:"mime_type"`
	RawBody             string            `json:"raw_body,omitempty"` // only used internally by browser package
	RawBodyLink         string            `json:"raw_body_link"`
	RawBodyHash         string            `json:"raw_body_hash"`
	ResponseTimestamp   int64             `json:"response_timestamp"`
	URLRequestTimestamp int64             `json:"url_request_timestamp"`
	IsDocument          bool              `json:"is_document"`
	WebCertificate      *WebCertificate   `json:"web_certificate,omitempty"`
	IsDeleted           bool              `json:"deleted"`
	LoadHostAddress     string            `json:"load_host_address"`
	LoadIPAddress       string            `json:"load_ip_address"`
}

// WebCertificate contains extracted certificate data from an HTTPResponse
type WebCertificate struct {
	OrgID                             int      `json:"organization_id,omitempty"`
	GroupID                           int      `json:"group_id,omitempty"`
	CertificateID                     int64    `json:"certificate_id,omitempty"` // from DB
	ResponseTimestamp                 int64    `json:"response_timestamp,omitempty"`
	HostAddress                       string   `json:"host_address"`
	IPAddress                         string   `json:"ip_address"`
	AddressHash                       string   `json:"address_hash"`
	Port                              string   `json:"port"`
	Protocol                          string   `json:"protocol"`                          // Protocol name (e.g. "TLS 1.2" or "QUIC").
	KeyExchange                       string   `json:"keyExchange"`                       // Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchangeGroup                  string   `json:"keyExchangeGroup,omitempty"`        // (EC)DH group used by the connection, if applicable.
	Cipher                            string   `json:"cipher"`                            // Cipher name.
	Mac                               string   `json:"mac,omitempty"`                     // TLS MAC. Note that AEAD ciphers do not have separate MACs.
	CertificateValue                  int      `json:"certificateId"`                     // from browser                    // Certificate ID value.
	SubjectName                       string   `json:"subjectName"`                       // Certificate subject name.
	SanList                           []string `json:"sanList"`                           // Subject Alternative Name (SAN) DNS names and IP addresses.
	Issuer                            string   `json:"issuer"`                            // Name of the issuing CA.
	ValidFrom                         int64    `json:"validFrom"`                         // Certificate valid from date.
	ValidTo                           int64    `json:"validTo"`                           // Certificate valid to (expiration) date
	CertificateTransparencyCompliance string   `json:"certificateTransparencyCompliance"` // Whether the request complied with Certificate Transparency policy enum values: unknown, not-compliant, compliant
	IsDeleted                         bool     `json:"deleted"`
}

// WebSnapshot for returning serialized dom and image snapshot links
type WebSnapshot struct {
	SnapshotID          int64    `json:"snapshot_id,omitempty"`
	OrgID               int      `json:"org_id,omitempty"`
	GroupID             int      `json:"group_id,omitempty"`
	SnapshotLink        string   `json:"snapshot_link"`
	SerializedDOMHash   string   `json:"serialized_dom_hash"`
	SerializedDOMLink   string   `json:"serialized_dom_link"`
	ResponseTimestamp   int64    `json:"response_timestamp"`
	IsDeleted           bool     `json:"deleted"`
	URL                 string   `json:"url"`
	AddressHash         string   `json:"address_hash"` // unfortunately we can't FK off address_id since we may get a new ip/host before it exists
	HostAddress         string   `json:"host_address"`
	IPAddress           string   `json:"ip_address"`
	ResponsePort        int      `json:"response_port"`
	RequestedPort       int      `json:"requested_port"`
	Scheme              string   `json:"scheme"`
	TechCategories      []string `json:"tech_categories"`
	TechNames           []string `json:"tech_names"`
	TechVersions        []string `json:"tech_versions"`
	TechMatchLocations  []string `json:"tech_match_locations"`
	TechMatchData       []string `json:"tech_match_data"`
	TechIcons           []string `json:"tech_icons"`
	TechWebsites        []string `json:"tech_websites"`
	LoadURL             string   `json:"load_url"`
	URLRequestTimestamp int64    `json:"request_timestamp"`
}

// WebResponseFilter used to filter results when searching web data.
type WebResponseFilter struct {
	OrgID   int         `json:"org_id"`
	GroupID int         `json:"group_id"`
	Filters *FilterType `json:"filters"`
	Start   int64       `json:"start"`
	Limit   int         `json:"limit"`
}

// WebSnapshotFilter for filtering websnapshots
type WebSnapshotFilter struct {
	OrgID   int         `json:"org_id"`
	GroupID int         `json:"group_id"`
	Start   int64       `json:"start"`
	Limit   int         `json:"limit"`
	Filters *FilterType `json:"filters"`
}

// WebCertificateFilter used to filter results when searching web data.
type WebCertificateFilter struct {
	OrgID   int         `json:"org_id"`
	GroupID int         `json:"group_id"`
	Start   int64       `json:"start"`
	Limit   int         `json:"limit"`
	Filters *FilterType `json:"filters"`
}

// ScanGroupWebDataStats returns general statistics for web related asset data
type ScanGroupWebDataStats struct {
	OrgID               int      `json:"org_id"`
	GroupID             int      `json:"group_id"`
	ExpiringCerts15Days int32    `json:"expiring_certs_15"`
	ExpiringCerts30Days int32    `json:"expiring_certs_30"`
	UniqueWebServers    int32    `json:"unique_web_servers"`
	ServerTypes         []string `json:"server_types"`
	ServerCounts        []int32  `json:"server_counts"`
}

// WebDataService adds and returns all web responses
type WebDataService interface {
	Init(config []byte) error
	Add(ctx context.Context, userContext UserContext, webData *WebData) (int, error)
	OrgStats(ctx context.Context, userContext UserContext) (oid int, orgStats []*ScanGroupWebDataStats, err error)
	GroupStats(ctx context.Context, userContext UserContext, groupID int) (oid int, groupStats *ScanGroupWebDataStats, err error)
	GetURLList(ctx context.Context, userContext UserContext, filter *WebResponseFilter) (int, []*URLListResponse, error)
	GetResponses(ctx context.Context, userContext UserContext, filter *WebResponseFilter) (int, []*HTTPResponse, error)
	GetCertificates(ctx context.Context, userContext UserContext, filter *WebCertificateFilter) (int, []*WebCertificate, error)
	GetSnapshots(ctx context.Context, userContext UserContext, filter *WebSnapshotFilter) (int, []*WebSnapshot, error)
	GetDomainDependency(ctx context.Context, userContext UserContext, filter *WebResponseFilter) (int, *WebDomainDependency, error)
	Archive(ctx context.Context, userContext UserContext, group *ScanGroup, archiveTime time.Time) (int, int, error)
}
