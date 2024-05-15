package v2

import lssc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/service_client"

func getAccessTokenUrl(psc lssc.IServiceClient) string {
	return psc.ServiceURL("auth", "token")
}
