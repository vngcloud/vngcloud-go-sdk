package gateway

import (
	lsidentitySvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/identity"
	lsportalSvcV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/portal"
)

type IIamGateway interface {
	V2() IIamGatewayV2
}

type IIamGatewayV2 interface {
	IdentityService() lsidentitySvcV2.IIdentityServiceV2
}

type IVServerGateway interface {
	V1() IVServerGatewayV1
}

type IVServerGatewayV1 interface {
	PortalService() lsportalSvcV1.IPortalServiceV1
}

type IVLBGateway interface{}

type IVBackUpGateway interface{}
