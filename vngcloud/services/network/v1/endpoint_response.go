package v1

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type GetEndpointByIdResponse struct {
	Data struct {
		Uuid              string `json:"uuid,omitempty"`
		EndpointName      string `json:"endpointName,omitempty"`
		EndpointServiceId string `json:"endpointServiceId,omitempty"`
		VpcId             string `json:"vpcId,omitempty"`
		EndpointUrl       string `json:"endpointUrl,omitempty"`
		Status            string `json:"status,omitempty"`
		Service           struct {
			Uuid       string `json:"uuid,omitempty"`
			Name       string `json:"name,omitempty"`
			EndpointIp string `json:"endpointIp,omitempty"`
		} `json:"service,omitempty"`
	} `json:"data"`
}

func (s *GetEndpointByIdResponse) ToEntityEndpoint() *lsentity.Endpoint {
	return &lsentity.Endpoint{
		Id:          s.Data.Uuid,
		Name:        s.Data.EndpointName,
		VpcId:       s.Data.VpcId,
		IPv4Address: s.Data.Service.EndpointIp,
		EndpointUrl: s.Data.EndpointUrl,
		Status:      s.Data.Status,
	}
}
