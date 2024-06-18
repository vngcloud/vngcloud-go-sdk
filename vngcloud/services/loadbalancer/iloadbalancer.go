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
	ListPoolsByLoadBalancerId(popts lslbSvcV2.IListPoolsByLoadBalancerIdRequest) (*lsentity.ListPools, lserr.ISdkError)
	UpdatePoolMembers(popts lslbSvcV2.IUpdatePoolMembersRequest) lserr.ISdkError
	ListPoolMembers(popts lslbSvcV2.IListPoolMembersRequest) (*lsentity.ListMembers, lserr.ISdkError)
	DeletePoolById(popt lslbSvcV2.IDeletePoolByIdRequest) lserr.ISdkError
	DeleteListenerById(popts lslbSvcV2.IDeleteListenerByIdRequest) lserr.ISdkError
	DeleteLoadBalancerById(popts lslbSvcV2.IDeleteLoadBalancerByIdRequest) lserr.ISdkError
}

type ILoadBalancerServiceInternal interface {
	CreateLoadBalancer(popts lslbSvcIn.ICreateLoadBalancerRequest) (*lsentity.LoadBalancer, lserr.ISdkError)
}
