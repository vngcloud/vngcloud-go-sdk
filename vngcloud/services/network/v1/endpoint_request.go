package v1

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
)

func NewGetEndpointByIdRequest(pendpointId string) IGetEndpointByIdRequest {
	opt := new(GetEndpointByIdRequest)
	opt.EndpointId = pendpointId
	return opt
}

type GetEndpointByIdRequest struct {
	lscommon.UserAgent
	lscommon.EndpointCommon
}

func (s *GetEndpointByIdRequest) AddUserAgent(pagent ...string) IGetEndpointByIdRequest {
	s.UserAgent.Agent = append(s.UserAgent.Agent, pagent...)
	return s
}

type CreateEndpointRequest struct {
	ResourceType string `json:"resourceType"`
	Action       string `json:"action"`
	ResourceInfo struct {
		IsBuyMorePoc      bool   `json:"isBuyMorePoc"`
		IsPoc             bool   `json:"isPoc"`
		IsEnableAutoRenew bool   `json:"isEnableAutoRenew"`
		EndpointName      string `json:"endpointName"`
		CategoryUuid      string `json:"categoryUuid"`
		ServiceUuid       string `json:"serviceUuid"`
		PackageUuid       string `json:"packageUuid"`
		VpcUuid           string `json:"vpcUuid"`
		SubnetUuid        string `json:"subnetUuid"`
		RegionUuid        string `json:"regionUuid"`
		ProjectUuid       string `json:"projectUuid"`
		Description       string `json:"description"`
	} `json:"resourceInfo"`

	lscommon.UserAgent
}

func (s *CreateEndpointRequest) AddUserAgent(pagent ...string) ICreateEndpointRequest {
	s.UserAgent.Agent = append(s.UserAgent.Agent, pagent...)
	return s
}

func (s *CreateEndpointRequest) ToRequestBody(psvc lsclient.IServiceClient) interface{} {
	s.ResourceType = "endpoint"
	s.Action = "create"
	s.ResourceInfo.RegionUuid = psvc.GetZoneId()
	s.ResourceInfo.ProjectUuid = psvc.GetProjectId()

	return s
}

func (s *CreateEndpointRequest) WithEndpointName(pendpointName string) ICreateEndpointRequest {
	s.ResourceInfo.EndpointName = pendpointName
	return s
}

func (s *CreateEndpointRequest) WithCategoryUuid(pcategoryUuid string) ICreateEndpointRequest {
	s.ResourceInfo.CategoryUuid = pcategoryUuid
	return s
}

func (s *CreateEndpointRequest) WithServiceUuid(pserviceUuid string) ICreateEndpointRequest {
	s.ResourceInfo.ServiceUuid = pserviceUuid
	return s
}

func (s *CreateEndpointRequest) WithPackageUuid(ppackageUuid string) ICreateEndpointRequest {
	s.ResourceInfo.PackageUuid = ppackageUuid
	return s
}

func (s *CreateEndpointRequest) WithVpcUuid(pvpcUuid string) ICreateEndpointRequest {
	s.ResourceInfo.VpcUuid = pvpcUuid
	return s
}

func (s *CreateEndpointRequest) WithSubnetUuid(psubnetUuid string) ICreateEndpointRequest {
	s.ResourceInfo.SubnetUuid = psubnetUuid
	return s
}

func (s *CreateEndpointRequest) WithDescription(pdesp string) ICreateEndpointRequest {
	s.ResourceInfo.Description = pdesp
	return s
}
