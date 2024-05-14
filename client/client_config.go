package client

type (
	ISdkAuthentication interface {
		WithAccessToken(paccessToken string) ISdkAuthentication
	}

	ISdkConfigure interface {
		GetClientId() string
		GetClientSecret() string
		GetIamEndpoint() string
		WithClientId(pclientId string) ISdkConfigure
		WithClientSecret(pclientSecret string) ISdkConfigure
		WithIamEndpoint(piamEndpoint string) ISdkConfigure
		WithVServerEndpoint(pvserverEndpoint string) ISdkConfigure
		WithVLBEndpoint(pvlbEndpoint string) ISdkConfigure
	}
)

type SdkAuthentication struct {
	accessToken string
}

func (s *SdkAuthentication) WithAccessToken(paccessToken string) ISdkAuthentication {
	s.accessToken = paccessToken
	return s
}

type authConfigure struct {
	clientId        string
	clientSecret    string
	iamEndpoint     string
	vserverEndpoint string
	vlbEndpoint     string
}

func (s *authConfigure) GetClientId() string {
	return s.clientId
}

func (s *authConfigure) GetClientSecret() string {
	return s.clientSecret
}

func (s *authConfigure) GetIamEndpoint() string {
	return s.iamEndpoint
}

func (s *authConfigure) WithClientId(pclientId string) ISdkConfigure {
	s.clientId = pclientId
	return s
}

func (s *authConfigure) WithClientSecret(pclientSecret string) ISdkConfigure {
	s.clientSecret = pclientSecret
	return s
}

func (s *authConfigure) WithIamEndpoint(piamEndpoint string) ISdkConfigure {
	s.iamEndpoint = piamEndpoint
	return s
}

func (s *authConfigure) WithVServerEndpoint(pvserverEndpoint string) ISdkConfigure {
	s.vserverEndpoint = pvserverEndpoint
	return s
}

func (s *authConfigure) WithVLBEndpoint(pvlbEndpoint string) ISdkConfigure {
	s.vlbEndpoint = pvlbEndpoint
	return s
}
