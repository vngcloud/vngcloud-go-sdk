package v1

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
)

type ListRecordsResponse struct {
	ListData  []*lsentity.DnsRecord `json:"listData"`
	Page      int                   `json:"page"`
	PageSize  int                   `json:"pageSize"`
	TotalPage int                   `json:"totalPage"`
	TotalItem int                   `json:"totalItem"`
}

func (r *ListRecordsResponse) ToEntityListRecords() *lsentity.ListDnsRecords {
	return &lsentity.ListDnsRecords{
		ListData:  r.ListData,
		Page:      r.Page,
		PageSize:  r.PageSize,
		TotalPage: r.TotalPage,
		TotalItem: r.TotalItem,
	}
}