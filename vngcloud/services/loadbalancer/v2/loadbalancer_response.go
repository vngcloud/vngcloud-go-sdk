package v2

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lstr "strings"
)

type CreateLoadBalancerResponse struct {
	UUID string `json:"uuid"`
}

type GetLoadBalancerByIdResponse struct {
	Data LoadBalancer `json:"data"`
}

type ListLoadBalancersResponse struct {
	ListData  []LoadBalancer `json:"listData"`
	Page      int            `json:"page"`
	PageSize  int            `json:"pageSize"`
	TotalPage int            `json:"totalPage"`
	TotalItem int            `json:"totalItem"`
}

type (
	LoadBalancer struct {
		UUID               string `json:"uuid"`
		Name               string `json:"name"`
		DisplayStatus      string `json:"displayStatus"`
		Address            string `json:"address"`
		PrivateSubnetID    string `json:"privateSubnetId"`
		PrivateSubnetCidr  string `json:"privateSubnetCidr"`
		Type               string `json:"type"`
		DisplayType        string `json:"displayType"`
		LoadBalancerSchema string `json:"loadBalancerSchema"`
		PackageID          string `json:"packageId"`
		Description        string `json:"description"`
		Location           string `json:"location"`
		CreatedAt          string `json:"createdAt"`
		UpdatedAt          string `json:"updatedAt"`
		PackageInfo        struct {
			PackageID        string `json:"packageId"`
			ConnectionNumber int    `json:"connectionNumber"`
			DataTransfer     int    `json:"dataTransfer"`
			Name             string `json:"name"`
		} `json:"packageInfo"`
		ProgressStatus string `json:"progressStatus"`
	}
)

func (s *CreateLoadBalancerResponse) ToEntityLoadBalancer() *lsentity.LoadBalancer {
	return &lsentity.LoadBalancer{
		UUID: s.UUID,
	}
}

func (s *GetLoadBalancerByIdResponse) ToEntityLoadBalancer() *lsentity.LoadBalancer {
	if s == nil {
		return nil
	}

	return s.Data.toEntityLoadBalancer()
}

func (s *ListLoadBalancersResponse) ToEntityListLoadBalancers() *lsentity.ListLoadBalancers {
	if s == nil || s.ListData == nil || len(s.ListData) < 1 {
		return &lsentity.ListLoadBalancers{
			Page:      0,
			PageSize:  0,
			TotalPage: 0,
			TotalItem: 0,
			Items:     make([]*lsentity.LoadBalancer, 0),
		}
	}

	result := &lsentity.ListLoadBalancers{
		Page:      s.Page,
		PageSize:  s.PageSize,
		TotalPage: s.TotalPage,
		TotalItem: s.TotalItem,
	}

	for _, itemLb := range s.ListData {
		result.Add(itemLb.toEntityLoadBalancer())
	}

	return result
}

func (s *LoadBalancer) toEntityLoadBalancer() *lsentity.LoadBalancer {
	internal := lstr.TrimSpace(lstr.ToUpper(s.LoadBalancerSchema)) == "INTERNAL"
	return &lsentity.LoadBalancer{
		UUID:               s.UUID,
		Name:               s.Name,
		Address:            s.Address,
		DisplayStatus:      s.DisplayStatus,
		PrivateSubnetID:    s.PrivateSubnetID,
		PrivateSubnetCidr:  s.PrivateSubnetCidr,
		Type:               s.Type,
		DisplayType:        s.DisplayType,
		LoadBalancerSchema: s.LoadBalancerSchema,
		PackageID:          s.PackageID,
		Description:        s.Description,
		Location:           s.Location,
		CreatedAt:          s.CreatedAt,
		UpdatedAt:          s.UpdatedAt,
		ProgressStatus:     s.ProgressStatus,

		// will be removed
		Status:   s.DisplayStatus,
		Internal: internal,
		SubnetID: s.PrivateSubnetID,
	}
}
