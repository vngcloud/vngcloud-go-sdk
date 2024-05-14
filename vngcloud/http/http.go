package http

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/client"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lhttp "net/http"
	ltime "time"
)

type httpClient struct {
	retryCount int
	delay      ltime.Duration
	sleep      ltime.Duration
	client     *lhttp.Client

	reauthFunc func() (lsclient.ISdkAuthentication, lserr.ISdkError)

	accessToken string
}

func NewHttpClient() IHttpClient {
	return &httpClient{
		retryCount: 0,
		client:     new(lhttp.Client),
	}
}

func (s *httpClient) WithRetryCount(pretryCount int) IHttpClient {
	s.retryCount = pretryCount
	return s
}

func (s *httpClient) WithDelay(pdelay ltime.Duration) IHttpClient {
	s.delay = pdelay
	return s
}

func (s *httpClient) WithSleep(psleep ltime.Duration) IHttpClient {
	s.sleep = psleep
	return s
}

func (s *httpClient) WithReauthFunc(preauthFunc func() (lsclient.ISdkAuthentication, lserr.ISdkError)) IHttpClient {
	s.reauthFunc = preauthFunc
	return s
}
