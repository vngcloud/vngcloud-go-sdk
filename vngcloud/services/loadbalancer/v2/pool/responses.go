package pool

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

type CreateResponse struct {
	UUID string `json:"uuid"`
}

func (p *CreateResponse) ToPoolObject() *objects.Pool {
	return &objects.Pool{
		UUID: p.UUID,
	}
}

// **************************************** ListPoolsBasedLoadBalancerResponse *****************************************
type PoolResponse struct {
	UUID              string           `json:"uuid"`
	Name              string           `json:"name"`
	Protocol          string           `json:"protocol"`
	Description       string           `json:"description,omitempty"`
	LoadBalanceMethod string           `json:"loadBalanceMethod"`
	DisplayStatus     string           `json:"displayStatus"`
	CreatedAt         string           `json:"createdAt"`
	UpdatedAt         string           `json:"updatedAt"`
	Stickiness        bool             `json:"stickiness"`
	TLSEncryption     bool             `json:"tlsEncryption"`
	ProgressStatus    string           `json:"progressStatus"`
	Members           []MemberResponse `json:"members"`
}
type MemberResponse struct {
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
type ListPoolsBasedLoadBalancerResponse struct {
	Data []PoolResponse `json:"data"`
}

func (s *ListPoolsBasedLoadBalancerResponse) ToListPoolObjects() []*objects.Pool {
	var pools []*objects.Pool

	if s == nil || s.Data == nil || len(s.Data) == 0 {
		return pools
	}

	for _, item := range s.Data {
		var members []objects.Member
		if item.Members != nil && len(item.Members) > 0 {
			for _, member := range item.Members {
				members = append(members, objects.Member{
					Address:        member.Address,
					Backup:         member.Backup,
					Name:           member.Name,
					UUID:           member.UUID,
					DisplayStatus:  member.DisplayStatus,
					ProtocolPort:   member.ProtocolPort,
					MonitorPort:    member.MonitorPort,
					SubnetID:       member.SubnetID,
					TypeCreate:     member.TypeCreate,
					CreatedAt:      member.CreatedAt,
					Weight:         member.Weight,
					PoolID:         member.PoolID,
					ProgressStatus: member.ProgressStatus,
				})
			}
		}

		pools = append(pools, &objects.Pool{
			UUID:              item.UUID,
			Name:              item.Name,
			Protocol:          item.Protocol,
			Description:       item.Description,
			LoadBalanceMethod: item.LoadBalanceMethod,
			Status:            item.DisplayStatus,
			Stickiness:        item.Stickiness,
			TLSEncryption:     item.TLSEncryption,
			Members:           members,
		})
	}
	return pools
}

type GetResponse struct {
	Data PoolResponse `json:"data"`
}

func (s *GetResponse) ToPoolObject() *objects.Pool {
	if s == nil {
		return nil
	}

	var members []objects.Member
	if s.Data.Members != nil && len(s.Data.Members) > 0 {
		for _, member := range s.Data.Members {
			members = append(members, objects.Member{
				Address:        member.Address,
				Backup:         member.Backup,
				Name:           member.Name,
				UUID:           member.UUID,
				DisplayStatus:  member.DisplayStatus,
				ProtocolPort:   member.ProtocolPort,
				MonitorPort:    member.MonitorPort,
				SubnetID:       member.SubnetID,
				TypeCreate:     member.TypeCreate,
				CreatedAt:      member.CreatedAt,
				Weight:         member.Weight,
				PoolID:         member.PoolID,
				ProgressStatus: member.ProgressStatus,
			})
		}
	}

	return &objects.Pool{
		UUID:              s.Data.UUID,
		Name:              s.Data.Name,
		Protocol:          s.Data.Protocol,
		Description:       s.Data.Description,
		LoadBalanceMethod: s.Data.LoadBalanceMethod,
		Status:            s.Data.DisplayStatus,
		Stickiness:        s.Data.Stickiness,
		TLSEncryption:     s.Data.TLSEncryption,
		Members:           members,
	}
}

type GetMemberResponse struct {
	Data []MemberResponse `json:"data"`
}

func (s *GetMemberResponse) ToListMemberObject() []*objects.Member {
	var members []*objects.Member

	if s == nil || s.Data == nil || len(s.Data) == 0 {
		return members
	}

	for _, item := range s.Data {
		members = append(members, &objects.Member{
			Address:        item.Address,
			Backup:         item.Backup,
			Name:           item.Name,
			UUID:           item.UUID,
			DisplayStatus:  item.DisplayStatus,
			ProtocolPort:   item.ProtocolPort,
			MonitorPort:    item.MonitorPort,
			SubnetID:       item.SubnetID,
			TypeCreate:     item.TypeCreate,
			CreatedAt:      item.CreatedAt,
			Weight:         item.Weight,
			PoolID:         item.PoolID,
			ProgressStatus: item.ProgressStatus,
		})
	}
	return members
}
