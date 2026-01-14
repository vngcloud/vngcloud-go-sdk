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
	defaultPackageId = "d391c404-51b0-4f4d-a946-ec35c8a6e2be"

	VStorageServiceId = "b9ba2b16-389e-48b7-9e75-4c991239da27"
	VCRServiceId      = "4f823540-4d64-4bf5-a3e2-322b098b601b"
	VServerServiceId  = "e3f1b087-2d8b-4258-8b49-f5c6a4d27411"
	IAMServiceId      = "f3d11a4c-f071-4009-88a6-4a21346c8708"
	VMonitorServiceId = "9cc0d21a-5c27-4295-9787-eea213e255a1"

	defaultListEndpointsRequestPage = 1
	defaultListEndpointsRequestSize = 10
)

type GetEndpointByIdRequest struct {
	lscommon.UserAgent
	lscommon.EndpointCommon
}

func (s *GetEndpointByIdRequest) AddUserAgent(pagent ...string) IGetEndpointByIdRequest {
	s.Agent = append(s.Agent, pagent...)
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
		PortalUserId      string `json:"portalUserId"`
		SubnetUuid        string `json:"subnetUuid"`
		RegionUuid        string `json:"regionUuid"`
		ProjectUuid       string `json:"projectUuid"`
		Description       string `json:"description"`
		EnableAZ          bool   `json:"enableAZ"`
		EnableDnsName     bool   `json:"enableDnsName"`
		Networking        []struct {
			Zone       string `json:"zone"`
			SubnetUuid string `json:"subnetUuid"`
		} `json:"networking"`
		Scaling struct {
			MinSize int `json:"minSize"`
			MaxSize int `json:"maxSize"`
		} `json:"scaling"`
	} `json:"resourceInfo"`

	lscommon.UserAgent
}

func (s *CreateEndpointRequest) ToMap() map[string]interface{} {
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
		"enableAZ":          s.ResourceInfo.EnableAZ,
		"enableDnsName":     s.ResourceInfo.EnableDnsName,
		"networking":        s.ResourceInfo.Networking,
		"scaling":           s.ResourceInfo.Scaling,
	}

	if len(s.Agent) > 0 {
		res["userAgent"] = s.Agent
	}

	return res
}

func (s *CreateEndpointRequest) AddUserAgent(pagent ...string) ICreateEndpointRequest {
	s.Agent = append(s.Agent, pagent...)
	return s
}

