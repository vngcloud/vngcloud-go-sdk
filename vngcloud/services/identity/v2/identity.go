package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lsdkErr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

type IdentityServiceV2 struct {
	IamClient lsclient.IServiceClient
}

func (s *IdentityServiceV2) GetAccessToken(popts IGetAccessTokenRequest) (*lsentity.AccessToken, lsdkErr.ISdkError) {
	resp := new(GetAccessTokenResponse)
	errResp := lsdkErr.NewErrorResponse()
	req := popts.ToRequest().WithJsonResponse(resp).WithJsonError(errResp)
	sdkErr := s.IamClient.Post(getAccessTokenUrl(s.IamClient), req)

	if sdkErr != nil {
		return nil, lsdkErr.ErrorHandler(sdkErr.GetError(),
			lsdkErr.WithErrorAuthenticationFailed(sdkErr.GetError(), errResp)).WithKVparameters(
			"clientId", popts.GetClientId())
	}

	return resp.ToEntityAccessToken(), nil
}
