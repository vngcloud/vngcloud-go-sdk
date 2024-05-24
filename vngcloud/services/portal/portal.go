package portal

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsportalSvcV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/portal/v1"
	lsportalSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/portal/v2"
)

func NewPortalServiceV1(psvcClient lsclient.IServiceClient) IPortalServiceV1 {
	return &lsportalSvcV1.PortalServiceV1{
		PortalClient: psvcClient,
	}
}

func NewPortalServiceV2(psvcClient lsclient.IServiceClient) IPortalServiceV2 {
	return &lsportalSvcV2.PortalServiceV2{
		PortalClient: psvcClient,
	}
}
