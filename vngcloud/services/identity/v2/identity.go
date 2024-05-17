package v2

import (
	lbase64 "encoding/base64"
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

type IdentityServiceV2 struct {
	IamClient lsclient.IServiceClient
}

func (s *IdentityServiceV2) GetAccessToken(popts IGetAccessTokenRequest) (*lsentity.AccessToken, lserr.ISdkError) {
	url := getAccessTokenUrl(s.IamClient)
	resp := new(GetAccessTokenResponse)
	errResp := lserr.NewErrorResponse()
	req := lsclient.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithSkipAuth(true).
		WithJsonError(errResp).
		WithJsonBody(popts.ToRequestBody()).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithHeader("Authorization", "Basic "+lbase64.StdEncoding.EncodeToString([]byte(popts.GetClientId()+":"+popts.GetClientSecret())))

	if _, sdkErr := s.IamClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr,
			lserr.WithErrorTooManyFailedLogin(errResp),
			lserr.WithErrorAuthenticationFailed(errResp)).WithKVparameters("clientId", popts.GetClientId())
	}

	return resp.ToEntityAccessToken(), nil
}