func (s *CreateEndpointRequest) ToRequestBody(psvc lsclient.IServiceClient) interface{} {
	s.ResourceType = "endpoint"
	s.Action = "create"
	s.ResourceInfo.EnableAZ = true
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

func (s *CreateEndpointRequest) WithPortalUserId(portalUserId string) ICreateEndpointRequest {
	s.ResourceInfo.PortalUserId = portalUserId
	return s
}

func (s *CreateEndpointRequest) GetPortalUserId() string {
	return s.ResourceInfo.PortalUserId
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

func (s *CreateEndpointRequest) WithEnableDnsName(pyes bool) ICreateEndpointRequest {
	s.ResourceInfo.EnableDnsName = pyes
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

func (s *CreateEndpointRequest) AddNetworking(zone, subnetUuid string) ICreateEndpointRequest {
	s.ResourceInfo.Networking = append(s.ResourceInfo.Networking, struct {
		Zone       string `json:"zone"`
		SubnetUuid string `json:"subnetUuid"`
	}{
		Zone:       zone,
		SubnetUuid: subnetUuid,
	})
	return s
}

func (s *CreateEndpointRequest) WithScaling(minSize int, maxSize int) ICreateEndpointRequest {
	s.ResourceInfo.Scaling = struct {
		MinSize int `json:"minSize"`
		MaxSize int `json:"maxSize"`
	}{
		MinSize: minSize,
		MaxSize: maxSize,
	}
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
	s.Agent = append(s.Agent, pagent...)
	return s
}

func (s *DeleteEndpointByIdRequest) ToRequestBody(psvc lsclient.IServiceClient) interface{} {
	s.ProjectUuid = psvc.GetProjectId()
	s.RegionUuid = psvc.GetZoneId()

	return s
}

func (s *DeleteEndpointByIdRequest) ToMap() map[string]interface{} {
	res := map[string]interface{}{
		"serviceId":  s.EndpointServiceUuid,
		"endpointId": s.EndpointId,
		"projectId":  s.ProjectUuid,
		"regionId":   s.RegionUuid,
		"vpcId":      s.VpcUuid,
	}

	if len(s.Agent) > 0 {
		res["userAgent"] = s.Agent
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

func (s *ListEndpointsRequest) ToMap() map[string]interface{} {
	res := map[string]interface{}{
		"page":  s.Page,
		"size":  s.Size,
		"vpcId": s.VpcId,
		"uuid":  s.Uuid,
	}

	if len(s.Agent) > 0 {
		res["userAgent"] = s.Agent
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
	s.Agent = append(s.Agent, pagent...)
	return s
}

// _____________________________________________________________________ ListTagsByEndpointIdRequest

type ListTagsByEndpointIdRequest struct {
	lscommon.UserAgent
	lscommon.EndpointCommon
	lscommon.PortalUser

	ProjectId string
	Id        string `q:"resourceUuid"`
}

func (s *ListTagsByEndpointIdRequest) ToListQuery() (string, error) {
	parser, _ := ljparser.GetParser()
	url, err := parser.UrlMe(s)
	if err != nil {
		return "", err
	}

	return url.String(), err
}

func (s *ListTagsByEndpointIdRequest) GetProjectId() string {
	return s.ProjectId
}

func (s *ListTagsByEndpointIdRequest) GetDefaultQuery() string {
	query := lfmt.Sprintf(`{"page":%d,"size":%d}`, defaultListEndpointsRequestPage, defaultListEndpointsRequestSize)
	query = "params=" + lurl.QueryEscape(query)
	return query
}

func (s *ListTagsByEndpointIdRequest) ToMap() map[string]interface{} {
	res := map[string]interface{}{
		"resourceUuid": s.Id,
	}

	if len(s.Agent) > 0 {
		res["userAgent"] = s.Agent
	}

	return res
}

func (s *ListTagsByEndpointIdRequest) GetMapHeaders() map[string]string {
	return s.PortalUser.GetMapHeaders()
}

func (s *ListTagsByEndpointIdRequest) AddUserAgent(pagent ...string) IListTagsByEndpointIdRequest {
	s.Agent = append(s.Agent, pagent...)
	return s
}

// _________________________________________________________________ CreateTagsWithEndpointIdRequest

type CreateTagsWithEndpointIdRequest struct {
	lscommon.UserAgent
	lscommon.EndpointCommon
	lscommon.PortalUser

	ProjectId    string
	ResourceUuid string `json:"resourceUuid"`
	Tags         []struct {
		TagKey   string `json:"tagKey"`
		TagValue string `json:"tagValue"`
	} `json:"tags"`

	SystemTag bool `json:"systemTag"`
}

func (s *CreateTagsWithEndpointIdRequest) ToMap() map[string]interface{} {
	res := map[string]interface{}{
		"resourceUuid": s.Id,
	}

	if len(s.Agent) > 0 {
		res["userAgent"] = s.Agent
	}

	res["tags"] = s.Tags

	return res
}

func (s *CreateTagsWithEndpointIdRequest) AddUserAgent(pagent ...string) ICreateTagsWithEndpointIdRequest {
	s.Agent = append(s.Agent, pagent...)
	return s
}

func (s *CreateTagsWithEndpointIdRequest) GetMapHeaders() map[string]string {
	return s.PortalUser.GetMapHeaders()
}

func (s *CreateTagsWithEndpointIdRequest) AddTag(pkey, pvalue string) ICreateTagsWithEndpointIdRequest {
	s.Tags = append(s.Tags, struct {
		TagKey   string `json:"tagKey"`
		TagValue string `json:"tagValue"`
	}{
		TagKey:   pkey,
		TagValue: pvalue,
	})

	return s
}

func (s *CreateTagsWithEndpointIdRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateTagsWithEndpointIdRequest) GetProjectId() string {
	return s.ProjectId
}

// ____________________________________________________________________ DeleteTagByEndpointIdRequest

type DeleteTagOfEndpointRequest struct {
	lscommon.UserAgent
	lscommon.PortalUser

	ProjectId string
	TagId     string
}

func (s *DeleteTagOfEndpointRequest) ToMap() map[string]interface{} {
	res := map[string]interface{}{
		"tagId": s.TagId,
	}

	if len(s.Agent) > 0 {
		res["userAgent"] = s.Agent
	}

	return res
}

func (s *DeleteTagOfEndpointRequest) AddUserAgent(pagent ...string) IDeleteTagOfEndpointRequest {
	s.Agent = append(s.Agent, pagent...)
	return s
}

func (s *DeleteTagOfEndpointRequest) GetMapHeaders() map[string]string {
	return s.PortalUser.GetMapHeaders()
}

func (s *DeleteTagOfEndpointRequest) GetTagId() string {
	return s.TagId
}

func (s *DeleteTagOfEndpointRequest) GetProjectId() string {
	return s.ProjectId
}

// _________________________________________________________________ UpdateTagValueOfEndpointRequest

type UpdateTagValueOfEndpointRequest struct {
	lscommon.UserAgent
	lscommon.PortalUser

	TagId     string
	ProjectId string
	TagValue  string `json:"tagValue"`
}

func (s *UpdateTagValueOfEndpointRequest) ToMap() map[string]interface{} {
	res := map[string]interface{}{
		"tagId":    s.TagId,
		"tagValue": s.TagValue,
	}

	if len(s.Agent) > 0 {
		res["userAgent"] = s.Agent
	}

	return res
}

func (s *UpdateTagValueOfEndpointRequest) AddUserAgent(pagent ...string) IUpdateTagValueOfEndpointRequest {
	s.Agent = append(s.Agent, pagent...)
	return s
}

func (s *UpdateTagValueOfEndpointRequest) GetMapHeaders() map[string]string {
	return s.PortalUser.GetMapHeaders()
}

func (s *UpdateTagValueOfEndpointRequest) GetTagId() string {
	return s.TagId
}

func (s *UpdateTagValueOfEndpointRequest) ToRequestBody() interface{} {
	return s
}

func (s *UpdateTagValueOfEndpointRequest) GetProjectId() string {
	return s.ProjectId
}
