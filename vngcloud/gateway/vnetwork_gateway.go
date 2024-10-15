package gateway

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsnwSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network"
)

type vnetworkGatewayV1 struct {
	networkService lsnwSvc.INetworkServiceV1
}

type vnetworkGatewayInternalV1 struct {
	networkService lsnwSvc.INetworkServiceInternalV1
}

func (s *vnetworkGatewayV1) NetworkService() lsnwSvc.INetworkServiceV1 {
	return s.networkService
}

func (s *vnetworkGatewayInternalV1) NetworkService() lsnwSvc.INetworkServiceInternalV1 {
	return s.networkService
}

func NewVNetworkGatewayV1(psvcClient lsclient.IServiceClient) IVNetworkGatewayV1 {
	return &vnetworkGatewayV1{
		networkService: lsnwSvc.NewNetworkServiceV1(psvcClient),
	}
}

func NewVNetworkGatewayInternalV1(psvcClient lsclient.IServiceClient) IVNetworkGatewayInternalV1 {
	return &vnetworkGatewayInternalV1{
		networkService: lsnwSvc.NewNetworkServiceInternalV1(psvcClient),
	}
}
