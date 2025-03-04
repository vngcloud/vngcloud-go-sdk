package glb

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	v1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/glb/v1"
)

func NewGLBServiceV1(psvcClient lsclient.IServiceClient) IGLBServiceV1 {
	return &v1.GLBServiceV1{
		VLBClient: psvcClient,
	}
}
