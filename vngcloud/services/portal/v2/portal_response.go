package v2

import (
	lstrconv "strconv"

	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
)

const (
	QtVolumeAttachLimit = QuotaName("VOLUME_PER_SERVER") // The maximum number of volumes that you can attach to an server
)

type (
	Quota struct {
		Description string    `json:"description,omitempty"`
		Name        QuotaName `json:"quotaName"`
		Type        QuotaType `json:"type"`
		Used        string    `json:"used"`
		Limit       int       `json:"limit"`
	}

	QuotaType string
	QuotaName string
)

type ListAllQuotaUsedResponse struct {
	Data []Quota `json:"data"`
}

func (s *ListAllQuotaUsedResponse) ToEntityListQuotas() *lsentity.ListQuotas {
	listQuotas := &lsentity.ListQuotas{
		Items: make([]*lsentity.Quota, 0),
	}
	for _, q := range s.Data {
		listQuotas.Items = append(listQuotas.Items, q.ToEntityQuota())
	}

	return listQuotas
}

func (s *Quota) ToEntityQuota() *lsentity.Quota {
	var (
		used int
		err  error
	)

	if used, err = lstrconv.Atoi(s.Used); err != nil {
		used = 0
	}

	return &lsentity.Quota{
		Description: s.Description,
		Name:        string(s.Name),
		Type:        string(s.Type),
		Limit:       s.Limit,
		Used:        used,
	}
}
