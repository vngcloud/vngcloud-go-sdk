package gateway

import (
	lscomputeSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/compute"
	lsidentitySvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/identity"
	lsnetworkSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network"
	lsportalSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/portal"
	lsvolumeSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/volume"
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

	// GetEndpoint returns the endpoint of the vServer service
	GetEndpoint() string
}

type IVServerGatewayV1 interface {
	PortalService() lsportalSvc.IPortalServiceV1
	VolumeService() lsvolumeSvc.IVolumeServiceV1
}

type IVServerGatewayV2 interface {
	NetworkService() lsnetworkSvc.INetworkServiceV2
	ComputeService() lscomputeSvc.IComputeServiceV2
	PortalService() lsportalSvc.IPortalServiceV2
	VolumeService() lsvolumeSvc.IVolumeServiceV2
}

type IVLBGateway interface{}

type IVBackUpGateway interface{}
