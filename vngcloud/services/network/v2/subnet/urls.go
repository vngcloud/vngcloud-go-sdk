package subnet

import "github.com/vngcloud/vngcloud-go-sdk/client"

func getURL(pSc *client.ServiceClient, pOpts IGetOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"networks",
		pOpts.GetNetworkUUID(),
		"subnets",
		pOpts.GetSubnetUUID())
}

func listByNetworkIDURL(pSc *client.ServiceClient, pOpts IListByNetworkIDOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"networks",
		pOpts.GetNetworkUUID(),
		"subnets")
}
