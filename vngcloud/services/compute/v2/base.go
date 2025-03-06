package v2

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

type ComputeServiceV2 struct {
	VServerClient lsclient.IServiceClient
}

func (s *ComputeServiceV2) getProjectId() string {
	return s.VServerClient.GetProjectId()
}

const (
	defaultOffsetListServerGroups = 0
	defaultLimitListServerGroups  = 10
)
