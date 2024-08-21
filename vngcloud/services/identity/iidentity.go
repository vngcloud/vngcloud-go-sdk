package identity

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lsdkErr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lsidentitySvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/identity/v2"
)

type IIdentityServiceV2 interface {
	GetAccessToken(popts lsidentitySvcV2.IGetAccessTokenRequest) (*lsentity.AccessToken, lsdkErr.IError)
}
