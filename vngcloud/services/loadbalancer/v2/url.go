package v2

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

func createLoadBalancerUrl(psc lsclient.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers")
}

func resizeLoadBalancerUrl(psc lsclient.IServiceClient, popts IResizeLoadBalancerRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"resize")
}

func listLoadBalancerPackagesUrl(psc lsclient.IServiceClient, popts IListLoadBalancerPackagesRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers", "packages") + "?zoneId=" + popts.GetZoneId()
}

func getLoadBalancerByIdUrl(psc lsclient.IServiceClient, popts IGetLoadBalancerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId())
}

func listLoadBalancersUrl(psc lsclient.IServiceClient, popts IListLoadBalancersRequest) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL(psc.GetProjectId(), "loadBalancers") + query
}

func getPoolHealthMonitorByIdUrl(psc lsclient.IServiceClient, popts IGetPoolHealthMonitorByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools",
		popts.GetPoolId(),
		"healthMonitor")
}

func createPoolUrl(psc lsclient.IServiceClient, popts ICreatePoolRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools")
}

func updatePoolUrl(psc lsclient.IServiceClient, popts IUpdatePoolRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools",
		popts.GetPoolId())
}

func createListenerUrl(psc lsclient.IServiceClient, popts ICreateListenerRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners")
}

func updateListenerUrl(psc lsclient.IServiceClient, popts IUpdateListenerRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId())
}

func listListenersByLoadBalancerIdUrl(psc lsclient.IServiceClient, popts IListListenersByLoadBalancerIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners")
}

func listPoolsByLoadBalancerIdUrl(psc lsclient.IServiceClient, popts IListPoolsByLoadBalancerIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools")
}

func updatePoolMembersUrl(psc lsclient.IServiceClient, popts IUpdatePoolMembersRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools",
		popts.GetPoolId(),
		"members")
}

func listPoolMembersUrl(psc lsclient.IServiceClient, popts IListPoolMembersRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools",
		popts.GetPoolId(),
		"members")
}

func deletePoolByIdUrl(psc lsclient.IServiceClient, popts IDeletePoolByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools",
		popts.GetPoolId())
}

func deleteListenerByIdUrl(psc lsclient.IServiceClient, popts IDeleteListenerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId())
}

func deleteLoadBalancerByIdUrl(psc lsclient.IServiceClient, popts IDeleteLoadBalancerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId())
}

func listTagsUrl(psc lsclient.IServiceClient, popts IListTagsRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"tag",
		"resource",
		popts.GetLoadBalancerId())
}

func createTagsUrl(psc lsclient.IServiceClient, popts ICreateTagsRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"tag",
		"resource",
		popts.GetLoadBalancerId())
}

func updateTagsUrl(psc lsclient.IServiceClient, popts IUpdateTagsRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"tag",
		"resource",
		popts.GetLoadBalancerId())
}

// Policy

func listPoliciesUrl(psc lsclient.IServiceClient, popts IListPoliciesRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId(),
		"l7policies",
	)
}

func createPolicyUrl(psc lsclient.IServiceClient, popts ICreatePolicyRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId(),
		"l7policies",
	)
}

func getPolicyByIdUrl(psc lsclient.IServiceClient, popts IGetPolicyByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId(),
		"l7policies",
		popts.GetPolicyId(),
	)
}

func updatePolicyUrl(psc lsclient.IServiceClient, popts IUpdatePolicyRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId(),
		"l7policies",
		popts.GetPolicyId(),
	)
}

func deletePolicyByIdUrl(psc lsclient.IServiceClient, popts IDeletePolicyByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId(),
		"l7policies",
		popts.GetPolicyId(),
	)
}

func reorderPoliciesUrl(psc lsclient.IServiceClient, popts IReorderPoliciesRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId(),
		"reorderL7Policies")
}

func getPoolByIdUrl(psc lsclient.IServiceClient, popts IGetPoolByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools",
		popts.GetPoolId())
}

func getListenerByIdUrl(psc lsclient.IServiceClient, popts IGetListenerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId())
}

func resizeLoadBalancerByIdUrl(psc lsclient.IServiceClient, popts IResizeLoadBalancerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"resize")
}

func scaleLoadBalancerUrl(psc lsclient.IServiceClient, popts IScaleLoadBalancerRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"rebalancing")
}

// --------------------------------------------------------

func listCertificatesUrl(psc lsclient.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"cas")
}

func createCertificateUrl(psc lsclient.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"cas")
}

func getCertificateByIdUrl(psc lsclient.IServiceClient, popts IGetCertificateByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"cas",
		popts.GetCertificateId())
}

func deleteCertificateByIdUrl(psc lsclient.IServiceClient, popts IDeleteCertificateByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"cas",
		popts.GetCertificateId())
}
