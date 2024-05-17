package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type CreateSecgroupResponse struct { //_________________________________________________________________________________
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

type GetSecgroupByIdResponse struct { //________________________________________________________________________________
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
