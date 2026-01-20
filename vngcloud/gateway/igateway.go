package gateway

import (
	lscomputeSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/compute"
	lsdnsSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/dns"
	"github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/glb"
	lsidentitySvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/identity"
	lslbSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/loadbalancer"
	lsnetworkSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network"
	lsportalSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/portal"
	lsServerSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/server"
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
	InternalV1() IVServerGatewayInternalV1
	V2() IVServerGatewayV2
	// GetEndpoint returns the endpoint of the vServer service
	GetEndpoint() string
}

type IVNetworkGateway interface {
	V1() IVNetworkGatewayV1
	V2() IVNetworkGatewayV1
	InternalV1() IVNetworkGatewayInternalV1
	GetEndpoint() string
}

type IVServerGatewayV1 interface {
	PortalService() lsportalSvc.IPortalServiceV1
	VolumeService() lsvolumeSvc.IVolumeServiceV1
}

type IVServerGatewayInternalV1 interface {
	ServerService() lsServerSvc.IServerServiceInternalV1
}

type IVServerGatewayV2 interface {
	NetworkService() lsnetworkSvc.INetworkServiceV2
	ComputeService() lscomputeSvc.IComputeServiceV2
	PortalService() lsportalSvc.IPortalServiceV2
	VolumeService() lsvolumeSvc.IVolumeServiceV2
}

type IVLBGatewayV2 interface {
	LoadBalancerService() lslbSvc.ILoadBalancerServiceV2
}

type IVNetworkGatewayV1 interface {
	NetworkService() lsnetworkSvc.INetworkServiceV1
}

type IVNetworkGatewayV2 interface {
	NetworkService() lsnetworkSvc.INetworkServiceV2
}

type IVNetworkGatewayInternalV1 interface {
	NetworkService() lsnetworkSvc.INetworkServiceInternalV1
}

type IVLBGatewayInternal interface {
	LoadBalancerService() lslbSvc.ILoadBalancerServiceInternal
}

type IVLBGateway interface {
	Internal() IVLBGatewayInternal
	V2() IVLBGatewayV2
	GetEndpoint() string
}

type IVBackUpGateway interface{}

type IGLBGateway interface {
	V1() IGLBGatewayV1
}

type IGLBGatewayV1 interface {
	GLBService() glb.IGLBServiceV1
}

type IVDnsGateway interface {
	V1() IVDnsGatewayV1
	Internal() IVDnsGatewayInternal
	GetEndpoint() string
}

type IVDnsGatewayV1 interface {
	DnsService() lsdnsSvc.IVDnsServiceV1
}

type IVDnsGatewayInternal interface {
	DnsService() lsdnsSvc.IVDnsServiceInternal
}
