package identity

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsidentitySvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/identity/v2"
)

func NewIdentityService(psvcClient lsclient.IServiceClient) IIdentityServiceV2 {
	return &lsidentitySvcV2.IdentityServiceV2{
		IamClient: psvcClient,
	}
}
