package gateway

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsportalSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/portal"
)

type vserverGatewayV1 struct {
	portalService lsportalSvc.IPortalServiceV1
}

func NewVServerGatewayV1(psvcClient lsclient.IServiceClient) IVServerGatewayV1 {
	return &vserverGatewayV1{
		portalService: lsportalSvc.NewPortalService(psvcClient),
	}
}

func (s *vserverGatewayV1) PortalService() lsportalSvc.IPortalServiceV1 {
	return s.portalService
}
