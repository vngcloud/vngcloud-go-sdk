package oauth2

import (
	"github.com/vngcloud/vngcloud-go-sdk/client"
)

func authURL(pSc *client.ServiceClient) string {
	return pSc.ServiceURL("auth", "token")
}
