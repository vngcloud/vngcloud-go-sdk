package v1

import (
	"github.com/vngcloud/vngcloud-go-sdk/client"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

// Get gets the under-user information from the portal information
func Get(pSc *client.ServiceClient, pProjectID string) (*objects.Portal, error) {
	portal := new(objects.Portal)
	_, err := pSc.Get(getURL(pSc, pProjectID), &client.RequestOpts{
		JSONResponse: portal,
		OkCodes:      []int{200},
	})

	return portal, err
}
