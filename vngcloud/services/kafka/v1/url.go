package v1

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

// Path layout under the vDB Kafka service base URL (/vdb-kafka/...):
//   clusters
//   clusters/{clusterId}
//   clusters/{clusterId}/users
//   clusters/{clusterId}/users/{userId}/authen-creds
//   clusters/{clusterId}/topics

func listClustersUrl(psc lsclient.IServiceClient) string {
	return psc.ServiceURL("clusters")
}

func getClusterUrl(psc lsclient.IServiceClient, popts IGetClusterRequest) string {
	return psc.ServiceURL("clusters", popts.GetClusterId())
}

func listUsersUrl(psc lsclient.IServiceClient, popts IListUsersRequest) string {
	return psc.ServiceURL("clusters", popts.GetClusterId(), "users")
}

func getUserAuthenCredsUrl(psc lsclient.IServiceClient, popts IGetUserAuthenCredsRequest) string {
	return psc.ServiceURL("clusters", popts.GetClusterId(), "users", popts.GetUserId(), "authen-creds")
}

func listTopicsUrl(psc lsclient.IServiceClient, popts IListTopicsRequest) string {
	return psc.ServiceURL("clusters", popts.GetClusterId(), "topics")
}

func createTopicUrl(psc lsclient.IServiceClient, popts ICreateTopicRequest) string {
	return psc.ServiceURL("clusters", popts.GetClusterId(), "topics")
}
