package entity

import "time"

type RecordValue struct {
	Value    string  `json:"value"`
	Location *string `json:"location"`
	Weight   *int    `json:"weight"`
}

type DnsRecord struct {
	RecordId            string        `json:"recordId"`
	SubDomain           string        `json:"subDomain"`
	HostedZoneId        string        `json:"hostedZoneId"`
	Status              string        `json:"status"`
	Type                string        `json:"type"`
	RoutingPolicy       string        `json:"routingPolicy"`
	Value               []RecordValue `json:"value"`
	TTL                 int           `json:"ttl"`
	EnableStickySession *bool         `json:"enableStickySession"`
	CreatedAt           time.Time     `json:"createdAt"`
	DeletedAt           *time.Time    `json:"deletedAt"`
	UpdatedAt           time.Time     `json:"updatedAt"`
}

type ListDnsRecords struct {
	ListData  []*DnsRecord `json:"listData"`
	Page      int          `json:"page"`
	PageSize  int          `json:"pageSize"`
	TotalPage int          `json:"totalPage"`
	TotalItem int          `json:"totalItem"`
}
