package inter

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

type LoadBalancerServiceInternal struct {
	VLBClient lsclient.IServiceClient
}
