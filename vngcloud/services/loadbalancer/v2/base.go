package v2

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

type LoadBalancerServiceV2 struct {
	VLBClient lsclient.IServiceClient
}

func (s *LoadBalancerServiceV2) getProjectId() string {
	return s.VLBClient.GetProjectId()
}