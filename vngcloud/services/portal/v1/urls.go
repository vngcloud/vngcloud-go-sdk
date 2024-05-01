package v1

import "github.com/vngcloud/vngcloud-go-sdk/client"

func getURL(pSc *client.ServiceClient, pProjectID string) string {
	return pSc.ServiceURL("projects", pProjectID, "detail")
}
