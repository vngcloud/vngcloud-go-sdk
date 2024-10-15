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
	WithPackageUuid(ppackageUuid string) ICreateEndpointRequest
	WithServiceUuid(pserviceUuid string) ICreateEndpointRequest
	WithCategoryUuid(pcategoryUuid string) ICreateEndpointRequest
	WithEndpointName(pendpointName string) ICreateEndpointRequest
	WithPoc(pyes bool) ICreateEndpointRequest
	WithBuyMorePoc(pyes bool) ICreateEndpointRequest
	WithEnableAutoRenew(pyes bool) ICreateEndpointRequest
	AddUserAgent(pagent ...string) ICreateEndpointRequest
	GetParameters() map[string]interface{}
}

type IDeleteEndpointByIdRequest interface {
	GetEndpointId() string
	AddUserAgent(pagent ...string) IDeleteEndpointByIdRequest
	ToRequestBody(psvc lsclient.IServiceClient) interface{}
	GetParameters() map[string]interface{}
}

type IListEndpointsRequest interface {
	WithPage(ppage int) IListEndpointsRequest
	WithSize(psize int) IListEndpointsRequest
	WithVpcId(pvpcId string) IListEndpointsRequest
	WithUuid(puuid string) IListEndpointsRequest
	ToListQuery() (string, error)
	GetDefaultQuery() string
	AddUserAgent(pagent ...string) IListEndpointsRequest
	GetParameters() map[string]interface{}
}

type IListTagsByEndpointIdRequest interface {
	ToListQuery() (string, error)
	GetDefaultQuery() string
	GetParameters() map[string]interface{}
	GetMapHeaders() map[string]string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IListTagsByEndpointIdRequest
}

type ICreateTagsWithEndpointIdRequest interface {
	GetParameters() map[string]interface{}
	AddUserAgent(pagent ...string) ICreateTagsWithEndpointIdRequest
	GetMapHeaders() map[string]string
	AddTag(pkey, pvalue string) ICreateTagsWithEndpointIdRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
}
