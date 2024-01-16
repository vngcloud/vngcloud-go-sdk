package policy

import (
	// "fmt"
	"github.com/vngcloud/vngcloud-go-sdk/client"
)

func listURL(pSc *client.ServiceClient, pOpts IListOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"loadBalancers",
		pOpts.GetLoadBalancerID(),
		"listeners",
		pOpts.GetListenerID(),
		"l7policies",
	)
}

func createURL(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"loadBalancers",
		pOpts.GetLoadBalancerID(),
		"listeners",
		pOpts.GetListenerID(),
		"l7policies",
	)
}

func deleteURL(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"loadBalancers",
		pOpts.GetLoadBalancerID(),
		"listeners",
		pOpts.GetListenerID(),
		"l7policies",
		pOpts.GetPolicyID(),
	)
}

func updateURL(pSc *client.ServiceClient, pOpts IUpdateOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"loadBalancers",
		pOpts.GetLoadBalancerID(),
		"listeners",
		pOpts.GetListenerID(),
		"l7policies",
		pOpts.GetPolicyID(),
	)
}
