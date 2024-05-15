package portal

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsportalSvcV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/portal/v1"
)

func NewPortalService(psvcClient lsclient.IServiceClient) IPortalServiceV1 {
	return &lsportalSvcV1.PortalServiceV1{
		PortalClient: psvcClient,
	}
}
