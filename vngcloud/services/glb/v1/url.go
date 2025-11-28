package v1

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

// ------------------------------------------------------------

func createGlobalPoolUrl(psc lsclient.IServiceClient, popts ICreateGlobalPoolRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
	)
}

func listGlobalPoolsUrl(psc lsclient.IServiceClient, popts IListGlobalPoolsRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
	)
}

func updateGlobalPoolUrl(psc lsclient.IServiceClient, popts IUpdateGlobalPoolRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
		popts.GetPoolId(),
	)
}

func deleteGlobalPoolUrl(psc lsclient.IServiceClient, popts IDeleteGlobalPoolRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
		popts.GetPoolId(),
	)
}

func listGlobalPoolMembersUrl(psc lsclient.IServiceClient, popts IListGlobalPoolMembersRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
		popts.GetPoolId(),
		"pool-members",
	)
}

func patchGlobalPoolMemberUrl(psc lsclient.IServiceClient, popts IPatchGlobalPoolMemberRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
		popts.GetPoolId(),
		"pool-members",
	)
}

func getGlobalPoolMemberUrl(psc lsclient.IServiceClient, popts IGetGlobalPoolMemberRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
		popts.GetPoolId(),
		"pool-members",
		popts.GetPoolMemberId(),
	)
}

func listGlobalListenersUrl(psc lsclient.IServiceClient, popts IListGlobalListenersRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-listeners",
	)
}

func createGlobalListenerUrl(psc lsclient.IServiceClient, popts ICreateGlobalListenerRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-listeners",
	)
}

func updateGlobalListenerUrl(psc lsclient.IServiceClient, popts IUpdateGlobalListenerRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-listeners",
		popts.GetListenerId(),
	)
}

func deleteGlobalListenerUrl(psc lsclient.IServiceClient, popts IDeleteGlobalListenerRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-listeners",
		popts.GetListenerId(),
	)
}

func getGlobalListenerUrl(psc lsclient.IServiceClient, popts IGetGlobalListenerRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-listeners",
		popts.GetListenerId(),
	)
}

// --------------------------------------------------------

func listGlobalLoadBalancersUrl(psc lsclient.IServiceClient, popts IListGlobalLoadBalancersRequest) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL("global-load-balancers") + query
}

func createGlobalLoadBalancerUrl(psc lsclient.IServiceClient, popts ICreateGlobalLoadBalancerRequest) string {
	return psc.ServiceURL("global-load-balancers")
}

func deleteGlobalLoadBalancerUrl(psc lsclient.IServiceClient, popts IDeleteGlobalLoadBalancerRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
	)
}

func getGlobalLoadBalancerByIdUrl(psc lsclient.IServiceClient, popts IGetGlobalLoadBalancerByIdRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
	)
}
