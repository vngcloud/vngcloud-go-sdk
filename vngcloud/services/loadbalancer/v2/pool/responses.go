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

type ListPoolsBasedLoadBalancerResponse struct {
	Data []struct {
		UUID              string `json:"uuid"`
		Name              string `json:"name"`
		Protocol          string `json:"protocol"`
		Description       string `json:"description,omitempty"`
		LoadBalanceMethod string `json:"loadBalanceMethod"`
		DisplayStatus     string `json:"displayStatus"`
		CreatedAt         string `json:"createdAt"`
		UpdatedAt         string `json:"updatedAt"`
		Stickiness        bool   `json:"stickiness"`
		TLSEncryption     bool   `json:"tlsEncryption"`
		ProgressStatus    string `json:"progressStatus"`
		Members           []struct {
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
		} `json:"members"`
	} `json:"data"`
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
					Address: member.Address,
					Backup:  member.Backup,
					Status:  member.DisplayStatus,
					Name:    member.Name,
					UUID:    member.UUID,
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
