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
