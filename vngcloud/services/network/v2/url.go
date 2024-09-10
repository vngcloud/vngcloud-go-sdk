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

func deleteSecgroupRuleByIdUrl(psc lsclient.IServiceClient, popts IDeleteSecgroupRuleByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups",
		popts.GetSecgroupId(),
		"secgroupRules",
		popts.GetSecgroupRuleId())
}

func listSecgroupRulesBySecgroupIdUrl(psc lsclient.IServiceClient, popts IListSecgroupRulesBySecgroupIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"secgroups",
		popts.GetSecgroupId(),
		"secGroupRules")
}

func getNetworkByIdUrl(psc lsclient.IServiceClient, popts IGetNetworkByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"networks",
		popts.GetNetworkId())
}

func getSubnetByIdUrl(psc lsclient.IServiceClient, popts IGetSubnetByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"networks",
		popts.GetNetworkId(),
		"subnets",
		popts.GetSubnetId())
}

func updateSubnetByIdUrl(psc lsclient.IServiceClient, popts IGetSubnetByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"networks",
		popts.GetNetworkId(),
		"subnets",
		popts.GetSubnetId())
}

func getAllAddressPairByVirtualSubnetIdUrl(psc lsclient.IServiceClient, popts IGetAllAddressPairByVirtualSubnetIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"virtual-subnets",
		popts.GetVirtualSubnetId(),
		"addressPairs")
}

func setAddressPairInVirtualSubnetUrl(psc lsclient.IServiceClient, popts IGetAllAddressPairByVirtualSubnetIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"virtual-subnets",
		popts.GetVirtualSubnetId(),
		"addressPairs")
}

func deleteAddressPairUrl(psc lsclient.IServiceClient, popts IDeleteAddressPairRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"virtual-subnets",
		"addressPairs",
		popts.GetAddressPairID())
}
