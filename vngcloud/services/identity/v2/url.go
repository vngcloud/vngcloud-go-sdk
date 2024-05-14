package v2

import lssc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/service_client"

func authURL(psc lssc.IServiceClient) string {
	return psc.ServiceURL("auth", "token")
}
