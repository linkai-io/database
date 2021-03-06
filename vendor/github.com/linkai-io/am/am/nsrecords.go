package am

// NSrecords is a map of record names to values
var NSRecords = map[string]int16{"NA": 0,
	"A":          1,
	"NS":         2,
	"MD":         3,
	"MF":         4,
	"CNAME":      5,
	"SOA":        6,
	"MB":         7,
	"MG":         8,
	"MR":         9,
	"NULL":       10,
	"PTR":        12,
	"HINFO":      13,
	"MINFO":      14,
	"MX":         15,
	"TXT":        16,
	"RP":         17,
	"AFSDB":      18,
	"X25":        19,
	"ISDN":       20,
	"RT":         21,
	"NSAPPTR":    23,
	"SIG":        24,
	"KEY":        25,
	"PX":         26,
	"GPOS":       27,
	"AAAA":       28,
	"LOC":        29,
	"NXT":        30,
	"EID":        31,
	"NIMLOC":     32,
	"SRV":        33,
	"ATMA":       34,
	"NAPTR":      35,
	"KX":         36,
	"CERT":       37,
	"DNAME":      39,
	"OPT":        41,
	"DS":         43,
	"SSHFP":      44,
	"RRSIG":      46,
	"NSEC":       47,
	"DNSKEY":     48,
	"DHCID":      49,
	"NSEC3":      50,
	"TLSA":       52,
	"SMIMEA":     53,
	"HIP":        55,
	"NINFO":      56,
	"RKEY":       57,
	"TALINK":     58,
	"CDS":        59,
	"CDNSKEY":    60,
	"OPENPGPKEY": 61,
	"CSYNC":      62,
	"SPF":        99,
	"UINFO":      100,
	"UID":        101,
	"GID":        102,
	"UNSPEC":     103,
	"NID":        104,
	"L32":        105,
	"L64":        106,
	"LP":         107,
	"EUI48":      108,
	"EUI64":      109,
	"URI":        256,
	"CAA":        257,
	"AVC":        258,
	"IXFR":       251,
	"AXFR":       252,
	"ANY":        255,
}
