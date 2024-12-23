package global

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

type LoadBalancerServiceGlobal struct {
	VLBClient     lsclient.IServiceClient
	VServerClient lsclient.IServiceClient
}

func (s *LoadBalancerServiceGlobal) getProjectId() string {
	return s.VLBClient.GetProjectId()
}

const (
	defaultOffsetListGlobalLoadBalancer = 0
	defaultLimitListGlobalLoadBalancer  = 10
)
