package v2

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

func getSecgroupByIdUrl(psc lsclient.IServiceClient, popts IGetSecgroupByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups",
		popts.GetSecgroupId())
}

func createSecgroupUrl(psc lsclient.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups")
}

func deleteSecgroupByIdUrl(psc lsclient.IServiceClient, popts IDeleteSecgroupByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups",
		popts.GetSecgroupId())
}

func createSecgroupRuleUrl(psc lsclient.IServiceClient, popts ICreateSecgroupRuleRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups",
		popts.GetSecgroupId(),
		"secgroupRules")
}
