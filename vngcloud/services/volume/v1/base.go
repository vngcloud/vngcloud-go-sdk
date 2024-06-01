package v1

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
)

type VolumeServiceV1 struct {
	VServerClient lsclient.IServiceClient
}

func (s *VolumeServiceV1) getProjectId() string {
	return s.VServerClient.GetProjectId()
}
