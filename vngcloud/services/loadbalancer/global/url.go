package global

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

func listGlobalLoadBalancersUrl(psc lsclient.IServiceClient, popts IListGlobalLoadBalancersRequest) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL("global-load-balancers") + query
}

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
