package volume

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsvolumeSvcV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/volume/v1"
	lsvolumeSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/volume/v2"
)

func NewVolumeServiceV2(psvcClient lsclient.IServiceClient) IVolumeServiceV2 {
	return &lsvolumeSvcV2.VolumeServiceV2{
		VServerClient: psvcClient,
	}
}

func NewVolumeServiceV1(psvcClient lsclient.IServiceClient) IVolumeServiceV1 {
	return &lsvolumeSvcV1.VolumeServiceV1{
		VServerClient: psvcClient,
	}
}
