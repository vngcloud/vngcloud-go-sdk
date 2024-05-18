package v2

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

type ComputeServiceV2 struct {
	VserverClient lsclient.IServiceClient
}

func (s *ComputeServiceV2) getProjectId() string {
	return s.VserverClient.GetProjectId()
}
