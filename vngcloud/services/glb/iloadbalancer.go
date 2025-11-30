package glb

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	v1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/glb/v1"
)

type IGLBServiceV1 interface {
	ListGlobalPools(popts v1.IListGlobalPoolsRequest) (*lsentity.ListGlobalPools, lserr.IError)
	CreateGlobalPool(popts v1.ICreateGlobalPoolRequest) (*lsentity.GlobalPool, lserr.IError)
	UpdateGlobalPool(popts v1.IUpdateGlobalPoolRequest) (*lsentity.GlobalPool, lserr.IError)
	DeleteGlobalPool(popts v1.IDeleteGlobalPoolRequest) lserr.IError

	ListGlobalPoolMembers(popts v1.IListGlobalPoolMembersRequest) (*lsentity.ListGlobalPoolMembers, lserr.IError)
	GetGlobalPoolMember(popts v1.IGetGlobalPoolMemberRequest) (*lsentity.GlobalPoolMember, lserr.IError)
	UpdateGlobalPoolMember(popts v1.IUpdateGlobalPoolMemberRequest) (*lsentity.GlobalPoolMember, lserr.IError)
	DeleteGlobalPoolMember(popts v1.IDeleteGlobalPoolMemberRequest) lserr.IError
	PatchGlobalPoolMembers(popts v1.IPatchGlobalPoolMembersRequest) lserr.IError

	ListGlobalListeners(popts v1.IListGlobalListenersRequest) (*lsentity.ListGlobalListeners, lserr.IError)
	CreateGlobalListener(popts v1.ICreateGlobalListenerRequest) (*lsentity.GlobalListener, lserr.IError)
	UpdateGlobalListener(popts v1.IUpdateGlobalListenerRequest) (*lsentity.GlobalListener, lserr.IError)
	DeleteGlobalListener(popts v1.IDeleteGlobalListenerRequest) lserr.IError
	GetGlobalListener(popts v1.IGetGlobalListenerRequest) (*lsentity.GlobalListener, lserr.IError)

	ListGlobalLoadBalancers(popts v1.IListGlobalLoadBalancersRequest) (*lsentity.ListGlobalLoadBalancers, lserr.IError)
	CreateGlobalLoadBalancer(popts v1.ICreateGlobalLoadBalancerRequest) (*lsentity.GlobalLoadBalancer, lserr.IError)
	DeleteGlobalLoadBalancer(popts v1.IDeleteGlobalLoadBalancerRequest) lserr.IError
	GetGlobalLoadBalancerById(popts v1.IGetGlobalLoadBalancerByIdRequest) (*lsentity.GlobalLoadBalancer, lserr.IError)

	ListGlobalPackages(popts v1.IListGlobalPackagesRequest) (*lsentity.ListGlobalPackages, lserr.IError)
	ListGlobalRegions(popts v1.IListGlobalRegionsRequest) (*lsentity.ListGlobalRegions, lserr.IError)
	GetGlobalLoadBalancerUsageHistories(popts v1.IGetGlobalLoadBalancerUsageHistoriesRequest) (*lsentity.ListGlobalLoadBalancerUsageHistories, lserr.IError)
}
