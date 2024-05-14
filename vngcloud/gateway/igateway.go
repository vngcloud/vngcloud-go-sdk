package gateway

import lsidentitySvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/identity"

type IIamGateway interface {
	V2() IIamGatewayV2
}

type IIamGatewayV2 interface {
	IdentityService() lsidentitySvc.IIdentityServiceV2
}

type IVServerGateway interface {
}

type IVLBGateway interface{}
