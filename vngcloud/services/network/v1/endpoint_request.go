package v1

import (
	lfmt "fmt"
	lurl "net/url"
	lstr "strings"

	ljparser "github.com/cuongpiger/joat/parser"

	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
)

const (
	defaultPackageId = "d33c1f28-cb27-442d-81ac-062f57bd52b9"

	VStorageServiceId = "b9ba2b16-389e-48b7-9e75-4c991239da27"
	VCRServiceId      = "4f823540-4d64-4bf5-a3e2-322b098b601b"
	VServerServiceId  = "e3f1b087-2d8b-4258-8b49-f5c6a4d27411"
	IAMServiceId      = "f3d11a4c-f071-4009-88a6-4a21346c8708"
	VMonitorServiceId = "9cc0d21a-5c27-4295-9787-eea213e255a1"

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
	opts.ResourceInfo.PackageUuid = defaultPackageId

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

func NewListTagsByEndpointIdRequest(pendpointId string) IListTagsByEndpointIdRequest {
	opt := new(ListTagsByEndpointIdRequest)
	opt.Id = pendpointId
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

func (s *CreateEndpointRequest) GetParameters() map[string]interface{} {
	res := map[string]interface{}{
		"resourceType":      s.ResourceType,
		"action":            s.Action,
		"isBuyMorePoc":      s.ResourceInfo.IsBuyMorePoc,
		"isPoc":             s.ResourceInfo.IsPoc,
		"isEnableAutoRenew": s.ResourceInfo.IsEnableAutoRenew,
		"endpointName":      s.ResourceInfo.EndpointName,
		"categoryID":        s.ResourceInfo.CategoryUuid,
		"serviceId":         s.ResourceInfo.ServiceUuid,
		"packageId":         s.ResourceInfo.PackageUuid,
		"vpcId":             s.ResourceInfo.VpcUuid,
		"subnetId":          s.ResourceInfo.SubnetUuid,
		"regionId":          s.ResourceInfo.RegionUuid,
		"projectId":         s.ResourceInfo.ProjectUuid,
		"description":       s.ResourceInfo.Description,
	}

	if s.UserAgent.Agent != nil && len(s.UserAgent.Agent) > 0 {
		res["userAgent"] = s.UserAgent.Agent
	}

	return res
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

func (s *CreateEndpointRequest) WithPoc(pyes bool) ICreateEndpointRequest {
	s.ResourceInfo.IsPoc = pyes
	return s
}

func (s *CreateEndpointRequest) WithBuyMorePoc(pyes bool) ICreateEndpointRequest {
	s.ResourceInfo.IsBuyMorePoc = pyes
	return s
}

func (s *CreateEndpointRequest) WithEnableAutoRenew(pyes bool) ICreateEndpointRequest {
	s.ResourceInfo.IsEnableAutoRenew = pyes
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

func (s *DeleteEndpointByIdRequest) GetParameters() map[string]interface{} {
	res := map[string]interface{}{
		"serviceId":  s.EndpointServiceUuid,
		"endpointId": s.EndpointId,
		"projectId":  s.ProjectUuid,
		"regionId":   s.RegionUuid,
		"vpcId":      s.VpcUuid,
	}

	if s.UserAgent.Agent != nil && len(s.UserAgent.Agent) > 0 {
		res["userAgent"] = s.UserAgent.Agent
	}

	return res
}

type ListEndpointsRequest struct {
	Page  int
	Size  int
	VpcId string
	Uuid  string
	lscommon.UserAgent
}

func (s *ListEndpointsRequest) GetParameters() map[string]interface{} {
	res := map[string]interface{}{
		"page":  s.Page,
		"size":  s.Size,
		"vpcId": s.VpcId,
		"uuid":  s.Uuid,
	}

	if s.UserAgent.Agent != nil && len(s.UserAgent.Agent) > 0 {
		res["userAgent"] = s.UserAgent.Agent
	}

	return res
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

func (s *ListEndpointsRequest) WithUuid(puuid string) IListEndpointsRequest {
	s.Uuid = puuid
	return s
}

func (s *ListEndpointsRequest) ToListQuery() (string, error) {
	var params []string
	if s.VpcId != "" {
		params = append(params, lfmt.Sprintf(`{"field":"vpcId","value":"%s"}`, s.VpcId))
	}

	if s.Uuid != "" {
		params = append(params, lfmt.Sprintf(`{"field":"uuid","value":"%s"}`, s.Uuid))
	}

	paramsFilter := lstr.Join(params, ",")
	query := lfmt.Sprintf(`{"page":%d,"size":%d,"search":[%s]}`, s.Page, s.Size, paramsFilter)
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

// _____________________________________________________________________ ListTagsByEndpointIdRequest

type ListTagsByEndpointIdRequest struct {
	lscommon.UserAgent
	lscommon.EndpointCommon

	Id string `q:"resourceUuid"`
}

func (s *ListTagsByEndpointIdRequest) ToListQuery() (string, error) {
	parser, _ := ljparser.GetParser()
	url, err := parser.UrlMe(s)
	if err != nil {
		return "", err
	}

	return url.String(), err
}

func (s *ListTagsByEndpointIdRequest) GetDefaultQuery() string {
	query := lfmt.Sprintf(`{"page":%d,"size":%d}`, defaultListEndpointsRequestPage, defaultListEndpointsRequestSize)
	query = "params=" + lurl.QueryEscape(query)
	return query
}

func (s *ListTagsByEndpointIdRequest) GetParameters() map[string]interface{} {
	res := map[string]interface{}{
		"resourceUuid": s.Id,
	}

	if s.UserAgent.Agent != nil && len(s.UserAgent.Agent) > 0 {
		res["userAgent"] = s.UserAgent.Agent
	}

	return res
}
