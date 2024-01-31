package server

import "github.com/vngcloud/vngcloud-go-sdk/client"

func getServerURL(pSc *client.ServiceClient, pOpts IGetOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"servers",
		pOpts.GetServerID())
}
