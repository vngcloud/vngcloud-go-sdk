package v2

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

func getSecgroupUrl(psc lsclient.IServiceClient, popts IGetSecgroupByIdRequest) string {
	return psc.ServiceURL(psc.GetProjectId(), "secgroups", popts.GetSecgroupId())
}

func createSecgroupUrl(psc lsclient.IServiceClient) string {
	return psc.ServiceURL(psc.GetProjectId(), "secgroups")
}
