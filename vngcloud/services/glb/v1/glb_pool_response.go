package v1

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type GlobalPoolResponse struct {
	CreatedAt                 string                      `json:"createdAt"`
	UpdatedAt                 string                      `json:"updatedAt"`
	DeletedAt                 *string                     `json:"deletedAt"`
	ID                        string                      `json:"id"`
	Name                      string                      `json:"name"`
	Description               string                      `json:"description"`
	GlobalLoadBalancerID      string                      `json:"globalLoadBalancerId"`
	Algorithm                 string                      `json:"algorithm"`
	StickySession             *string                     `json:"stickySession"`
	TLSEnabled                *string                     `json:"tlsEnabled"`
	Protocol                  string                      `json:"protocol"`
	Status                    string                      `json:"status"`
	Health                    *HealthResponse             `json:"health"`
	GlobalPoolMembersResponse *[]GlobalPoolMemberResponse `json:"globalPoolMembers"`
}

func (s *GlobalPoolResponse) ToEntityPool() *lsentity.GlobalPool {
	return &lsentity.GlobalPool{
		CreatedAt:            s.CreatedAt,
		UpdatedAt:            s.UpdatedAt,
		DeletedAt:            s.DeletedAt,
		ID:                   s.ID,
		Name:                 s.Name,
		Description:          s.Description,
		GlobalLoadBalancerID: s.GlobalLoadBalancerID,
		Algorithm:            s.Algorithm,
		StickySession:        s.StickySession,
		TLSEnabled:           s.TLSEnabled,
		Protocol:             s.Protocol,
		Status:               s.Status,
		Health:               s.Health.ToEntityGlobalPoolHealthMonitor(),
	}
}

type HealthResponse struct {
	CreatedAt            string  `json:"createdAt"`
	UpdatedAt            string  `json:"updatedAt"`
	DeletedAt            *string `json:"deletedAt"`
	ID                   string  `json:"id"`
	GlobalPoolID         string  `json:"globalPoolId"`
	GlobalLoadBalancerID string  `json:"globalLoadBalancerId"`
	Protocol             string  `json:"protocol"`
	Path                 *string `json:"path"`
	Timeout              int     `json:"timeout"`
	IntervalTime         int     `json:"intervalTime"`
	HealthyThreshold     int     `json:"healthyThreshold"`
	UnhealthyThreshold   int     `json:"unhealthyThreshold"`
	HttpVersion          *string `json:"httpVersion"`
	HttpMethod           *string `json:"httpMethod"`
	DomainName           *string `json:"domainName"`
	SuccessCode          *string `json:"successCode"`
	Status               string  `json:"status"`
}

func (s *HealthResponse) ToEntityGlobalPoolHealthMonitor() *lsentity.GlobalPoolHealthMonitor {
	return &lsentity.GlobalPoolHealthMonitor{
		CreatedAt:            s.CreatedAt,
		UpdatedAt:            s.UpdatedAt,
		DeletedAt:            s.DeletedAt,
		ID:                   s.ID,
		GlobalPoolID:         s.GlobalPoolID,
		GlobalLoadBalancerID: s.GlobalLoadBalancerID,
		Protocol:             s.Protocol,
		Path:                 s.Path,
		Timeout:              s.Timeout,
		IntervalTime:         s.IntervalTime,
		HealthyThreshold:     s.HealthyThreshold,
		UnhealthyThreshold:   s.UnhealthyThreshold,
		DomainName:           s.DomainName,
		HTTPVersion:          s.HttpVersion,
		HTTPMethod:           s.HttpMethod,
		SuccessCode:          s.SuccessCode,
		Status:               s.Status,
	}
}

type GlobalPoolMemberResponse struct {
	CreatedAt            string                  `json:"createdAt"`
	UpdatedAt            string                  `json:"updatedAt"`
	DeletedAt            *string                 `json:"deletedAt"`
	ID                   string                  `json:"id"`
	Name                 string                  `json:"name"`
	Description          string                  `json:"description"`
	Region               string                  `json:"region"`
	GlobalPoolID         string                  `json:"globalPoolId"`
	GlobalLoadBalancerID string                  `json:"globalLoadBalancerId"`
	TrafficDial          int                     `json:"trafficDial"`
	VpcID                string                  `json:"vpcId"`
	Type                 string                  `json:"type"`
	Status               string                  `json:"status"`
	Members              []*GlobalMemberResponse `json:"members"`
}

func (s *GlobalPoolMemberResponse) ToEntityGlobalPoolMember() *lsentity.GlobalPoolMember {
	members := make([]*lsentity.GlobalPoolMemberDetail, 0, len(s.Members))
	for _, member := range s.Members {
		members = append(members, member.ToEntityGlobalMember())
	}
	return &lsentity.GlobalPoolMember{
		CreatedAt:            s.CreatedAt,
		UpdatedAt:            s.UpdatedAt,
		DeletedAt:            s.DeletedAt,
		ID:                   s.ID,
		Name:                 s.Name,
		Description:          s.Description,
		Region:               s.Region,
		GlobalPoolID:         s.GlobalPoolID,
		GlobalLoadBalancerID: s.GlobalLoadBalancerID,
		TrafficDial:          s.TrafficDial,
		VpcID:                s.VpcID,
		Type:                 s.Type,
		Status:               s.Status,
		Members:              &lsentity.ListGlobalMembers{Items: members},
	}
}

