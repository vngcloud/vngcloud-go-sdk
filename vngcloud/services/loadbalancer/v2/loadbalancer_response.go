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
	internal := lstr.TrimSpace(lstr.ToUpper(s.Data.LoadBalancerSchema)) == "INTERNAL"
	return &lsentity.LoadBalancer{
		UUID:               s.Data.UUID,
		Name:               s.Data.Name,
		Address:            s.Data.Address,
		DisplayStatus:      s.Data.DisplayStatus,
		PrivateSubnetID:    s.Data.PrivateSubnetID,
		PrivateSubnetCidr:  s.Data.PrivateSubnetCidr,
		Type:               s.Data.Type,
		DisplayType:        s.Data.DisplayType,
		LoadBalancerSchema: s.Data.LoadBalancerSchema,
		PackageID:          s.Data.PackageID,
		Description:        s.Data.Description,
		Location:           s.Data.Location,
		CreatedAt:          s.Data.CreatedAt,
		UpdatedAt:          s.Data.UpdatedAt,
		ProgressStatus:     s.Data.ProgressStatus,

		// will be removed
		Status:   s.Data.DisplayStatus,
		Internal: internal,
		SubnetID: s.Data.PrivateSubnetID,
	}
}
