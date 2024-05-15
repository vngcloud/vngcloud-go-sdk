package client

import (
	lctx "context"
	lsync "sync"
	ltime "time"

	lshttp "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsgateway "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/gateway"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lssvcIdentityV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/identity/v2"
)

var (
	_ IClient = new(client)
)

const (
	IamOauth2 authOpts = "IamOauth2"
)

type (
	client struct {
		context    lctx.Context
		httpClient lshttp.IHttpClient

		iamGateway     lsgateway.IIamGateway
		vserverGateway lsgateway.IVServerGateway
		vlbGateway     lsgateway.IVLBGateway
		vbackupGateway lsgateway.VBackUpGateway

		mut       *lsync.RWMutex
		reauthMut *reauthLock
	}

	reauthLock struct {
		ongoing *reauthFuture

		lsync.RWMutex
	}

	reauthFuture struct {
		done   chan struct{}
		sdkErr lserr.ISdkError
	}

	authOpts string
)

func NewClient() IClient {
	c := new(client)
	c.context = lctx.TODO()

	return c
}

func NewSdkConfigure() ISdkConfigure {
	return &sdkConfigure{}
}

func (s *client) WithHttpClient(pclient lshttp.IHttpClient) IClient {
	s.httpClient = pclient
	return s
}

func (s *client) WithContext(pctx lctx.Context) IClient {
	s.context = pctx
	return s
}

func (s *client) WithAuthOption(pauthOpts authOpts, pauthConfig ISdkConfigure) IClient {
	if s.httpClient == nil {
		s.httpClient = lshttp.NewHttpClient()
	}

	switch pauthOpts {
	case IamOauth2:
		s.httpClient.WithReauthFunc(s.usingIamOauth2AsAuthOption(pauthConfig))
	default:
		s.httpClient.WithReauthFunc(s.usingIamOauth2AsAuthOption(pauthConfig))
	}

	return s
}

func (s *client) WithRetryCount(pretry int) IClient {
	if s.httpClient == nil {
		s.httpClient = lshttp.NewHttpClient()
	}

	s.httpClient.WithRetryCount(pretry)
	return s
}

func (s *client) WithKvDefaultHeaders(pargs ...string) IClient {
	if s.httpClient == nil {
		s.httpClient = lshttp.NewHttpClient()
	}

	s.httpClient.WithKvDefaultHeaders(pargs...)
	return s
}

func (s *client) WithDelay(pdelay ltime.Duration) IClient {
	if s.httpClient == nil {
		s.httpClient = lshttp.NewHttpClient()
	}

	s.httpClient.WithDelay(pdelay)
	return s
}

func (s *client) WithSleep(psleep ltime.Duration) IClient {
	if s.httpClient == nil {
		s.httpClient = lshttp.NewHttpClient()
	}

	s.httpClient.WithSleep(psleep)
	return s
}

func (s *client) Configure(psdkCfg ISdkConfigure) IClient {
	if s.httpClient == nil {
		s.httpClient = lshttp.NewHttpClient()
	}

	s.httpClient.WithReauthFunc(s.usingIamOauth2AsAuthOption(psdkCfg))

	if s.iamGateway == nil {
		s.iamGateway = lsgateway.NewIamGateway(psdkCfg.GetIamEndpoint(), s.httpClient)
	}

	return s
}

func (s *client) IamGateway() lsgateway.IIamGateway {
	return s.iamGateway
}

func (s *client) VServerGateway() lsgateway.IVServerGateway {
	return s.vserverGateway
}

func (s *client) VLBGateway() lsgateway.IVLBGateway {
	return s.vlbGateway
}

func (s *client) VBackUpGateway() lsgateway.VBackUpGateway {
	return s.vbackupGateway
}

func (s *client) usingIamOauth2AsAuthOption(pauthConfig ISdkConfigure) func() (lshttp.ISdkAuthentication, lserr.ISdkError) {
	authFunc := func() (lshttp.ISdkAuthentication, lserr.ISdkError) {
		token, err := s.iamGateway.V2().IdentityService().GetAccessToken(
			lssvcIdentityV2.NewGetAccessTokenRequest(pauthConfig.GetClientId(), pauthConfig.GetClientSecret()))
		if err != nil {
			return nil, err
		}

		return token.ToSdkAuthentication(), nil
	}

	return authFunc
}
