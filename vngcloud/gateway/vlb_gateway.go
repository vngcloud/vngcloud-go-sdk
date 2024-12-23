package gateway

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lslbSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/loadbalancer"
)

type vlbGatewayV2 struct {
	lbService lslbSvc.ILoadBalancerServiceV2
}

type vlbGatewayInternal struct {
	lbService lslbSvc.ILoadBalancerServiceInternal
}

type vlbGatewayGlobal struct {
	lbService lslbSvc.ILoadBalancerServiceGlobal
}

func (s *vlbGatewayGlobal) LoadBalancerService() lslbSvc.ILoadBalancerServiceGlobal {
	return s.lbService
}

func NewVLBGatewayV2(plbSvcClient, pserverSvcClient lsclient.IServiceClient) IVLBGatewayV2 {
	return &vlbGatewayV2{
		lbService: lslbSvc.NewLoadBalancerServiceV2(plbSvcClient, pserverSvcClient),
	}
}

func NewVLBGatewayInternal(psvcClient lsclient.IServiceClient) IVLBGatewayInternal {
	return &vlbGatewayInternal{
		lbService: lslbSvc.NewLoadBalancerServiceInternal(psvcClient),
	}
}

func NewVLBGatewayGlobal(psvcClient lsclient.IServiceClient) IVLBGatewayGlobal {
	return &vlbGatewayGlobal{
		lbService: lslbSvc.NewLoadBalancerServiceGlobal(psvcClient),
	}
}

func (s *vlbGatewayInternal) LoadBalancerService() lslbSvc.ILoadBalancerServiceInternal {
	return s.lbService
}

func (s *vlbGatewayV2) LoadBalancerService() lslbSvc.ILoadBalancerServiceV2 {
	return s.lbService
}
