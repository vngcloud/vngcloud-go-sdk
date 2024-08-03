package v1

import (
	lfmt "fmt"
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
	lurl "net/url"
)

const (
	defaultPackageId = "d33c1f28-cb27-442d-81ac-062f57bd52b9"

	vstorageServiceId = "b9ba2b16-389e-48b7-9e75-4c991239da27"

	defaultListEndpointsRequestPage = 1
	defaultListEndpointsRequestSize = 10
)

func NewGetEndpointByIdRequest(pendpointId string) IGetEndpointByIdRequest {
	opt := new(GetEndpointByIdRequest)
	opt.EndpointId = pendpointId
	return opt
}

func NewCreateEndpointRequest(pname, pserviceId, pvpcId, psubnetId string) ICreateEndpointRequest {
	opts := new(CreateEndpointRequest)
	opts.ResourceInfo.EndpointName = pname
	opts.ResourceInfo.ServiceUuid = pserviceId
	opts.ResourceInfo.VpcUuid = pvpcId
	opts.ResourceInfo.SubnetUuid = psubnetId

	return opts
}

func NewDeleteEndpointByIdRequest(pendpointId, pvpcId, pendpointServiceId string) IDeleteEndpointByIdRequest {
	opt := new(DeleteEndpointByIdRequest)
	opt.EndpointId = pendpointId
	opt.EndpointUuid = pendpointId
	opt.VpcUuid = pvpcId
	opt.EndpointServiceUuid = pendpointServiceId

	return opt
}

func NewListEndpointsRequest(ppage, psize int) IListEndpointsRequest {
	return &ListEndpointsRequest{
		Page: ppage,
		Size: psize,
	}
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
	s.ResourceInfo.PackageUuid = defaultPackageId

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

type DeleteEndpointByIdRequest struct {
	EndpointServiceUuid string `json:"endpointServiceUuid"`
	EndpointUuid        string `json:"endpointUuid"`
	ProjectUuid         string `json:"projectUuid"`
	RegionUuid          string `json:"regionUuid"`
	VpcUuid             string `json:"vpcUuid"`

	lscommon.UserAgent
	lscommon.EndpointCommon
}

func (s *DeleteEndpointByIdRequest) AddUserAgent(pagent ...string) IDeleteEndpointByIdRequest {
	s.UserAgent.Agent = append(s.UserAgent.Agent, pagent...)
	return s
}

func (s *DeleteEndpointByIdRequest) ToRequestBody(psvc lsclient.IServiceClient) interface{} {
	s.ProjectUuid = psvc.GetProjectId()
	s.RegionUuid = psvc.GetZoneId()

	return s
}

type ListEndpointsRequest struct {
	Page  int
	Size  int
	VpcId string
	lscommon.UserAgent
}

func (s *ListEndpointsRequest) WithPage(ppage int) IListEndpointsRequest {
	s.Page = ppage
	return s
}

func (s *ListEndpointsRequest) WithSize(psize int) IListEndpointsRequest {
	s.Size = psize
	return s
}

func (s *ListEndpointsRequest) WithVpcId(pvpcId string) IListEndpointsRequest {
	s.VpcId = pvpcId
	return s
}

func (s *ListEndpointsRequest) ToListQuery() (string, error) {
	params := ""
	if s.VpcId != "" {
		params = lfmt.Sprintf(`{"field":"vpcId","value":"%s"}`, s.VpcId)
	}

	query := lfmt.Sprintf(`{"page":%d,"size":%d,"search":[%s]}`, s.Page, s.Size, params)
	query = "params=" + lurl.QueryEscape(query)

	return query, nil
}

func (s *ListEndpointsRequest) GetDefaultQuery() string {
	query := lfmt.Sprintf(`{"page":%d,"size":%d}`, defaultListEndpointsRequestPage, defaultListEndpointsRequestSize)
	query = "params=" + lurl.QueryEscape(query)
	return query
}

func (s *ListEndpointsRequest) AddUserAgent(pagent ...string) IListEndpointsRequest {
	s.UserAgent.Agent = append(s.UserAgent.Agent, pagent...)
	return s
}
