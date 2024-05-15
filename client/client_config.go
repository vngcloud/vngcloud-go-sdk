package client

import ljutils "github.com/cuongpiger/joat/utils"

type (
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

type sdkConfigure struct {
	clientId        string
	clientSecret    string
	iamEndpoint     string
	vserverEndpoint string
	vlbEndpoint     string
}

func (s *sdkConfigure) GetClientId() string {
	return s.clientId
}

func (s *sdkConfigure) GetClientSecret() string {
	return s.clientSecret
}

func (s *sdkConfigure) GetIamEndpoint() string {
	return s.iamEndpoint
}

func (s *sdkConfigure) WithClientId(pclientId string) ISdkConfigure {
	s.clientId = pclientId
	return s
}

func (s *sdkConfigure) WithClientSecret(pclientSecret string) ISdkConfigure {
	s.clientSecret = pclientSecret
	return s
}

func (s *sdkConfigure) WithIamEndpoint(piamEndpoint string) ISdkConfigure {
	s.iamEndpoint = ljutils.NormalizeURL(piamEndpoint)
	return s
}

func (s *sdkConfigure) WithVServerEndpoint(pvserverEndpoint string) ISdkConfigure {
	s.vserverEndpoint = pvserverEndpoint
	return s
}

func (s *sdkConfigure) WithVLBEndpoint(pvlbEndpoint string) ISdkConfigure {
	s.vlbEndpoint = pvlbEndpoint
	return s
}
