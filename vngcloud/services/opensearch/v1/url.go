package v1

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

// URL layout under base = https://vdb.console.vngcloud.vn/vdb/open-search/v1
//   /{projectId}/open-search           — list clusters
//   /{projectId}/open-search/{id}      — get cluster detail

func listClustersUrl(psc lsclient.IServiceClient) string {
	return psc.ServiceURL(psc.GetProjectId(), "open-search")
}

func getClusterUrl(psc lsclient.IServiceClient, popts IGetClusterRequest) string {
	return psc.ServiceURL(psc.GetProjectId(), "open-search", popts.GetClusterId())
}

func listEndpointsUrl(psc lsclient.IServiceClient, popts IListEndpointsRequest) string {
	return psc.ServiceURL(psc.GetProjectId(), "open-search", popts.GetClusterId(), "endpoints")
}
