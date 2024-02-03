package network

import "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"

type GetResponse struct {
	Status      string   `json:"status"`
	ElasticIps  []string `json:"elasticIps"`
	DisplayName string   `json:"displayName"`
	ID          string   `json:"id"`
	CreatedAt   string   `json:"createdAt"`
	Cidr        string   `json:"cidr"`
}

func (s *GetResponse) ToNetworkObject() *objects.Network {
	if s == nil {
		return nil
	}

	return &objects.Network{
		Status:     s.Status,
		ElasticIps: s.ElasticIps,
		Name:       s.DisplayName,
		UUID:       s.ID,
		CreatedAt:  s.CreatedAt,
		Cidr:       s.Cidr,
	}
}
