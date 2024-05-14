package v2

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lsdkErr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lssc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/service_client"
)

type IdentityServiceV2 struct {
	ServiceClient lssc.IServiceClient
}

func (s *IdentityServiceV2) GetAccessToken(popts IGetAccessTokenRequest) (*lsentity.AccessToken, lsdkErr.ISdkError) {
	return nil, nil
}
