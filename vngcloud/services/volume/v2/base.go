package v2

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

type VolumeServiceV2 struct {
	VServerClient lsclient.IServiceClient
}

func (s *VolumeServiceV2) getProjectId() string {
	return s.VServerClient.GetProjectId()
}
