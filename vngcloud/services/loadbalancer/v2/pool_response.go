package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

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
