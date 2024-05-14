package gateway

import (
	lssc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/service_client"
	lsidentitySvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/identity"
)

type iamGatewayV2 struct {
	identityService lsidentitySvc.IIdentityServiceV2
}

func NewIamGatewayV2(psvcClient lssc.IServiceClient) IIamGatewayV2 {
	return &iamGatewayV2{
		identityService: lsidentitySvc.NewIdentityService(psvcClient),
	}
}

func (s *iamGatewayV2) IdentityService() lsidentitySvc.IIdentityServiceV2 {
	return s.identityService
}
