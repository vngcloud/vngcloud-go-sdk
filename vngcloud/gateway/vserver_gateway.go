package gateway

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lscomputeSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/compute"
	lsnetworkSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network"
	lsportalSvcV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/portal"
)

type vserverGatewayV1 struct {
	portalService lsportalSvcV1.IPortalServiceV1
}

type vserverGatewayV2 struct {
	networkService lsnetworkSvcV2.INetworkServiceV2
	computeService lscomputeSvcV2.IComputeServiceV2
}

func NewVServerGatewayV1(psvcClient lsclient.IServiceClient) IVServerGatewayV1 {
	return &vserverGatewayV1{
		portalService: lsportalSvcV1.NewPortalService(psvcClient),
	}
}

func NewVServerGatewayV2(psvcClient lsclient.IServiceClient) IVServerGatewayV2 {
	return &vserverGatewayV2{
		networkService: lsnetworkSvcV2.NewNetworkService(psvcClient),
		computeService: lscomputeSvcV2.NewComputeService(psvcClient),
	}
}

func (s *vserverGatewayV1) PortalService() lsportalSvcV1.IPortalServiceV1 {
	return s.portalService
}

func (s *vserverGatewayV2) NetworkService() lsnetworkSvcV2.INetworkServiceV2 {
	return s.networkService
}

func (s *vserverGatewayV2) ComputeService() lscomputeSvcV2.IComputeServiceV2 {
	return s.computeService
}
