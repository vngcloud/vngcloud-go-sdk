package flavor

import (
	"github.com/vngcloud/vngcloud-go-sdk/client"
)

func getURL(pSc *client.ServiceClient, pOpts IGetOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"flavors",
		pOpts.GetFlavorID())
}
