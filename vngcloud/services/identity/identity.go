package identity

import (
	lssc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/service_client"
	lsidentitySvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/identity/v2"
)

func NewIdentityService(psvcClient lssc.IServiceClient) IIdentityServiceV2 {
	return &lsidentitySvcV2.IdentityServiceV2{
		IamClient: psvcClient,
	}
}
