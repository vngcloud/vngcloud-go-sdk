package v1

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

type NetworkServiceV1 struct {
	VNetworkClient lsclient.IServiceClient
}

func (s *NetworkServiceV1) getProjectId() string {
	return s.VNetworkClient.GetProjectId()
}

type NetworkServiceInternalV1 struct {
	VNetworkClient lsclient.IServiceClient
}

func (s *NetworkServiceInternalV1) getProjectId() string {
	return s.VNetworkClient.GetProjectId()
}
