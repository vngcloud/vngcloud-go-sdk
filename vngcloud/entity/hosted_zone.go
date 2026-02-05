package entity

import "time"

type VpcMapRegion struct {
	VpcId  string `json:"vpcId"`
	Region string `json:"region"`
}

type HostedZone struct {
	HostedZoneId      string         `json:"hostedZoneId"`
	DomainName        string         `json:"domainName"`
	Status            string         `json:"status"`
	Description       string         `json:"description"`
	Type              string         `json:"type"`
	CountRecords      int            `json:"countRecords"`
	AssocVpcIds       []string       `json:"assocVpcIds"`
	AssocVpcMapRegion []VpcMapRegion `json:"assocVpcMapRegion"`
	PortalUserId      int            `json:"portalUserId"`
	CreatedAt         time.Time      `json:"createdAt"`
	DeletedAt         *time.Time     `json:"deletedAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
}

type ListHostedZone struct {
	ListData  []*HostedZone `json:"listData"`
	Page      int           `json:"page"`
	PageSize  int           `json:"pageSize"`
	TotalPage int           `json:"totalPage"`
	TotalItem int           `json:"totalItem"`
}
