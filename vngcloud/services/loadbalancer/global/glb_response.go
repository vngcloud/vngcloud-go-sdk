package global

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
	vips := make([]*lsentity.GlobalLoadBalancerVIP, 0)
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

	domains := make([]*lsentity.GlobalLoadBalancerDomain, 0)
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
		result.Items = append(result.Items, itemLb.toEntityGlobalLoadBalancer())
	}

	return result
}

func (s *GlobalLoadBalancerResponse) toEntityGlobalLoadBalancer() *lsentity.GlobalLoadBalancer {
	vips := make([]*lsentity.GlobalLoadBalancerVIP, 0)
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

	domains := make([]*lsentity.GlobalLoadBalancerDomain, 0)
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

type CreateGlobalLoadBalancerResponse struct {
	GlobalLoadBalancer GlobalLoadBalancerResponse `json:"globalLoadBalancer"`
	GlobalListener     GlobalListenerResponse     `json:"globalListener"`
	GlobalPool         GlobalPoolResponse         `json:"globalPool"`
}

func (s *CreateGlobalLoadBalancerResponse) ToEntityGlobalLoadBalancer() *lsentity.GlobalLoadBalancer {
	return s.GlobalLoadBalancer.ToEntityGlobalLoadBalancer()
}
