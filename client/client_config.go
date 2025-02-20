package client

import ljutils "github.com/cuongpiger/joat/utils"

type (
	ISdkConfigure interface {
		GetClientId() string
		GetClientSecret() string
		GetProjectId() string
		GetUserId() string
		GetZoneId() string
		GetIamEndpoint() string
		GetVServerEndpoint() string
		GetVLBEndpoint() string
		GetVNetworkEndpoint() string
		GetUserAgent() string
		WithUserId(puserId string) ISdkConfigure
		WithZoneId(pzoneId string) ISdkConfigure
		WithClientId(pclientId string) ISdkConfigure
		WithClientSecret(pclientSecret string) ISdkConfigure
		WithUserAgent(puserAgent string) ISdkConfigure
		WithProjectId(pprojectId string) ISdkConfigure
		WithIamEndpoint(piamEndpoint string) ISdkConfigure
		WithVServerEndpoint(pvserverEndpoint string) ISdkConfigure
		WithVLBEndpoint(pvlbEndpoint string) ISdkConfigure
		WithVNetworkEndpoint(pvnetworkEndpoint string) ISdkConfigure
		WithVLBGlobalEndpoint(pvlbEndpoint string) ISdkConfigure
	}
)

type sdkConfigure struct {
	clientId          string
	clientSecret      string
	projectId         string
	zoneId            string
	userId            string
	iamEndpoint       string
	vserverEndpoint   string
	vlbEndpoint       string
	vlbGlobalEndpoint string
	vnetworkEndpoint  string
	userAgent         string
}

func (s *sdkConfigure) GetClientId() string {
	return s.clientId
}

func (s *sdkConfigure) GetClientSecret() string {
	return s.clientSecret
}

func (s *sdkConfigure) GetProjectId() string {
	return s.projectId
}

func (s *sdkConfigure) GetUserId() string {
	return s.userId
}

func (s *sdkConfigure) GetZoneId() string {
	return s.zoneId
}

func (s *sdkConfigure) GetIamEndpoint() string {
	return s.iamEndpoint
}

func (s *sdkConfigure) GetVServerEndpoint() string {
	return s.vserverEndpoint
}

func (s *sdkConfigure) GetVLBEndpoint() string {
	return s.vlbEndpoint
}

func (s *sdkConfigure) GetVNetworkEndpoint() string {
	return s.vnetworkEndpoint
}

func (s *sdkConfigure) GetUserAgent() string {
	return s.userAgent
}

func (s *sdkConfigure) WithUserAgent(puserAgent string) ISdkConfigure {
	s.userAgent = puserAgent
	return s
}

func (s *sdkConfigure) WithClientId(pclientId string) ISdkConfigure {
	s.clientId = pclientId
	return s
}

func (s *sdkConfigure) WithClientSecret(pclientSecret string) ISdkConfigure {
	s.clientSecret = pclientSecret
	return s
}

func (s *sdkConfigure) WithUserId(puserId string) ISdkConfigure {
	s.userId = puserId
	return s
}

func (s *sdkConfigure) WithZoneId(pzoneId string) ISdkConfigure {
	s.zoneId = pzoneId
	return s
}

func (s *sdkConfigure) WithProjectId(pprojectId string) ISdkConfigure {
	s.projectId = pprojectId
	return s
}

func (s *sdkConfigure) WithIamEndpoint(piamEndpoint string) ISdkConfigure {
	s.iamEndpoint = ljutils.NormalizeURL(piamEndpoint)
	return s
}

func (s *sdkConfigure) WithVServerEndpoint(pvserverEndpoint string) ISdkConfigure {
	s.vserverEndpoint = ljutils.NormalizeURL(pvserverEndpoint)
	return s
}

func (s *sdkConfigure) WithVLBEndpoint(pvlbEndpoint string) ISdkConfigure {
	s.vlbEndpoint = ljutils.NormalizeURL(pvlbEndpoint)
	return s
}

func (s *sdkConfigure) WithVNetworkEndpoint(pvnetworkEndpoint string) ISdkConfigure {
	s.vnetworkEndpoint = ljutils.NormalizeURL(pvnetworkEndpoint)
	return s
}

func (s *sdkConfigure) WithVLBGlobalEndpoint(pvlbEndpoint string) ISdkConfigure {
	s.vlbGlobalEndpoint = ljutils.NormalizeURL(pvlbEndpoint)
	return s
}
