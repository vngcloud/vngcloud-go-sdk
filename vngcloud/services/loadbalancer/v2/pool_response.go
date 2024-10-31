package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type GetPoolHealthMonitorByIdResponse struct {
	Data struct {
		UUID                string  `json:"uuid"`
		Timeout             int     `json:"timeout"`
		CreatedAt           string  `json:"createdAt"`
		UpdatedAt           string  `json:"updatedAt"`
		DomainName          *string `json:"domainName"`
		HttpVersion         *string `json:"httpVersion"`
		HealthCheckProtocol string  `json:"healthCheckProtocol"`
		Interval            int     `json:"interval"`
		HealthyThreshold    int     `json:"healthyThreshold"`
		UnhealthyThreshold  int     `json:"unhealthyThreshold"`
		HealthCheckMethod   *string `json:"healthCheckMethod"`
		HealthCheckPath     *string `json:"healthCheckPath"`
		SuccessCode         *string `json:"successCode"`
		ProgressStatus      string  `json:"progressStatus"`
		DisplayStatus       string  `json:"displayStatus"`
	} `json:"data"`
}

func (s *GetPoolHealthMonitorByIdResponse) ToEntityHealthMonitor() *lsentity.HealthMonitor {
	if s == nil {
		return nil
	}

	return &lsentity.HealthMonitor{
		Timeout:             s.Data.Timeout,
		CreatedAt:           s.Data.CreatedAt,
		UpdatedAt:           s.Data.UpdatedAt,
		DomainName:          s.Data.DomainName,
		HttpVersion:         s.Data.HttpVersion,
		HealthCheckProtocol: s.Data.HealthCheckProtocol,
		Interval:            s.Data.Interval,
		HealthyThreshold:    s.Data.HealthyThreshold,
		UnhealthyThreshold:  s.Data.UnhealthyThreshold,
		HealthCheckMethod:   s.Data.HealthCheckMethod,
		HealthCheckPath:     s.Data.HealthCheckPath,
		SuccessCode:         s.Data.SuccessCode,
		ProgressStatus:      s.Data.ProgressStatus,
		DisplayStatus:       s.Data.DisplayStatus,
	}
}

type CreatePoolResponse struct {
	UUID string `json:"uuid"`
}

type ListPoolsByLoadBalancerIdResponse struct {
	Data []Pool `json:"data"`
}

type ListPoolMembersResponse struct {
	Data []PoolMember `json:"data"`
}

type GetPoolByIdResponse struct {
	Data Pool `json:"data"`
}

type Pool struct {
	UUID              string       `json:"uuid"`
	Name              string       `json:"name"`
	Protocol          string       `json:"protocol"`
	Description       string       `json:"description,omitempty"`
	LoadBalanceMethod string       `json:"loadBalanceMethod"`
	DisplayStatus     string       `json:"displayStatus"`
	CreatedAt         string       `json:"createdAt"`
	UpdatedAt         string       `json:"updatedAt"`
	Stickiness        bool         `json:"stickiness"`
	TLSEncryption     bool         `json:"tlsEncryption"`
	ProgressStatus    string       `json:"progressStatus"`
	Members           []PoolMember `json:"members"`
}

type PoolMember struct {
	Address        string `json:"address"`
	Backup         bool   `json:"backup"`
	CreatedAt      string `json:"createdAt"`
	DisplayStatus  string `json:"displayStatus"`
	MonitorPort    int    `json:"monitorPort"`
	Name           string `json:"name"`
	PoolID         string `json:"poolId"`
	ProgressStatus string `json:"progressStatus"`
	ProtocolPort   int    `json:"protocolPort"`
	SubnetID       string `json:"subnetId"`
	TypeCreate     string `json:"typeCreate"`
	UpdateAt       string `json:"updateAt,omitempty"`
	UUID           string `json:"uuid"`
	Weight         int    `json:"weight"`
}

func (s *CreatePoolResponse) ToEntityPool() *lsentity.Pool {
	return &lsentity.Pool{
		UUID: s.UUID,
	}
}

func (s *ListPoolsByLoadBalancerIdResponse) ToEntityListPools() *lsentity.ListPools {
	listPools := new(lsentity.ListPools)
	for _, pool := range s.Data {
		listPools.Add(pool.toEntityPool())
	}

	return listPools
}

func (s *PoolMember) toEntityMember() *lsentity.Member {
	return &lsentity.Member{
		Address:        s.Address,
		Backup:         s.Backup,
		Name:           s.Name,
		UUID:           s.UUID,
		DisplayStatus:  s.DisplayStatus,
		ProtocolPort:   s.ProtocolPort,
		MonitorPort:    s.MonitorPort,
		SubnetID:       s.SubnetID,
		TypeCreate:     s.TypeCreate,
		CreatedAt:      s.CreatedAt,
		Weight:         s.Weight,
		PoolID:         s.PoolID,
		ProgressStatus: s.ProgressStatus,
	}
}

func (s *Pool) toEntityListMembers() *lsentity.ListMembers {
	listMembers := &lsentity.ListMembers{}
	for _, member := range s.Members {
		listMembers.Add(member.toEntityMember())
	}
	return listMembers
}

func (s *Pool) toEntityPool() *lsentity.Pool {
	return &lsentity.Pool{
		UUID:              s.UUID,
		Name:              s.Name,
		Protocol:          s.Protocol,
		Description:       s.Description,
		LoadBalanceMethod: s.LoadBalanceMethod,
		Status:            s.DisplayStatus,
		Stickiness:        s.Stickiness,
		TLSEncryption:     s.TLSEncryption,
		Members:           s.toEntityListMembers(),
	}
}

func (s *ListPoolMembersResponse) ToEntityListMembers() *lsentity.ListMembers {
	listMembers := &lsentity.ListMembers{}
	for _, member := range s.Data {
		listMembers.Add(member.toEntityMember())
	}
	return listMembers
}

func (s *GetPoolByIdResponse) ToEntityPool() *lsentity.Pool {
	return s.Data.toEntityPool()
}
