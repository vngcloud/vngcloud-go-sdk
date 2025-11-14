package gateway

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lscomputeSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/compute"
	lsnetworkSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network"
	lsportalSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/portal"
	lsServerSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/server"
	lsvolumeSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/volume"
)

type vserverGatewayV1 struct {
	portalService lsportalSvc.IPortalServiceV1
	volumeService lsvolumeSvc.IVolumeServiceV1
}

type vServerGatewayInternalV1 struct {
	serverService lsServerSvc.IServerServiceInternalV1
}

type vserverGatewayV2 struct {
	networkService lsnetworkSvc.INetworkServiceV2
	computeService lscomputeSvc.IComputeServiceV2
	portalService  lsportalSvc.IPortalServiceV2
	volumeService  lsvolumeSvc.IVolumeServiceV2
}

func NewVServerGatewayV1(psvcClient lsclient.IServiceClient) IVServerGatewayV1 {
	return &vserverGatewayV1{
		portalService: lsportalSvc.NewPortalServiceV1(psvcClient),
		volumeService: lsvolumeSvc.NewVolumeServiceV1(psvcClient),
	}
}

func NewVServerGatewayV2(psvcClient lsclient.IServiceClient) IVServerGatewayV2 {
	return &vserverGatewayV2{
		networkService: lsnetworkSvc.NewNetworkServiceV2(psvcClient),
		computeService: lscomputeSvc.NewComputeServiceV2(psvcClient),
		portalService:  lsportalSvc.NewPortalServiceV2(psvcClient),
		volumeService:  lsvolumeSvc.NewVolumeServiceV2(psvcClient),
	}
}

func NewVServerGatewayInternalV1(psvcClient lsclient.IServiceClient) IVServerGatewayInternalV1 {
	return &vServerGatewayInternalV1{
		serverService: lsServerSvc.NewServerServiceInternalV1(psvcClient),
	}
}

func (s *vserverGatewayV1) PortalService() lsportalSvc.IPortalServiceV1 {
	return s.portalService
}

func (s *vserverGatewayV1) VolumeService() lsvolumeSvc.IVolumeServiceV1 {
	return s.volumeService
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

func (s *vserverGatewayV2) VolumeService() lsvolumeSvc.IVolumeServiceV2 {
	return s.volumeService
}

func (s *vServerGatewayInternalV1) ServerService() lsServerSvc.IServerServiceInternalV1 {
	return s.serverService
}
