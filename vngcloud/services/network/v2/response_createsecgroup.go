package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type CreateSecgroupResponse struct {
	Data struct {
		ID           int     `json:"id"`
		UUID         string  `json:"uuid"`
		CreatedAt    string  `json:"createdAt"`
		DeletedAt    *string `json:"deletedAt,omitempty"`
		Status       string  `json:"status"`
		SecgroupId   int     `json:"secgroupId"`
		SecgroupName string  `json:"secgroupName"`
		ProjectUuid  string  `json:"projectUuid"`
		Description  string  `json:"description"`
		UpdatedAt    *string `json:"updatedAt,omitempty"`
		IsSystem     bool    `json:"isSystem"`
		Type         string  `json:"type"`
	} `json:"data"`
}

func (s *CreateSecgroupResponse) ToEntitySecgroup() *lsentity.Secgroup {
	return &lsentity.Secgroup{
		Id:          s.Data.UUID,
		Name:        s.Data.SecgroupName,
		Description: s.Data.Description,
		Status:      s.Data.Status,
	}
}
