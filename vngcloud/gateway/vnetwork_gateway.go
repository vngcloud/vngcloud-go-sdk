package gateway

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsnwSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network"
)

type vnetworkGatewayV1 struct {
	networkService lsnwSvc.INetworkServiceV1
}

func (s *vnetworkGatewayV1) NetworkService() lsnwSvc.INetworkServiceV1 {
	return s.networkService
}

func NewVNetworkGatewayV1(psvcClient lsclient.IServiceClient) IVNetworkGatewayV1 {
	return &vnetworkGatewayV1{
		networkService: lsnwSvc.NewNetworkServiceV1(psvcClient),
	}
}
