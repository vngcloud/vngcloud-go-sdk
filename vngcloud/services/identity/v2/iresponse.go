package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type IGetAccessTokenResponse interface {
	ToEntityAccessToken() *lsentity.AccessToken
}
