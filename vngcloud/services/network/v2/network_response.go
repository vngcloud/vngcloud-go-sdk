package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type GetNetworkByIdResponse struct {
	Status      string   `json:"status"`
	ElasticIps  []string `json:"elasticIps"`
	DisplayName string   `json:"displayName"`
	ID          string   `json:"id"`
	CreatedAt   string   `json:"createdAt"`
	Cidr        string   `json:"cidr"`
}

func (s *GetNetworkByIdResponse) ToEntityNetwork() *lsentity.Network {
	return &lsentity.Network{
		Status:     s.Status,
		ElasticIps: s.ElasticIps,
		Name:       s.DisplayName,
		Id:         s.ID,
		CreatedAt:  s.CreatedAt,
		Cidr:       s.Cidr,
	}
}
