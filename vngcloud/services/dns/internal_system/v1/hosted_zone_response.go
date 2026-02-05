package v1

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
)

type GetHostedZoneByIdResponse struct {
	Data *lsentity.HostedZone `json:"data"`
}

func (r *GetHostedZoneByIdResponse) ToEntityHostedZone() *lsentity.HostedZone {
	return r.Data
}

type ListHostedZonesResponse struct {
	ListData  []*lsentity.HostedZone `json:"listData"`
	Page      int                    `json:"page"`
	PageSize  int                    `json:"pageSize"`
	TotalPage int                    `json:"totalPage"`
	TotalItem int                    `json:"totalItem"`
}

func (r *ListHostedZonesResponse) ToEntityListHostedZones() *lsentity.ListHostedZone {
	return &lsentity.ListHostedZone{
		ListData:  r.ListData,
		Page:      r.Page,
		PageSize:  r.PageSize,
		TotalPage: r.TotalPage,
		TotalItem: r.TotalItem,
	}
}

type CreateHostedZoneResponse struct {
	Data *lsentity.HostedZone `json:"data"`
}

func (r *CreateHostedZoneResponse) ToEntityHostedZone() *lsentity.HostedZone {
	return r.Data
}
