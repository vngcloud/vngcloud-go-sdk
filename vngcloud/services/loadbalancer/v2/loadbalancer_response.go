package v2

import (
	lstr "strings"

	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
)

type CreateLoadBalancerResponse struct {
	UUID string `json:"uuid"`
}

type ResizeLoadBalancerResponse struct {
	UUID string `json:"uuid"`
}

func (s *ResizeLoadBalancerResponse) ToEntityLoadBalancer() *lsentity.LoadBalancer {
	return &lsentity.LoadBalancer{
		UUID: s.UUID,
	}
}

type ScaleLoadBalancerResponse struct {
	UUID string `json:"uuid"`
}

func (s *ScaleLoadBalancerResponse) ToEntityLoadBalancer() *lsentity.LoadBalancer {
	return &lsentity.LoadBalancer{
		UUID: s.UUID,
	}
}

type ListLoadBalancerPackagesResponse struct {
	ListData []LoadBalancerPackageResponse `json:"listData"`
}

type LoadBalancerPackageResponse struct {
	UUID             string `json:"uuid"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	ConnectionNumber int    `json:"connectionNumber"`
	DataTransfer     int    `json:"dataTransfer"`
	Mode             string `json:"mode"`
	LbType           string `json:"lbType"`
	DisplayLbType    string `json:"displayLbType"`
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
		BackendSubnetID    string `json:"backendSubnetId"`
		PackageInfo        struct {
			PackageID        string `json:"packageId"`
			ConnectionNumber int    `json:"connectionNumber"`
			DataTransfer     int    `json:"dataTransfer"`
			Name             string `json:"name"`
		} `json:"packageInfo"`
		ProgressStatus string `json:"progressStatus"`
		AutoScalable   bool   `json:"autoScalable"`
		Zone           struct {
			UUID string `json:"uuid"`
			// Name        string `json:"name"`
			// ZoneType    string `json:"zoneType"`
			// IsDefault   bool   `json:"isDefault"`
			// Description string `json:"description"`
			// IsEnabled   bool   `json:"isEnabled"`
			// VolumeCount int    `json:"volumeCount"`
			// ServerCount int    `json:"serverCount"`
		} `json:"zone"`
		MinSize    int    `json:"minSize"`
		MaxSize    int    `json:"maxSize"`
		TotalNodes int    `json:"totalNodes"`
		Nodes      []Node `json:"nodes"`
	}

	Node struct {
		Status   string `json:"status"`
		ZoneID   string `json:"zoneId"`
		ZoneName string `json:"zoneName"`
		SubnetID string `json:"subnetId"`
	}
)

func (s *CreateLoadBalancerResponse) ToEntityLoadBalancer() *lsentity.LoadBalancer {
	return &lsentity.LoadBalancer{
		UUID: s.UUID,
	}
}

func (s *ListLoadBalancerPackagesResponse) ToEntityListLoadBalancerPackages() *lsentity.ListLoadBalancerPackages {
	if s == nil || s.ListData == nil || len(s.ListData) < 1 {
		return &lsentity.ListLoadBalancerPackages{
			Items: make([]*lsentity.LoadBalancerPackage, 0),
		}
	}

	result := &lsentity.ListLoadBalancerPackages{
		Items: make([]*lsentity.LoadBalancerPackage, 0),
	}

	for _, item := range s.ListData {
		result.Items = append(result.Items, &lsentity.LoadBalancerPackage{
			UUID:             item.UUID,
			Name:             item.Name,
			Type:             item.Type,
			ConnectionNumber: item.ConnectionNumber,
			DataTransfer:     item.DataTransfer,
			Mode:             item.Mode,
			LbType:           item.LbType,
			DisplayLbType:    item.DisplayLbType,
		})
	}

	return result
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

	// Convert nodes from response to entity
	nodes := make([]*lsentity.Node, 0, len(s.Nodes))
	for _, node := range s.Nodes {
		nodes = append(nodes, &lsentity.Node{
			Status:   node.Status,
			ZoneID:   node.ZoneID,
			ZoneName: node.ZoneName,
			SubnetID: node.SubnetID,
		})
	}

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
		AutoScalable:       s.AutoScalable,
		ZoneID:             s.Zone.UUID,
		MinSize:            s.MinSize,
		MaxSize:            s.MaxSize,
		TotalNodes:         s.TotalNodes,
		Nodes:              nodes,
		BackendSubnetID: func() string {
			if s.BackendSubnetID != "" {
				return s.BackendSubnetID
			}
			return s.PrivateSubnetID
		}(),

		// will be removed
		Status:   s.DisplayStatus,
		Internal: internal,
	}
}
