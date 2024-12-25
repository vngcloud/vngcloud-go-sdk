package loadbalancer

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	"github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/loadbalancer/global"
	lslbSvcIn "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/loadbalancer/inter"
	lslbSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/loadbalancer/v2"
)

type ILoadBalancerServiceV2 interface {
	CreateLoadBalancer(popts lslbSvcV2.ICreateLoadBalancerRequest) (*lsentity.LoadBalancer, lserr.IError)
	ResizeLoadBalancer(popts lslbSvcV2.IResizeLoadBalancerRequest) (*lsentity.LoadBalancer, lserr.IError)
	ListLoadBalancerPackages(popts lslbSvcV2.IListLoadBalancerPackagesRequest) (*lsentity.ListLoadBalancerPackages, lserr.IError)
	GetLoadBalancerById(popts lslbSvcV2.IGetLoadBalancerByIdRequest) (*lsentity.LoadBalancer, lserr.IError)
	ListLoadBalancers(popts lslbSvcV2.IListLoadBalancersRequest) (*lsentity.ListLoadBalancers, lserr.IError)
	GetPoolHealthMonitorById(popts lslbSvcV2.IGetPoolHealthMonitorByIdRequest) (*lsentity.HealthMonitor, lserr.IError)
	CreatePool(popts lslbSvcV2.ICreatePoolRequest) (*lsentity.Pool, lserr.IError)
	UpdatePool(popts lslbSvcV2.IUpdatePoolRequest) lserr.IError
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

	ListPolicies(popts lslbSvcV2.IListPoliciesRequest) (*lsentity.ListPolicies, lserr.IError)
	CreatePolicy(popts lslbSvcV2.ICreatePolicyRequest) (*lsentity.Policy, lserr.IError)
	GetPolicyById(popts lslbSvcV2.IGetPolicyByIdRequest) (*lsentity.Policy, lserr.IError)
	UpdatePolicy(popts lslbSvcV2.IUpdatePolicyRequest) lserr.IError
	DeletePolicyById(popts lslbSvcV2.IDeletePolicyByIdRequest) lserr.IError
	GetPoolById(popts lslbSvcV2.IGetPoolByIdRequest) (*lsentity.Pool, lserr.IError)
	GetListenerById(popts lslbSvcV2.IGetListenerByIdRequest) (*lsentity.Listener, lserr.IError)
	ResizeLoadBalancerById(popts lslbSvcV2.IResizeLoadBalancerByIdRequest) lserr.IError
}

type ILoadBalancerServiceInternal interface {
	CreateLoadBalancer(popts lslbSvcIn.ICreateLoadBalancerRequest) (*lsentity.LoadBalancer, lserr.IError)
}

