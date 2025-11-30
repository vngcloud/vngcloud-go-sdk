package v2

import (
	ltime "time"

	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
)

type GetAccessTokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	TokenType        string `json:"token_type"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
}

func (s *GetAccessTokenResponse) ToEntityAccessToken() *lsentity.AccessToken {
	return &lsentity.AccessToken{
		Token:     s.AccessToken,
		ExpiresAt: ltime.Now().Add(ltime.Duration(s.ExpiresIn) * ltime.Second).UnixNano(),
	}
}
