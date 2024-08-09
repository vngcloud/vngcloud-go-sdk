package v1

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type Endpoint struct {
	Uuid              string `json:"uuid,omitempty"`
	EndpointName      string `json:"endpointName,omitempty"`
	EndpointServiceId string `json:"endpointServiceId,omitempty"`
	VpcId             string `json:"vpcId,omitempty"`
	EndpointUrl       string `json:"endpointUrl,omitempty"`
	EndpointIp        string `json:"endpointIp,omitempty"`
	Status            string `json:"status,omitempty"`
}

func (s *Endpoint) toEntityEndpoint() *lsentity.Endpoint {
	return &lsentity.Endpoint{
		Id:          s.Uuid,
		Name:        s.EndpointName,
		VpcId:       s.VpcId,
		IPv4Address: s.EndpointIp,
		EndpointUrl: s.EndpointUrl,
		Status:      s.Status,
	}
}

type GetEndpointByIdResponse struct {
	Data Endpoint `json:"data"`
}

func (s *GetEndpointByIdResponse) ToEntityEndpoint() *lsentity.Endpoint {
	return s.Data.toEntityEndpoint()
}

type CreateEndpointResponse struct {
	Data struct {
		Uuid string `json:"uuid,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"data"`
}

func (s *CreateEndpointResponse) ToEntityEndpoint() *lsentity.Endpoint {
	return &lsentity.Endpoint{
		Id:   s.Data.Uuid,
		Name: s.Data.Name,
	}
}

type ListEndpointsResponse struct {
	Data      []Endpoint `json:"data"`
	Page      int        `json:"page"`
	Size      int        `json:"size"`
	TotalPage int        `json:"totalPage"`
	Total     int        `json:"total"`
}

func (s *ListEndpointsResponse) ToEntityListEndpoints() *lsentity.ListEndpoints {
	var items []*lsentity.Endpoint
	for _, item := range s.Data {
		items = append(items, item.toEntityEndpoint())
	}
	return &lsentity.ListEndpoints{
		Items:     items,
		Page:      s.Page,
		PageSize:  s.Size,
		TotalPage: s.TotalPage,
		TotalItem: s.Total,
	}
}
