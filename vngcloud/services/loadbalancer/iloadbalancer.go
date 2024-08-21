package loadbalancer

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lslbSvcIn "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/loadbalancer/inter"
	lslbSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/loadbalancer/v2"
)

type ILoadBalancerServiceV2 interface {
	CreateLoadBalancer(popts lslbSvcV2.ICreateLoadBalancerRequest) (*lsentity.LoadBalancer, lserr.IError)
	GetLoadBalancerById(popts lslbSvcV2.IGetLoadBalancerByIdRequest) (*lsentity.LoadBalancer, lserr.IError)
	ListLoadBalancers(popts lslbSvcV2.IListLoadBalancersRequest) (*lsentity.ListLoadBalancers, lserr.IError)
	CreatePool(popts lslbSvcV2.ICreatePoolRequest) (*lsentity.Pool, lserr.IError)
	CreateListener(popts lslbSvcV2.ICreateListenerRequest) (*lsentity.Listener, lserr.IError)
	UpdateListener(popts lslbSvcV2.IUpdateListenerRequest) lserr.IError
	ListListenersByLoadBalancerId(popts lslbSvcV2.IListListenersByLoadBalancerIdRequest) (*lsentity.ListListeners, lserr.IError)
	ListPoolsByLoadBalancerId(popts lslbSvcV2.IListPoolsByLoadBalancerIdRequest) (*lsentity.ListPools, lserr.IError)
	UpdatePoolMembers(popts lslbSvcV2.IUpdatePoolMembersRequest) lserr.IError
	ListPoolMembers(popts lslbSvcV2.IListPoolMembersRequest) (*lsentity.ListMembers, lserr.IError)
	DeletePoolById(popt lslbSvcV2.IDeletePoolByIdRequest) lserr.IError
	DeleteListenerById(popts lslbSvcV2.IDeleteListenerByIdRequest) lserr.IError
	DeleteLoadBalancerById(popts lslbSvcV2.IDeleteLoadBalancerByIdRequest) lserr.IError
	ListTags(popts lslbSvcV2.IListTagsRequest) (*lsentity.ListTags, lserr.IError)
	CreateTags(popts lslbSvcV2.ICreateTagsRequest) lserr.IError
	UpdateTags(popts lslbSvcV2.IUpdateTagsRequest) lserr.IError
}

type ILoadBalancerServiceInternal interface {
	CreateLoadBalancer(popts lslbSvcIn.ICreateLoadBalancerRequest) (*lsentity.LoadBalancer, lserr.IError)
}
