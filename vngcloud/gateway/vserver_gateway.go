package gateway

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lscomputeSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/compute"
	lsnetworkSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network"
	lsportalSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/portal"
)

type vserverGatewayV1 struct {
	portalService lsportalSvc.IPortalServiceV1
}

type vserverGatewayV2 struct {
	networkService lsnetworkSvc.INetworkServiceV2
	computeService lscomputeSvc.IComputeServiceV2
	portalService  lsportalSvc.IPortalServiceV2
}

func NewVServerGatewayV1(psvcClient lsclient.IServiceClient) IVServerGatewayV1 {
	return &vserverGatewayV1{
		portalService: lsportalSvc.NewPortalServiceV1(psvcClient),
	}
}

func NewVServerGatewayV2(psvcClient lsclient.IServiceClient) IVServerGatewayV2 {
	return &vserverGatewayV2{
		networkService: lsnetworkSvc.NewNetworkServiceV2(psvcClient),
		computeService: lscomputeSvc.NewComputeServiceV2(psvcClient),
		portalService:  lsportalSvc.NewPortalServiceV2(psvcClient),
	}
}

func (s *vserverGatewayV1) PortalService() lsportalSvc.IPortalServiceV1 {
	return s.portalService
}

func (s *vserverGatewayV2) NetworkService() lsnetworkSvc.INetworkServiceV2 {
	return s.networkService
}

func (s *vserverGatewayV2) ComputeService() lscomputeSvc.IComputeServiceV2 {
	return s.computeService
}

func (s *vserverGatewayV2) PortalService() lsportalSvc.IPortalServiceV2 {
	return s.portalService
}
