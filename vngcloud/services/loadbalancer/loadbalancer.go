package loadbalancer

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsloadbalancerInternal "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/loadbalancer/inter"
	lsloadbalancerV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/loadbalancer/v2"
)

func NewLoadBalancerServiceV2(plbSvcClient, pserverSvcClient lsclient.IServiceClient) ILoadBalancerServiceV2 {
	return &lsloadbalancerV2.LoadBalancerServiceV2{
		VLBClient:     plbSvcClient,
		VServerClient: pserverSvcClient,
	}
}

func NewLoadBalancerServiceInternal(psvcClient lsclient.IServiceClient) ILoadBalancerServiceInternal {
	return &lsloadbalancerInternal.LoadBalancerServiceInternal{
		VLBClient: psvcClient,
	}
}
