package v2

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

func createServerUrl(psc lsclient.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"servers")
}

func getServerByIdUrl(psc lsclient.IServiceClient, popts IGetServerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"servers",
		popts.GetServerId())
}

func deleteServerByIdUrl(psc lsclient.IServiceClient, popts IDeleteServerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"servers",
		popts.GetServerId())
}
