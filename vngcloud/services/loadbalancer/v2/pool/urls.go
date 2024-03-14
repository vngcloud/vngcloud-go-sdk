package pool

import "github.com/vngcloud/vngcloud-go-sdk/client"

func createURL(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"loadBalancers",
		pOpts.GetLoadBalancerID(),
		"pools")
}

func listPoolsBasedLoadBalancerURL(pSc *client.ServiceClient, pOpts IListPoolsBasedLoadBalancerOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"loadBalancers",
		pOpts.GetLoadBalancerID(),
		"pools")
}

func deleteURL(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"loadBalancers",
		pOpts.GetLoadBalancerID(),
		"pools",
		pOpts.GetPoolID())
}

func updatePoolMembersURL(pSc *client.ServiceClient, pOpts IUpdatePoolMembersOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"loadBalancers",
		pOpts.GetLoadBalancerID(),
		"pools",
		pOpts.GetPoolID(),
		"members")
}

func getURL(pSc *client.ServiceClient, pOpts IGetOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"loadBalancers",
		pOpts.GetLoadBalancerID(),
		"pools",
		pOpts.GetPoolID())
}

func getMemberURL(pSc *client.ServiceClient, pOpts IGetMemberOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"loadBalancers",
		pOpts.GetLoadBalancerID(),
		"pools",
		pOpts.GetPoolID(),
		"members")
}

func updateURL(pSc *client.ServiceClient, pOpts IUpdateOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"loadBalancers",
		pOpts.GetLoadBalancerID(),
		"pools",
		pOpts.GetPoolID(),
	)
}
