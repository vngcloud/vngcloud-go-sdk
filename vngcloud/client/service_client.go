package client

import (
	lstr "strings"
	ltime "time"

	ljutils "github.com/cuongpiger/joat/utils"
	lsdkErr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

type serviceClient struct {
	name        string
	endpoint    string
	moreHeaders map[string]string
	client      IHttpClient
}

func NewServiceClient() IServiceClient {
	return &serviceClient{}
}

func (s *serviceClient) WithEndpoint(pendpoint string) IServiceClient {
	s.endpoint = ljutils.NormalizeURL(pendpoint)
	return s
}

func (s *serviceClient) WithName(pname string) IServiceClient {
	s.name = pname
	return s
}

func (s *serviceClient) WithMoreHeaders(pmoreHeaders map[string]string) IServiceClient {
	s.moreHeaders = pmoreHeaders
	return s
}

func (s *serviceClient) WithKVheader(pkey string, pvalue string) IServiceClient {
	s.moreHeaders[pkey] = pvalue
	return s
}

func (s *serviceClient) WithClient(pclient IHttpClient) IServiceClient {
	s.client = pclient
	return s
}

func (s *serviceClient) ServiceURL(pparts ...string) string {
	return s.endpoint + lstr.Join(pparts, "/")
}

func (s *serviceClient) Post(purl string, preq IRequest) lsdkErr.ISdkError {
	return s.client.DoRequest(purl, preq.WithRequestMethod(MethodPost))
}

type SdkAuthentication struct {
	accessToken string
	expiresAt   int64
}

func (s *SdkAuthentication) WithAccessToken(paccessToken string) ISdkAuthentication {
	s.accessToken = paccessToken
	return s
}

func (s *SdkAuthentication) WithExpiresAt(pexpiresAt int64) ISdkAuthentication {
	s.expiresAt = pexpiresAt
	return s
}

func (s *SdkAuthentication) NeedReauth() bool {
	if s.accessToken == "" {
		return true
	}

	ea := ltime.Unix(0, s.expiresAt)
	return !(ea.Sub(ltime.Now()) >= 5*ltime.Minute)
}

func (s *SdkAuthentication) UpdateAuth(pauth ISdkAuthentication) {
	s.accessToken = pauth.GetAccessToken()
	s.expiresAt = pauth.GetExpiresAt()
}

func (s *SdkAuthentication) GetAccessToken() string {
	return s.accessToken
}

func (s *SdkAuthentication) GetExpiresAt() int64 {
	return s.expiresAt
}
