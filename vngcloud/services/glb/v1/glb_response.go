package v1

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
)

type (
	GlobalLoadBalancerResponse struct {
		CreatedAt   string                             `json:"createdAt"`
		UpdatedAt   string                             `json:"updatedAt"`
		DeletedAt   string                             `json:"deletedAt"`
		ID          string                             `json:"id"`
		Name        string                             `json:"name"`
		Description string                             `json:"description"`
		Status      string                             `json:"status"`
		Package     string                             `json:"package"`
		Type        string                             `json:"type"`
		UserId      int                                `json:"userId"`
		Vips        []GlobalLoadBalancerVIPResponse    `json:"vips"`
		Domains     []GlobalLoadBalancerDomainResponse `json:"domains"`
	}

	GlobalLoadBalancerVIPResponse struct {
		ID                   int    `json:"id"`
		CreatedAt            string `json:"createdAt"`
		UpdatedAt            string `json:"updatedAt"`
		DeletedAt            string `json:"deletedAt"`
		Address              string `json:"address"`
		Status               string `json:"status"`
		Region               string `json:"region"`
		GlobalLoadBalancerID string `json:"globalLoadBalancerId"`
	}

	GlobalLoadBalancerDomainResponse struct {
		CreatedAt            string `json:"createdAt"`
		UpdatedAt            string `json:"updatedAt"`
		DeletedAt            string `json:"deletedAt"`
		ID                   int    `json:"id"`
		Hostname             string `json:"hostname"`
		Status               string `json:"status"`
		GlobalLoadBalancerID string `json:"globalLoadBalancerId"`
		DNSHostedZoneID      string `json:"dnsHostedZoneId"`
		DNSServerID          string `json:"dnsServerId"`
	}
)

func (s *GlobalLoadBalancerResponse) ToEntityGlobalLoadBalancer() *lsentity.GlobalLoadBalancer {
	vips := make([]*lsentity.GlobalLoadBalancerVIP, 0, len(s.Vips))
	for _, vip := range s.Vips {
		vips = append(vips, &lsentity.GlobalLoadBalancerVIP{
			ID:                   vip.ID,
			CreatedAt:            vip.CreatedAt,
			UpdatedAt:            vip.UpdatedAt,
			DeletedAt:            vip.DeletedAt,
			Address:              vip.Address,
			Status:               vip.Status,
			Region:               vip.Region,
			GlobalLoadBalancerID: vip.GlobalLoadBalancerID,
		})
	}

	domains := make([]*lsentity.GlobalLoadBalancerDomain, 0, len(s.Domains))
	for _, domain := range s.Domains {
		domains = append(domains, &lsentity.GlobalLoadBalancerDomain{
			CreatedAt:            domain.CreatedAt,
			UpdatedAt:            domain.UpdatedAt,
			DeletedAt:            domain.DeletedAt,
			ID:                   domain.ID,
			Hostname:             domain.Hostname,
			Status:               domain.Status,
			GlobalLoadBalancerID: domain.GlobalLoadBalancerID,
			DNSHostedZoneID:      domain.DNSHostedZoneID,
			DNSServerID:          domain.DNSServerID,
		})
	}

	return &lsentity.GlobalLoadBalancer{
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
		DeletedAt:   s.DeletedAt,
		ID:          s.ID,
		Name:        s.Name,
		Description: s.Description,
		Status:      s.Status,
		Package:     s.Package,
		Type:        s.Type,
		UserId:      s.UserId,
		Vips:        vips,
		Domains:     domains,
	}
}

// --------------------------------------------------
type ListGlobalLoadBalancersResponse struct {
	Items  []GlobalLoadBalancerResponse `json:"items"`
	Limit  int                          `json:"limit"`
	Total  int                          `json:"total"`
	Offset int                          `json:"offset"`
}

func (s *ListGlobalLoadBalancersResponse) ToEntityListGlobalLoadBalancers() *lsentity.ListGlobalLoadBalancers {
	result := &lsentity.ListGlobalLoadBalancers{
		Items:  make([]*lsentity.GlobalLoadBalancer, 0),
		Limit:  0,
		Total:  0,
		Offset: 0,
	}

	if s == nil || s.Items == nil || len(s.Items) < 1 {
		return result
	}

	result.Limit = s.Limit
	result.Total = s.Total
	result.Offset = s.Offset

	for _, itemLb := range s.Items {
		result.Items = append(result.Items, itemLb.ToEntityGlobalLoadBalancer())
	}

	return result
}

// --------------------------------------------------

type CreateGlobalLoadBalancerResponse struct {
	GlobalLoadBalancer GlobalLoadBalancerResponse `json:"globalLoadBalancer"`
	GlobalListener     GlobalListenerResponse     `json:"globalListener"`
	GlobalPool         GlobalPoolResponse         `json:"globalPool"`
}

func (s *CreateGlobalLoadBalancerResponse) ToEntityGlobalLoadBalancer() *lsentity.GlobalLoadBalancer {
	return s.GlobalLoadBalancer.ToEntityGlobalLoadBalancer()
}

// --------------------------------------------------

type GetGlobalLoadBalancerByIdResponse GlobalLoadBalancerResponse

func (s *GetGlobalLoadBalancerByIdResponse) ToEntityGlobalLoadBalancer() *lsentity.GlobalLoadBalancer {
	return (*GlobalLoadBalancerResponse)(s).ToEntityGlobalLoadBalancer()
}

// --------------------------------------------------

type ListGlobalPackagesResponse []GlobalPackageResponse