type GlobalMemberResponse struct {
	CreatedAt            string  `json:"createdAt"`
	UpdatedAt            string  `json:"updatedAt"`
	DeletedAt            *string `json:"deletedAt"`
	ID                   string  `json:"id"`
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	GlobalPoolMemberID   string  `json:"globalPoolMemberId"`
	GlobalLoadBalancerID string  `json:"globalLoadBalancerId"`
	SubnetID             string  `json:"subnetId"`
	Address              string  `json:"address"`
	Weight               int     `json:"weight"`
	Port                 int     `json:"port"`
	MonitorPort          int     `json:"monitorPort"`
	BackupRole           bool    `json:"backupRole"`
	Status               string  `json:"status"`
}

func (s *GlobalMemberResponse) ToEntityGlobalMember() *lsentity.GlobalPoolMemberDetail {
	return &lsentity.GlobalPoolMemberDetail{
		CreatedAt:            s.CreatedAt,
		UpdatedAt:            s.UpdatedAt,
		DeletedAt:            s.DeletedAt,
		ID:                   s.ID,
		Name:                 s.Name,
		Description:          s.Description,
		GlobalPoolMemberID:   s.GlobalPoolMemberID,
		GlobalLoadBalancerID: s.GlobalLoadBalancerID,
		SubnetID:             s.SubnetID,
		Address:              s.Address,
		Weight:               s.Weight,
		Port:                 s.Port,
		MonitorPort:          s.MonitorPort,
		BackupRole:           s.BackupRole,
		Status:               s.Status,
	}
}

// ----------------------------------------------------------

type ListGlobalPoolsResponse []*GlobalPoolResponse

func (s *ListGlobalPoolsResponse) ToEntityListGlobalPools() *lsentity.ListGlobalPools {
	result := &lsentity.ListGlobalPools{
		Items: make([]*lsentity.GlobalPool, 0),
	}

	if s == nil || len(*s) < 1 {
		return result
	}

	for _, pool := range *s {
		result.Items = append(result.Items, pool.ToEntityPool())
	}

	return result
}

// ----------------------------------------------------------

type CreateGlobalPoolResponse struct {
	ID                   string                      `json:"id"`
	Name                 string                      `json:"name"`
	Description          string                      `json:"description"`
	GlobalLoadBalancerID string                      `json:"globalLoadBalancerId"`
	Algorithm            string                      `json:"algorithm"`
	StickySession        *string                     `json:"stickySession"`
	TLSEnabled           *string                     `json:"tlsEnabled"`
	Protocol             string                      `json:"protocol"`
	Health               *HealthResponse             `json:"health"`
	GlobalPoolMembers    []*GlobalPoolMemberResponse `json:"globalPoolMembers"`
}

func (s *CreateGlobalPoolResponse) ToEntityPool() *lsentity.GlobalPool {
	return &lsentity.GlobalPool{
		ID:                   s.ID,
		Name:                 s.Name,
		Description:          s.Description,
		GlobalLoadBalancerID: s.GlobalLoadBalancerID,
		Algorithm:            s.Algorithm,
		StickySession:        s.StickySession,
		TLSEnabled:           s.TLSEnabled,
		Protocol:             s.Protocol,
	}
}

// ----------------------------------------------------------

type UpdateGlobalPoolResponse struct {
	ID string `json:"id"`
}

func (s *UpdateGlobalPoolResponse) ToEntityPool() *lsentity.GlobalPool {
	return &lsentity.GlobalPool{
		ID: s.ID,
	}
}

// ----------------------------------------------------------

type ListGlobalPoolMembersResponse []*GlobalPoolMemberResponse

func (s *ListGlobalPoolMembersResponse) ToEntityListGlobalPoolMembers() *lsentity.ListGlobalPoolMembers {
	result := &lsentity.ListGlobalPoolMembers{
		Items: make([]*lsentity.GlobalPoolMember, 0),
	}

	if s == nil || len(*s) < 1 {
		return result
	}

	for _, member := range *s {
		result.Items = append(result.Items, member.ToEntityGlobalPoolMember())
	}

	return result
}

// ----------------------------------------------------------

type GetGlobalPoolMemberResponse GlobalPoolMemberResponse

func (s *GetGlobalPoolMemberResponse) ToEntityGlobalPoolMember() *lsentity.GlobalPoolMember {
	members := make([]*lsentity.GlobalPoolMemberDetail, 0, len(s.Members))
	for _, member := range s.Members {
		members = append(members, member.ToEntityGlobalMember())
	}
	return &lsentity.GlobalPoolMember{
		CreatedAt:            s.CreatedAt,
		UpdatedAt:            s.UpdatedAt,
		DeletedAt:            s.DeletedAt,
		ID:                   s.ID,
		Name:                 s.Name,
		Description:          s.Description,
		Region:               s.Region,
		GlobalPoolID:         s.GlobalPoolID,
		GlobalLoadBalancerID: s.GlobalLoadBalancerID,
		TrafficDial:          s.TrafficDial,
		VpcID:                s.VpcID,
		Type:                 s.Type,
		Status:               s.Status,
		Members:              &lsentity.ListGlobalMembers{Items: members},
	}
}

// ----------------------------------------------------------

type UpdateGlobalPoolMemberResponse struct {
	ID                   string `json:"id"`
	GlobalPoolID         string `json:"globalPoolId"`
	GlobalLoadBalancerID string `json:"globalLoadBalancerId"`
	TrafficDial          int    `json:"trafficDial"`
}

func (s *UpdateGlobalPoolMemberResponse) ToEntityGlobalPoolMember() *lsentity.GlobalPoolMember {
	return &lsentity.GlobalPoolMember{
		ID:                   s.ID,
		GlobalPoolID:         s.GlobalPoolID,
		GlobalLoadBalancerID: s.GlobalLoadBalancerID,
		TrafficDial:          s.TrafficDial,
	}
}

// ----------------------------------------------------------
