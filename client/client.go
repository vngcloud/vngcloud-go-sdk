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
		projectId  string
		zoneId     string
		userId     string
		httpClient lsclient.IHttpClient
		userAgent  string
		authOpt    lsclient.AuthOpts

		iamGateway      lsgateway.IIamGateway
		vserverGateway  lsgateway.IVServerGateway
		vlbGateway      lsgateway.IVLBGateway
		vbackupGateway  lsgateway.IVBackUpGateway
		vnetworkGateway lsgateway.IVNetworkGateway
		glbGateway      lsgateway.IGLBGateway
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
		s.httpClient = lsclient.NewHttpClient(s.context)
	}

	s.authOpt = pauthOpts // Assign the auth option to the client

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
		s.httpClient = lsclient.NewHttpClient(s.context)
	}

	s.httpClient.WithRetryCount(pretry)
	return s
}

func (s *client) WithKvDefaultHeaders(pargs ...string) IClient {
	if s.httpClient == nil {
		s.httpClient = lsclient.NewHttpClient(s.context)
	}

	s.httpClient.WithKvDefaultHeaders(pargs...)
	return s
}

func (s *client) WithSleep(psleep ltime.Duration) IClient {
	if s.httpClient == nil {
		s.httpClient = lsclient.NewHttpClient(s.context)
	}

	s.httpClient.WithSleep(psleep)
	return s
}

func (s *client) WithProjectId(pprojectId string) IClient {
	s.projectId = pprojectId
	if s.httpClient == nil {
		return s
	}

	// So it needs to reconfigure the gateway project id
	if s.vserverGateway != nil {
		s.vserverGateway = lsgateway.NewVServerGateway(s.vserverGateway.GetEndpoint(), s.projectId, s.httpClient)
	}

	if s.vlbGateway != nil {
		s.vlbGateway = lsgateway.NewVLBGateway(
			s.vlbGateway.GetEndpoint(),
			s.vserverGateway.GetEndpoint(),
			s.projectId,
			s.httpClient,
		)
	}

	if s.vnetworkGateway != nil {
		s.vnetworkGateway = lsgateway.NewVNetworkGateway(
			s.vnetworkGateway.GetEndpoint(),
			s.zoneId,
			s.projectId,
			s.userId,
			s.httpClient,
		)
	}

	return s
}

func (s *client) WithUserId(puserId string) IClient {
	s.userId = puserId
	if s.vnetworkGateway != nil {
		s.vnetworkGateway = lsgateway.NewVNetworkGateway(
			s.vnetworkGateway.GetEndpoint(),
			s.zoneId,
			s.projectId,
			s.userId,
			s.httpClient,
		)
	}

	return s
}

func (s *client) Configure(psdkCfg ISdkConfigure) IClient {
	s.projectId = psdkCfg.GetProjectId()
	s.userId = psdkCfg.GetUserId()
	if s.httpClient == nil {
		s.httpClient = lsclient.NewHttpClient(s.context)
	}

	if s.iamGateway == nil && psdkCfg.GetIamEndpoint() != "" {
		s.iamGateway = lsgateway.NewIamGateway(psdkCfg.GetIamEndpoint(), s.projectId, s.httpClient)
	}

	if s.vserverGateway == nil && psdkCfg.GetVServerEndpoint() != "" {
		s.vserverGateway = lsgateway.NewVServerGateway(
			psdkCfg.GetVServerEndpoint(),
			s.projectId,
			s.httpClient,
		)
	}

	if s.vlbGateway == nil && psdkCfg.GetVLBEndpoint() != "" && psdkCfg.GetVServerEndpoint() != "" {
		s.vlbGateway = lsgateway.NewVLBGateway(
			psdkCfg.GetVLBEndpoint(),
			psdkCfg.GetVServerEndpoint(),
			s.projectId,
			s.httpClient,
		)
	}

	if s.vnetworkGateway == nil && psdkCfg.GetVNetworkEndpoint() != "" {
		s.vnetworkGateway = lsgateway.NewVNetworkGateway(
			psdkCfg.GetVNetworkEndpoint(),
			psdkCfg.GetZoneId(),
			s.projectId,
			s.userId,
			s.httpClient,
		)
	}

	if s.glbGateway == nil && psdkCfg.GetGLBEndpoint() != "" {
		s.glbGateway = lsgateway.NewGLBGateway(psdkCfg.GetGLBEndpoint(), s.httpClient)
	}

	s.httpClient.WithReauthFunc(lsclient.IamOauth2, s.usingIamOauth2AsAuthOption(psdkCfg))
	s.userAgent = psdkCfg.GetUserAgent()

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

func (s *client) VBackUpGateway() lsgateway.IVBackUpGateway {
	return s.vbackupGateway
}

func (s *client) VNetworkGateway() lsgateway.IVNetworkGateway {
	return s.vnetworkGateway
}

func (s *client) GLBGateway() lsgateway.IGLBGateway {
	return s.glbGateway
}

func (s *client) usingIamOauth2AsAuthOption(pauthConfig ISdkConfigure) func() (lsclient.ISdkAuthentication, lserr.IError) {
	authFunc := func() (lsclient.ISdkAuthentication, lserr.IError) {
		token, err := s.iamGateway.V2().IdentityService().GetAccessToken(
			lssvcIdentityV2.NewGetAccessTokenRequest(pauthConfig.GetClientId(), pauthConfig.GetClientSecret()))
		if err != nil {
			return nil, err
		}

		return token.ToSdkAuthentication(), nil
	}

	return authFunc
}

func (s *client) GetUserAgent() string {
	return s.userAgent
}
