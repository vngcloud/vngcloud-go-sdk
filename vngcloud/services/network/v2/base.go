package v2

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

type NetworkServiceV2 struct {
	VserverClient lsclient.IServiceClient
}

func (s *NetworkServiceV2) getProjectId() string {
	return s.VserverClient.GetProjectId()
}
