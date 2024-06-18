package loadbalancer

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lslbSvcIn "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/loadbalancer/inter"
	lslbSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/loadbalancer/v2"
)

type ILoadBalancerServiceV2 interface {
	CreateLoadBalancer(popts lslbSvcV2.ICreateLoadBalancerRequest) (*lsentity.LoadBalancer, lserr.ISdkError)
	GetLoadBalancerById(popts lslbSvcV2.IGetLoadBalancerByIdRequest) (*lsentity.LoadBalancer, lserr.ISdkError)
	ListLoadBalancers(popts lslbSvcV2.IListLoadBalancersRequest) (*lsentity.ListLoadBalancers, lserr.ISdkError)
	CreatePool(popts lslbSvcV2.ICreatePoolRequest) (*lsentity.Pool, lserr.ISdkError)
	CreateListener(popts lslbSvcV2.ICreateListenerRequest) (*lsentity.Listener, lserr.ISdkError)
	UpdateListener(popts lslbSvcV2.IUpdateListenerRequest) lserr.ISdkError
	ListListenersByLoadBalancerId(popts lslbSvcV2.IListListenersByLoadBalancerIdRequest) (*lsentity.ListListeners, lserr.ISdkError)
}

type ILoadBalancerServiceInternal interface {
	CreateLoadBalancer(popts lslbSvcIn.ICreateLoadBalancerRequest) (*lsentity.LoadBalancer, lserr.ISdkError)
}