type GlobalPackageResponse struct {
	ID                          string                     `json:"id"`
	Name                        string                     `json:"name"`
	Description                 string                     `json:"description"`
	DescriptionEn               string                     `json:"descriptionEn"`
	Detail                      interface{}                `json:"detail"`
	Enabled                     bool                       `json:"enabled"`
	BaseSku                     string                     `json:"baseSku"`
	BaseConnectionRate          int                        `json:"baseConnectionRate"`
	BaseDomesticTrafficTotal    int                        `json:"baseDomesticTrafficTotal"`
	BaseNonDomesticTrafficTotal int                        `json:"baseNonDomesticTrafficTotal"`
	ConnectionSku               string                     `json:"connectionSku"`
	DomesticTrafficSku          string                     `json:"domesticTrafficSku"`
	NonDomesticTrafficSku       string                     `json:"nonDomesticTrafficSku"`
	CreatedAt                   string                     `json:"createdAt"`
	UpdatedAt                   string                     `json:"updatedAt"`
	VlbPackages                 []VlbGlobalPackageResponse `json:"vlbPackages"`
}

type VlbGlobalPackageResponse struct {
	ID           int    `json:"id"`
	GlbPackageID string `json:"glb_package_id"`
	Region       string `json:"region"`
	VlbPackageID string `json:"vlb_package_id"`
	CreatedAt    string `json:"created_at"`
}

func (s *ListGlobalPackagesResponse) ToEntityListGlobalPackages() *lsentity.ListGlobalPackages {
	packages := make([]lsentity.GlobalPackage, 0)
	if s != nil {
		for _, item := range *s {
			packages = append(packages, *item.ToEntityGlobalPackage())
		}
	}
	return &lsentity.ListGlobalPackages{Items: packages}
}

func (s *GlobalPackageResponse) ToEntityGlobalPackage() *lsentity.GlobalPackage {
	vlbPackages := make([]lsentity.VlbPackage, 0, len(s.VlbPackages))
	for _, vlb := range s.VlbPackages {
		vlbPackages = append(vlbPackages, lsentity.VlbPackage{
			ID:           vlb.ID,
			GlbPackageID: vlb.GlbPackageID,
			Region:       vlb.Region,
			VlbPackageID: vlb.VlbPackageID,
			CreatedAt:    vlb.CreatedAt,
		})
	}

	return &lsentity.GlobalPackage{
		ID:                          s.ID,
		Name:                        s.Name,
		Description:                 s.Description,
		DescriptionEn:               s.DescriptionEn,
		Detail:                      s.Detail,
		Enabled:                     s.Enabled,
		BaseSku:                     s.BaseSku,
		BaseConnectionRate:          s.BaseConnectionRate,
		BaseDomesticTrafficTotal:    s.BaseDomesticTrafficTotal,
		BaseNonDomesticTrafficTotal: s.BaseNonDomesticTrafficTotal,
		ConnectionSku:               s.ConnectionSku,
		DomesticTrafficSku:          s.DomesticTrafficSku,
		NonDomesticTrafficSku:       s.NonDomesticTrafficSku,
		CreatedAt:                   s.CreatedAt,
		UpdatedAt:                   s.UpdatedAt,
		VlbPackages:                 vlbPackages,
	}
}

// --------------------------------------------------

type ListGlobalRegionsResponse []GlobalRegionResponse

type GlobalRegionResponse struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	VServerEndpoint  string `json:"vserverEndpoint"`
	VlbEndpoint      string `json:"vlbEndpoint"`
	UIServerEndpoint string `json:"uiServerEndpoint"`
}

func (s *ListGlobalRegionsResponse) ToEntityListGlobalRegions() *lsentity.ListGlobalRegions {
	regions := make([]lsentity.GlobalRegion, 0)
	if s != nil {
		for _, item := range *s {
			regions = append(regions, *item.ToEntityGlobalRegion())
		}
	}
	return &lsentity.ListGlobalRegions{Items: regions}
}

func (s *GlobalRegionResponse) ToEntityGlobalRegion() *lsentity.GlobalRegion {
	return &lsentity.GlobalRegion{
		ID:               s.ID,
		Name:             s.Name,
		VServerEndpoint:  s.VServerEndpoint,
		VlbEndpoint:      s.VlbEndpoint,
		UIServerEndpoint: s.UIServerEndpoint,
	}
}

// --------------------------------------------------

type GetGlobalLoadBalancerUsageHistoriesResponse struct {
	Type  string                                   `json:"type"`
	Items []GlobalLoadBalancerUsageHistoryResponse `json:"items"`
	From  string                                   `json:"from"`
	To    string                                   `json:"to"`
}

type GlobalLoadBalancerUsageHistoryResponse struct {
	Timestamp string  `json:"timestamp"`
	Value     float64 `json:"value"`
	Type      string  `json:"type"`
}

func (s *GetGlobalLoadBalancerUsageHistoriesResponse) ToEntityGlobalLoadBalancerUsageHistories() *lsentity.ListGlobalLoadBalancerUsageHistories {
	histories := make([]lsentity.GlobalLoadBalancerUsageHistory, 0)
	if s != nil && s.Items != nil {
		for _, item := range s.Items {
			histories = append(histories, *item.ToEntityGlobalLoadBalancerUsageHistory())
		}
	}
	return &lsentity.ListGlobalLoadBalancerUsageHistories{
		Type:  s.Type,
		Items: histories,
		From:  s.From,
		To:    s.To,
	}
}

func (s *GlobalLoadBalancerUsageHistoryResponse) ToEntityGlobalLoadBalancerUsageHistory() *lsentity.GlobalLoadBalancerUsageHistory {
	return &lsentity.GlobalLoadBalancerUsageHistory{
		Timestamp: s.Timestamp,
		Value:     s.Value,
		Type:      s.Type,
	}
}
