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
	AddUserAgent(pagent ...string) ICreateEndpointRequest
}

type IDeleteEndpointByIdRequest interface {
	GetEndpointId() string
	AddUserAgent(pagent ...string) IDeleteEndpointByIdRequest
	ToRequestBody(psvc lsclient.IServiceClient) interface{}
}
