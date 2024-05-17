package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type GetSecgroupByIdResponse struct {
	Data struct {
		ID          string  `json:"id"`
		Name        string  `json:"name"`
		Description *string `json:"description"`
		Status      string  `json:"status"`
		CreatedAt   string  `json:"createdAt"`
		IsSystem    bool    `json:"isSystem"`
	} `json:"data"`
}

func (s *GetSecgroupByIdResponse) ToEntitySecgroup() *lsentity.Secgroup {
	return &lsentity.Secgroup{
		Id:          s.Data.ID,
		Name:        s.Data.Name,
		Description: *s.Data.Description,
		Status:      s.Data.Status,
	}
}
