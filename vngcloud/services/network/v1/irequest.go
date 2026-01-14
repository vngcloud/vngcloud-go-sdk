package v1

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

type IGetEndpointByIdRequest interface {
	GetEndpointId() string
	AddUserAgent(pagent ...string) IGetEndpointByIdRequest
}

type ICreateEndpointRequest interface {
	ToRequestBody(psvc lsclient.IServiceClient) interface{}
	WithDescription(pdesp string) ICreateEndpointRequest
	WithSubnetUuid(psubnetUuid string) ICreateEndpointRequest
	WithVpcUuid(pvpcUuid string) ICreateEndpointRequest
	GetPortalUserId() string
	WithPortalUserId(portalUserId string) ICreateEndpointRequest
	WithPackageUuid(ppackageUuid string) ICreateEndpointRequest
	WithServiceUuid(pserviceUuid string) ICreateEndpointRequest
	WithCategoryUuid(pcategoryUuid string) ICreateEndpointRequest
	WithEndpointName(pendpointName string) ICreateEndpointRequest
	WithPoc(pyes bool) ICreateEndpointRequest
	WithEnableDnsName(pyes bool) ICreateEndpointRequest
	WithBuyMorePoc(pyes bool) ICreateEndpointRequest
	WithEnableAutoRenew(pyes bool) ICreateEndpointRequest
	AddNetworking(zone string, subnetUuid string) ICreateEndpointRequest
	WithScaling(minSize int, maxSize int) ICreateEndpointRequest
	AddUserAgent(pagent ...string) ICreateEndpointRequest
	ToMap() map[string]interface{}
}

type IDeleteEndpointByIdRequest interface {
	GetEndpointId() string
	AddUserAgent(pagent ...string) IDeleteEndpointByIdRequest
	ParseUserAgent() string
	ToRequestBody(psvc lsclient.IServiceClient) interface{}
	ToMap() map[string]interface{}
}

type IListEndpointsRequest interface {
	WithPage(ppage int) IListEndpointsRequest
	WithSize(psize int) IListEndpointsRequest
	WithVpcId(pvpcId string) IListEndpointsRequest
	WithUuid(puuid string) IListEndpointsRequest
	ToListQuery() (string, error)
	GetDefaultQuery() string
	AddUserAgent(pagent ...string) IListEndpointsRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IListTagsByEndpointIdRequest interface {
	ToListQuery() (string, error)
	GetDefaultQuery() string
	ToMap() map[string]interface{}
	GetMapHeaders() map[string]string
	ParseUserAgent() string
	GetProjectId() string
	AddUserAgent(pagent ...string) IListTagsByEndpointIdRequest
}

type ICreateTagsWithEndpointIdRequest interface {
	ToMap() map[string]interface{}
	AddUserAgent(pagent ...string) ICreateTagsWithEndpointIdRequest
	GetMapHeaders() map[string]string
	AddTag(pkey, pvalue string) ICreateTagsWithEndpointIdRequest
	ParseUserAgent() string
	GetProjectId() string
	ToRequestBody() interface{}
}

type IDeleteTagOfEndpointRequest interface {
	ToMap() map[string]interface{}
	AddUserAgent(pagent ...string) IDeleteTagOfEndpointRequest
	GetMapHeaders() map[string]string
	ParseUserAgent() string
	GetTagId() string
	GetProjectId() string
}

type IUpdateTagValueOfEndpointRequest interface {
	ToMap() map[string]interface{}
	AddUserAgent(pagent ...string) IUpdateTagValueOfEndpointRequest
	GetMapHeaders() map[string]string
	ParseUserAgent() string
	GetTagId() string
	GetProjectId() string
	ToRequestBody() interface{}
}
