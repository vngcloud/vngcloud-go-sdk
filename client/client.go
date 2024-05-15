package client

import (
	lctx "context"
	ltime "time"

	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsgateway "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/gateway"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lssvcIdentityV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/identity/v2"
)

var (
	_ IClient = new(client)
)

type (
	client struct {
		context    lctx.Context
		httpClient lsclient.IHttpClient

		iamGateway     lsgateway.IIamGateway
		vserverGateway lsgateway.IVServerGateway
		vlbGateway     lsgateway.IVLBGateway
		vbackupGateway lsgateway.VBackUpGateway
	}
)

func NewClient(pctx lctx.Context) IClient {
	c := new(client)
	c.context = pctx

	return c
}

func NewSdkConfigure() ISdkConfigure {
	return &sdkConfigure{}
}

func (s *client) WithHttpClient(pclient lsclient.IHttpClient) IClient {
	s.httpClient = pclient
	return s
}

func (s *client) WithContext(pctx lctx.Context) IClient {
	s.context = pctx
	return s
}

func (s *client) WithAuthOption(pauthOpts lsclient.AuthOpts, pauthConfig ISdkConfigure) IClient {
	if s.httpClient == nil {
		s.httpClient = lsclient.NewHttpClient()
	}

	switch pauthOpts {
	case lsclient.IamOauth2:
		s.httpClient.WithReauthFunc(lsclient.IamOauth2, s.usingIamOauth2AsAuthOption(pauthConfig)).
			WithKvDefaultHeaders("Content-Type", "application/json")
	default:
		s.httpClient.WithReauthFunc(lsclient.IamOauth2, s.usingIamOauth2AsAuthOption(pauthConfig)).
			WithKvDefaultHeaders("Content-Type", "application/json")
	}

	return s
}

func (s *client) WithRetryCount(pretry int) IClient {
	if s.httpClient == nil {
		s.httpClient = lsclient.NewHttpClient()
	}

	s.httpClient.WithRetryCount(pretry)
	return s
}

func (s *client) WithKvDefaultHeaders(pargs ...string) IClient {
	if s.httpClient == nil {
		s.httpClient = lsclient.NewHttpClient()
	}

	s.httpClient.WithKvDefaultHeaders(pargs...)
	return s
}

func (s *client) WithDelay(pdelay ltime.Duration) IClient {
	if s.httpClient == nil {
		s.httpClient = lsclient.NewHttpClient()
	}

	s.httpClient.WithDelay(pdelay)
	return s
}

func (s *client) WithSleep(psleep ltime.Duration) IClient {
	if s.httpClient == nil {
		s.httpClient = lsclient.NewHttpClient()
	}

	s.httpClient.WithSleep(psleep)
	return s
}

func (s *client) Configure(psdkCfg ISdkConfigure) IClient {
	if s.httpClient == nil {
		s.httpClient = lsclient.NewHttpClient()
	}

	s.httpClient.WithReauthFunc(lsclient.IamOauth2, s.usingIamOauth2AsAuthOption(psdkCfg))
	if s.iamGateway == nil {
		s.iamGateway = lsgateway.NewIamGateway(psdkCfg.GetIamEndpoint(), s.httpClient)
	}

	if s.vserverGateway == nil {
		s.vserverGateway = lsgateway.NewVServerGateway(psdkCfg.GetVServerEndpoint(), s.httpClient)
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

func (s *client) usingIamOauth2AsAuthOption(pauthConfig ISdkConfigure) func() (lsclient.ISdkAuthentication, lserr.ISdkError) {
	authFunc := func() (lsclient.ISdkAuthentication, lserr.ISdkError) {
		token, err := s.iamGateway.V2().IdentityService().GetAccessToken(
			lssvcIdentityV2.NewGetAccessTokenRequest(pauthConfig.GetClientId(), pauthConfig.GetClientSecret()))
		if err != nil {
			return nil, err
		}

		return token.ToSdkAuthentication(), nil
	}

	return authFunc
}
