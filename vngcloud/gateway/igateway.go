package gateway

import (
	lscomputeSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/compute"
	lsidentitySvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/identity"
	lsnetworkSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network"
	lsportalSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/portal"
)

type IIamGateway interface {
	V2() IIamGatewayV2
}

type IIamGatewayV2 interface {
	IdentityService() lsidentitySvc.IIdentityServiceV2
}

type IVServerGateway interface {
	V1() IVServerGatewayV1
	V2() IVServerGatewayV2
}

type IVServerGatewayV1 interface {
	PortalService() lsportalSvc.IPortalServiceV1
}

type IVServerGatewayV2 interface {
	NetworkService() lsnetworkSvc.INetworkServiceV2
	ComputeService() lscomputeSvc.IComputeServiceV2
	PortalService() lsportalSvc.IPortalServiceV2
}

type IVLBGateway interface{}

type IVBackUpGateway interface{}
