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

func updateServerSecgroupsByServerIdUrl(psc lsclient.IServiceClient, popts IUpdateServerSecgroupsByServerIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"servers",
		popts.GetServerId(),
		"update-sec-group")
}

func attachBlockVolumeUrl(psc lsclient.IServiceClient, popts IAttachBlockVolumeRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"servers",
		popts.GetServerId(),
		"attach")
}

func detachBlockVolumeUrl(psc lsclient.IServiceClient, popts IDetachBlockVolumeRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"servers",
		popts.GetServerId(),
		"detach",
	)
}

func attachFloatingIpUrl(psc lsclient.IServiceClient, popts IAttachFloatingIpRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"servers",
		popts.GetServerId(),
		"wan-ips",
		"auto",
		"attach")

}

func detachFloatingIpUrl(psc lsclient.IServiceClient, popts IDetachFloatingIpRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"servers",
		popts.GetServerId(),
		"wan-ips",
		popts.GetWanId(),
		"detach")
}

func listServerGroupPoliciesUrl(psc lsclient.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"serverGroups",
		"policies",
	)
}

func deleteServerGroupByIdUrl(psc lsclient.IServiceClient, popts IDeleteServerGroupByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"serverGroups",
		popts.GetServerGroupId(),
	)
}

func listServerGroupsUrl(psc lsclient.IServiceClient, popts IListServerGroupsRequest) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL(psc.GetProjectId(), "serverGroups") + query
}

func createServerGroupUrl(psc lsclient.IServiceClient, _ ICreateServerGroupRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"serverGroups",
	)
}
