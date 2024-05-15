package v1

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

func getPortalInfoUrl(psc lsclient.IServiceClient, popts IGetPortalInfoRequest) string {
	return psc.ServiceURL("projects", popts.GetBackEndProjectId(), "detail")
}
