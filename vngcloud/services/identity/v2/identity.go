package v2

import (
	lbase64 "encoding/base64"
	"fmt"

	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lsdkErr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

type IdentityServiceV2 struct {
	IamClient lsclient.IServiceClient
}

func (s *IdentityServiceV2) GetAccessToken(popts IGetAccessTokenRequest) (*lsentity.AccessToken, lsdkErr.ISdkError) {
	url := getAccessTokenUrl(s.IamClient)
	resp := new(GetAccessTokenResponse)
	errResp := lsdkErr.NewErrorResponse()
	req := lsclient.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp).
		WithJsonBody(popts.ToRequestBody()).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithHeader("Authorization", "Basic "+lbase64.StdEncoding.EncodeToString([]byte(popts.GetClientId()+":"+popts.GetClientSecret())))

	_, sdkErr := s.IamClient.Post(url, req)
	if sdkErr != nil {
		fmt.Println("sdkErr: ", sdkErr.GetError())
		return nil, lsdkErr.ErrorHandler(sdkErr.GetError())
	}

	return resp.ToEntityAccessToken(), nil
}
