package loadbalancer

import (
	"github.com/vngcloud/vngcloud-go-sdk/client"
)

func createURL(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"loadBalancers")
}

func getURL(pSc *client.ServiceClient, pOpts IGetOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"loadBalancers",
		pOpts.GetLoadBalancerID())
}

func deleteURL(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"loadBalancers",
		pOpts.GetLoadBalancerID())
}

func listBySubnetIDURL(pSc *client.ServiceClient, pOpts IListBySubnetIDOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"loadBalancers",
		"subnet",
		pOpts.GetSubnetID())
}

func listURL(psc *client.ServiceClient, popts IListOptsBuilder) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL(popts.GetProjectID(), "loadBalancers") + query
}

func updateURL(pSc *client.ServiceClient, pOpts IUpdateOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"loadBalancers",
		pOpts.GetLoadBalancerID(),
		"resize")
}

func createTagUrl(psc *client.ServiceClient, popts ICreateTagOptsBuilder) string {
	return psc.ServiceURL(
		popts.GetProjectID(),
		"tag",
		"resource",
		popts.GetLoadBalancerID())
}

func listTagsURL(psc *client.ServiceClient, popts IListTagsOptsBuilder) string {
	return psc.ServiceURL(
		popts.GetProjectID(),
		"tag",
		"resource",
		popts.GetLoadBalancerID())
}

func updateTagURL(psc *client.ServiceClient, popts IUpdateTagOptsBuilder) string {
	return psc.ServiceURL(
		popts.GetProjectID(),
		"tag",
		"resource",
		popts.GetLoadBalancerID())
}
