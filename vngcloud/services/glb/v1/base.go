package v1

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

type GLBServiceV1 struct {
	VLBClient     lsclient.IServiceClient
	VServerClient lsclient.IServiceClient
}

func (s *GLBServiceV1) getProjectId() string {
	return s.VLBClient.GetProjectId()
}

const (
	defaultOffsetListGlobalLoadBalancer = 0
	defaultLimitListGlobalLoadBalancer  = 10
)