type ILoadBalancerServiceGlobal interface {
	ListGlobalLoadBalancers(popts global.IListGlobalLoadBalancersRequest) (*lsentity.ListGlobalLoadBalancers, lserr.IError)

	ListGlobalPools(popts global.IListGlobalPoolsRequest) (*lsentity.ListGlobalPools, lserr.IError)
	CreateGlobalPool(popts global.ICreateGlobalPoolRequest) (*lsentity.GlobalPool, lserr.IError)
	UpdateGlobalPool(popts global.IUpdateGlobalPoolRequest) (*lsentity.GlobalPool, lserr.IError)
	DeleteGlobalPool(popts global.IDeleteGlobalPoolRequest) lserr.IError

	ListGlobalPoolMembers(popts global.IListGlobalPoolMembersRequest) (*lsentity.ListGlobalPoolMembers, lserr.IError)
	PatchGlobalPoolMember(popts global.IPatchGlobalPoolMemberRequest) lserr.IError

	ListGlobalListeners(popts global.IListGlobalListenersRequest) (*lsentity.ListGlobalListeners, lserr.IError)
	CreateGlobalListener(popts global.ICreateGlobalListenerRequest) (*lsentity.GlobalListener, lserr.IError)
	UpdateGlobalListener(popts global.IUpdateGlobalListenerRequest) (*lsentity.GlobalListener, lserr.IError)
	DeleteGlobalListener(popts global.IDeleteGlobalListenerRequest) lserr.IError
	// GetGlobalListenerById(popts global.IGetGlobalListenerByIdRequest) (*lsentity.GlobalListener, lserr.IError)

	// GetGlobalLoadBalancerPoolById(popts global.IGetGlobalLoadBalancerPoolByIdRequest) (*lsentity.GlobalLoadBalancerPool, lserr.IError)
	// GetGlobalLoadBalancerPoolMemberById(popts global.IGetGlobalLoadBalancerPoolMemberByIdRequest) (*lsentity.GlobalLoadBalancerPoolMember, lserr.IError)
	// UpdateGlobalLoadBalancerPoolMember(popts global.IUpdateGlobalLoadBalancerPoolMemberRequest) lserr.IError
	// DeleteGlobalLoadBalancerPoolMemberById(popts global.IDeleteGlobalLoadBalancerPoolMemberByIdRequest) lserr.IError
	// CreateGlobalLoadBalancer(popts global.ICreateGlobalLoadBalancerRequest) (*lsentity.GlobalLoadBalancer, lserr.IError)
	// UpdateGlobalLoadBalancer(popts global.IUpdateGlobalLoadBalancerRequest) lserr.IError
	// GetGlobalLoadBalancerById(popts global.IGetGlobalLoadBalancerByIdRequest) (*lsentity.GlobalLoadBalancer, lserr.IError)
	// DeleteGlobalLoadBalancerById(popts global.IDeleteGlobalLoadBalancerByIdRequest) lserr.IError
	// ListGlobalLoadBalancerHealthMonitors(popts global.IListGlobalLoadBalancerHealthMonitorsRequest) (*lsentity.ListGlobalLoadBalancerHealthMonitors, lserr.IError)
	// GetGlobalLoadBalancerHealthMonitorById(popts global.IGetGlobalLoadBalancerHealthMonitorByIdRequest) (*lsentity.GlobalLoadBalancerHealthMonitor, lserr.IError)
	// CreateGlobalLoadBalancerHealthMonitor(popts global.ICreateGlobalLoadBalancerHealthMonitorRequest) (*lsentity.GlobalLoadBalancerHealthMonitor, lserr.IError)
	// UpdateGlobalLoadBalancerHealthMonitor(popts global.IUpdateGlobalLoadBalancerHealthMonitorRequest) lserr.IError
	// DeleteGlobalLoadBalancerHealthMonitorById(popts global.IDeleteGlobalLoadBalancerHealthMonitorByIdRequest) lserr.IError
	// ListGlobalLoadBalancerTags(popts global.IListGlobalLoadBalancerTagsRequest) (*lsentity.ListGlobalLoadBalancerTags, lserr.IError)
	// CreateGlobalLoadBalancerTags(popts global.ICreateGlobalLoadBalancerTagsRequest) lserr.IError
	// UpdateGlobalLoadBalancerTags(popts global.IUpdateGlobalLoadBalancerTagsRequest) lserr.IError
	// ListGlobalLoadBalancerPolicies(popts global.IListGlobalLoadBalancerPoliciesRequest) (*lsentity.ListGlobalLoadBalancerPolicies, lserr.IError)
	// GetGlobalLoadBalancerPolicyById(popts global.IGetGlobalLoadBalancerPolicyByIdRequest) (*lsentity.GlobalLoadBalancerPolicy, lserr.IError)
	// CreateGlobalLoadBalancerPolicy(popts global.ICreateGlobalLoadBalancerPolicyRequest) (*lsentity.GlobalLoadBalancerPolicy, lserr.IError)
	// UpdateGlobalLoadBalancerPolicy(popts global.IUpdateGlobalLoadBalancerPolicyRequest) lserr.IError
	// DeleteGlobalLoadBalancerPolicyById(popts global.IDeleteGlobalLoadBalancerPolicyByIdRequest) lserr.IError
	// GetGlobalLoadBalancerPoolById(popts global.IGetGlobalLoadBalancerPoolByIdRequest) (*lsentity.GlobalLoadBalancerPool, lserr.IError)
	// GetGlobalLoadBalancerListenerById(popts global.IGetGlobalLoadBalancerListenerByIdRequest) (*lsentity.GlobalLoadBalancerListener, lserr.IError)
	// GetGlobalLoadBalancerHealthMonitorById(popts global.IGetGlobalLoadBalancerHealthMonitorByIdRequest) (*lsentity.GlobalLoadBalancerHealthMonitor, lserr.IError)
	// GetGlobalLoadBalancerPolicyById(popts global.IGetGlobalLoadBalancerPolicyByIdRequest) (*lsentity.GlobalLoadBalancerPolicy, lserr.IError)
	// GetGlobalLoadBalancerPoolMemberById(popts global.IGetGlobalLoadBalancerPoolMemberByIdRequest) (*lsentity.GlobalLoadBalancerPoolMember, lserr.IError)
}
