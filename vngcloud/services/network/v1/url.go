package v1

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

func getEndpointByIdUrl(psc lsclient.IServiceClient, popts IGetEndpointByIdRequest) string {
	return psc.ServiceURL(
		psc.GetZoneId(),
		psc.GetProjectId(),
		"endpoints",
		popts.GetEndpointId())
}

func createEndpointUrl(psc lsclient.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetZoneId(),
		psc.GetProjectId(),
		"endpoints")
}
