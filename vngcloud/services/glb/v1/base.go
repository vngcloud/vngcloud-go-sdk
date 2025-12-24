package v1

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

type GLBServiceV1 struct {
	VLBClient     lsclient.IServiceClient
	VServerClient lsclient.IServiceClient
}

const (
	defaultOffsetListGlobalLoadBalancer = 0
	defaultLimitListGlobalLoadBalancer  = 10
)
