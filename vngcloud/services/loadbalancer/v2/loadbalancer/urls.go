package loadbalancer

import (
	"fmt"
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

func listURL(pSc *client.ServiceClient, pOpts IListOptsBuilder) string {
	size := 999999
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		fmt.Sprintf("loadBalancers?name=&page=1&size=%d", size))
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
