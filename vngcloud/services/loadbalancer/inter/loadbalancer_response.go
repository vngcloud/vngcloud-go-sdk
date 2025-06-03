package inter

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type CreateLoadBalancerResponse struct {
	UUID string `json:"uuid"`
}

func (s *CreateLoadBalancerResponse) ToEntityLoadBalancer() *lsentity.LoadBalancer {
	return &lsentity.LoadBalancer{
		UUID: s.UUID,
	}
}
