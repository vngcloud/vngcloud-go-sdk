package compute

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lscomputeSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/compute/v2"
)

func NewComputeService(psvcClient lsclient.IServiceClient) IComputeServiceV2 {
	return &lscomputeSvcV2.ComputeServiceV2{
		VserverClient: psvcClient,
	}
}