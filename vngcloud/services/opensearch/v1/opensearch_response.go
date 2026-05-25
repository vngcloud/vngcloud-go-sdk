package v1

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
)

// OpenSearchClusterResponse mirrors the wire payload for a single cluster.
type OpenSearchClusterResponse struct {
	Id          string `json:"id"`
	ClusterId   string `json:"clusterId"`
	ProjectId   string `json:"projectId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Version     string `json:"version"`
	Region      string `json:"region"`

	NumberOfNodes int32 `json:"numberOfNodes"`
	StorageSize   int32 `json:"storageSize"`

	EnableTls     bool `json:"enableTls"`
	PublicAccess  bool `json:"publicAccess"`
	PrivateAccess bool `json:"privateAccess"`
	EncryptVolume bool `json:"encryptVolume"`
	EnableVpcDns  bool `json:"enableVpcDns"`

	VpcId    string `json:"vpcId"`
	VpcCidr  string `json:"vpcCidr"`
	SubnetId string `json:"subnetId"`

	PackageId     string `json:"packageId"`
	StorageTypeId string `json:"storageTypeId"`
	ConfigGroupId string `json:"configGroupId"`
	PortalUserId  int32  `json:"portalUserId"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"deletedAt"`

	PackageDetail OpenSearchPackageDetailDto `json:"packageDetail"`
	StorageType   OpenSearchStorageTypeDto   `json:"storageType"`
	ConfigGroup   OpenSearchConfigGroupDto   `json:"configGroup"`
	BillingStatus OpenSearchBillingStatusDto `json:"billingStatus"`
}

type OpenSearchPackageDetailDto struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	Type               string `json:"type"`
	Description        string `json:"description"`
	NetworkPerformance string `json:"networkPerformance"`
	IsDefault          bool   `json:"isDefault"`
	Ram                int32  `json:"ram"`
	Cpu                int32  `json:"cpu"`
	Platform           string `json:"platform"`
}

type OpenSearchStorageTypeDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Min  int32  `json:"min"`
	Max  int32  `json:"max"`
}

type OpenSearchConfigGroupDto struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type OpenSearchBillingStatusDto struct {
	Status      string `json:"status"`
	RenewPeriod int32  `json:"renewPeriod"`
}

// ListClustersResponse wraps the list endpoint payload:
//
//	{ "listData": [...], "page": 0, "pageSize": 0, "totalPage": 0, "totalItem": 0 }
type ListClustersResponse struct {
	ListData  []OpenSearchClusterResponse `json:"listData"`
	Page      int32                       `json:"page"`
	PageSize  int32                       `json:"pageSize"`
	TotalPage int32                       `json:"totalPage"`
	TotalItem int32                       `json:"totalItem"`
}

// GetClusterResponse wraps { "data": {...} }.
type GetClusterResponse struct {
	Data OpenSearchClusterResponse `json:"data"`
}

func (s *OpenSearchClusterResponse) ToEntity() *lsentity.OpenSearchCluster {
	if s == nil {
		return nil
	}
	return &lsentity.OpenSearchCluster{
		Id:            s.Id,
		ClusterId:     s.ClusterId,
		ProjectId:     s.ProjectId,
		Name:          s.Name,
		Description:   s.Description,
		Status:        s.Status,
		Version:       s.Version,
		Region:        s.Region,
		NumberOfNodes: s.NumberOfNodes,
		StorageSize:   s.StorageSize,
		EnableTls:     s.EnableTls,
		PublicAccess:  s.PublicAccess,
		PrivateAccess: s.PrivateAccess,
		EncryptVolume: s.EncryptVolume,
		EnableVpcDns:  s.EnableVpcDns,
		VpcId:         s.VpcId,
		VpcCidr:       s.VpcCidr,
		SubnetId:      s.SubnetId,
		PackageId:     s.PackageId,
		StorageTypeId: s.StorageTypeId,
		ConfigGroupId: s.ConfigGroupId,
		PortalUserId:  s.PortalUserId,
		CreatedAt:     s.CreatedAt,
		UpdatedAt:     s.UpdatedAt,
		DeletedAt:     s.DeletedAt,
		PackageDetail: lsentity.OpenSearchPackageDetail{
			Id:                 s.PackageDetail.Id,
			Name:               s.PackageDetail.Name,
			Type:               s.PackageDetail.Type,
			Description:        s.PackageDetail.Description,
			NetworkPerformance: s.PackageDetail.NetworkPerformance,
			IsDefault:          s.PackageDetail.IsDefault,
			Ram:                s.PackageDetail.Ram,
			Cpu:                s.PackageDetail.Cpu,
			Platform:           s.PackageDetail.Platform,
		},
		StorageType: lsentity.OpenSearchStorageType{
			Id:   s.StorageType.Id,
			Name: s.StorageType.Name,
			Min:  s.StorageType.Min,
			Max:  s.StorageType.Max,
		},
		ConfigGroup: lsentity.OpenSearchConfigGroup{
			Name:    s.ConfigGroup.Name,
			Version: s.ConfigGroup.Version,
		},
		BillingStatus: lsentity.OpenSearchBillingStatus{
			Status:      s.BillingStatus.Status,
			RenewPeriod: s.BillingStatus.RenewPeriod,
		},
	}
}

func (s *ListClustersResponse) ToEntity() []lsentity.OpenSearchCluster {
	out := make([]lsentity.OpenSearchCluster, 0, len(s.ListData))
	for i := range s.ListData {
		out = append(out, *s.ListData[i].ToEntity())
	}
	return out
}

// OpenSearchEndpointResponse mirrors one element of the /endpoints payload.
type OpenSearchEndpointResponse struct {
	Url     string   `json:"url"`
	Type    string   `json:"type"` // "dashboard" (:443) | "log" (:9200)
	Cidrs   []string `json:"cidrs"`
	Private bool     `json:"private"`
}

// ListEndpointsResponse wraps { "data": [...] }.
type ListEndpointsResponse struct {
	Data []OpenSearchEndpointResponse `json:"data"`
}

func (s *OpenSearchEndpointResponse) ToEntity() *lsentity.OpenSearchEndpoint {
	if s == nil {
		return nil
	}
	return &lsentity.OpenSearchEndpoint{
		Url:     s.Url,
		Type:    s.Type,
		Cidrs:   s.Cidrs,
		Private: s.Private,
	}
}

func (s *ListEndpointsResponse) ToEntity() []lsentity.OpenSearchEndpoint {
	out := make([]lsentity.OpenSearchEndpoint, 0, len(s.Data))
	for i := range s.Data {
		out = append(out, *s.Data[i].ToEntity())
	}
	return out
}
