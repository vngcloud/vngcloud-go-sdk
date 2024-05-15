package v2

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lsdkErr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lssc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/service_client"
)

type IdentityServiceV2 struct {
	IamClient lssc.IServiceClient
}

func (s *IdentityServiceV2) GetAccessToken(popts IGetAccessTokenRequest) (*lsentity.AccessToken, lsdkErr.ISdkError) {
	sdkErr := s.IamClient.Post(getAccessTokenUrl(s.IamClient), popts.ToRequest())
	if sdkErr != nil {
		return nil, sdkErr
	}

	return nil, nil
}
