package v1

type HostedZoneType string

const (
	// HostedZoneTypePrivate represents a private hosted zone
	HostedZoneTypePrivate HostedZoneType = "PRIVATE"
)

// DnsRecordType represents the type of DNS record
type DnsRecordType string

const (
	// DnsRecordTypeA represents an A record
	DnsRecordTypeA DnsRecordType = "A"

	// DnsRecordTypePTR represents a PTR record
	DnsRecordTypePTR DnsRecordType = "PTR"

	// DnsRecordTypeSRV represents an SRV record
	DnsRecordTypeSRV DnsRecordType = "SRV"

	// DnsRecordTypeMX represents an MX record
	DnsRecordTypeMX DnsRecordType = "MX"

	// DnsRecordTypeCNAME represents a CNAME record
	DnsRecordTypeCNAME DnsRecordType = "CNAME"

	// DnsRecordTypeTXT represents a TXT record
	DnsRecordTypeTXT DnsRecordType = "TXT"
)

// RoutingPolicy represents the routing policy for DNS records
type RoutingPolicy string

const (
	// RoutingPolicySimple represents simple routing
	RoutingPolicySimple RoutingPolicy = "simple-routing"

	// RoutingPolicyWeighted represents weighted routing
	RoutingPolicyWeighted RoutingPolicy = "weighted"

	// RoutingPolicyGeolocation represents geolocation routing
	RoutingPolicyGeolocation RoutingPolicy = "geolocation"
)
