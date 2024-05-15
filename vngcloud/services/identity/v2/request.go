package v2

import (
	lbase64 "encoding/base64"

	lshttp "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
)

type IGetAccessTokenRequest interface {
	WithClientId(pclientId string) IGetAccessTokenRequest
	WithClientSecret(pclientSecret string) IGetAccessTokenRequest
	GetClientId() string
	ToRequest() lshttp.IRequest
}

type getAccessTokenRequest struct {
	ClientId     string
	ClientSecret string
}

func NewGetAccessTokenRequest(pclientId, pclientSecret string) IGetAccessTokenRequest {
	return &getAccessTokenRequest{
		ClientId:     pclientId,
		ClientSecret: pclientSecret,
	}
}

func (s *getAccessTokenRequest) WithClientId(pclientId string) IGetAccessTokenRequest {
	s.ClientId = pclientId
	return s
}

func (s *getAccessTokenRequest) WithClientSecret(pclientSecret string) IGetAccessTokenRequest {
	s.ClientSecret = pclientSecret
	return s
}

func (s *getAccessTokenRequest) GetClientId() string {
	return s.ClientId
}

func (s *getAccessTokenRequest) ToRequest() lshttp.IRequest {
	return lshttp.NewRequest().
		WithJsonBody(map[string]string{
			"grant_type": "client_credentials",
		}).
		WithSkipAuth(true).
		WithOkCodes(200).
		WithHeader("Authorization", "Basic "+lbase64.StdEncoding.EncodeToString([]byte(s.ClientId+":"+s.ClientSecret)))
}
