package entity

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/client"

type AccessToken struct {
	token string
}

func (s *AccessToken) ToSdkAuthentication() lsclient.ISdkAuthentication {
	return new(lsclient.SdkAuthentication).WithAccessToken(s.token)
}
