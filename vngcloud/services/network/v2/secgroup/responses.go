package secgroup

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

type CreateResponse struct {
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

func (s *CreateResponse) ToSecgroupObject() *objects.Secgroup {
	if s == nil {
		return nil
	}

	return &objects.Secgroup{
		UUID:        s.Data.UUID,
		Name:        s.Data.SecgroupName,
		Description: s.Data.Description,
		Status:      s.Data.Status,
	}
}

type GetResponse struct {
	Data struct {
		ID          string  `json:"id"`
		Name        string  `json:"name"`
		Description *string `json:"description"`
		Status      string  `json:"status"`
		CreatedAt   string  `json:"createdAt"`
		IsSystem    bool    `json:"isSystem"`
	} `json:"data"`
}

func (s *GetResponse) ToSecgroupObject() *objects.Secgroup {
	if s == nil {
		return nil
	}

	return &objects.Secgroup{
		UUID:        s.Data.ID,
		Name:        s.Data.Name,
		Description: *s.Data.Description,
		Status:      s.Data.Status,
	}
}

// *************************************************** ListResponse ****************************************************

type ListResponse struct {
	ListData []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
		Status      string `json:"status"`
		CreatedAt   string `json:"createdAt"`
		IsSystem    bool   `json:"isSystem"`
	} `json:"listData"`
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	TotalPage int `json:"totalPage"`
	TotalItem int `json:"totalItem"`
}

func (s *ListResponse) ToListSecgroupObjects() []*objects.Secgroup {
	var secgroups []*objects.Secgroup
	if s == nil || s.ListData == nil || len(s.ListData) == 0 {
		return secgroups
	}

	for _, secgroup := range s.ListData {
		secgroups = append(secgroups, &objects.Secgroup{
			UUID:        secgroup.ID,
			Name:        secgroup.Name,
			Description: secgroup.Description,
			Status:      secgroup.Status,
		})
	}

	return secgroups
}
