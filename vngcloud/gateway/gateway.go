package gateway

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
)

type iamGateway struct {
	iamGatewayV2 IIamGatewayV2
}

func NewIamGateway(pendpoint string, phc lsclient.IHttpClient) IIamGateway {
	iamSvcV2 := lsclient.NewServiceClient().WithEndpoint(pendpoint + "v2").WithClient(phc)
	return &iamGateway{
		iamGatewayV2: NewIamGatewayV2(iamSvcV2),
	}
}

func (s *iamGateway) V2() IIamGatewayV2 {
	return s.iamGatewayV2
}
