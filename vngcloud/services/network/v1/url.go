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

func deleteEndpointByIdUrl(psc lsclient.IServiceClient, popts IDeleteEndpointByIdRequest) string {
	return psc.ServiceURL(
		psc.GetZoneId(),
		psc.GetProjectId(),
		"endpoints",
		popts.GetEndpointId())
}

func listEndpointsUrl(psc lsclient.IServiceClient, popts IListEndpointsRequest) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL(psc.GetZoneId(), psc.GetProjectId(), "endpoints?") + query
}

func listTagsByEndpointIdUrl(psc lsclient.IServiceClient, popts IListTagsByEndpointIdRequest) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL(
		popts.GetProjectId(),
		"tags") + query
}

func createTagsWithEndpointIdUrl(psc lsclient.IServiceClient, popts ICreateTagsWithEndpointIdRequest) string {
	return psc.ServiceURL(
		popts.GetProjectId(),
		"tags")
}

func deleteTagOfEndpointUrl(psc lsclient.IServiceClient, popts IDeleteTagOfEndpointRequest) string {
	return psc.ServiceURL(
		popts.GetProjectId(),
		"tags",
		popts.GetTagId())
}

func updateTagValueOfEndpointUrl(psc lsclient.IServiceClient, popts IUpdateTagValueOfEndpointRequest) string {
	return psc.ServiceURL(
		popts.GetProjectId(),
		"tags",
		popts.GetTagId())
}
