package gateway

import (
	lshttp "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/http"
	lssc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/service_client"
)

type iamGateway struct {
	iamGatewayV2 IIamGatewayV2
}

func NewIamGateway(pendpoint string, phc lshttp.IHttpClient) IIamGateway {
	svc := lssc.NewServiceClient().WithEndpoint(pendpoint).WithClient(phc)
	return &iamGateway{
		iamGatewayV2: NewIamGatewayV2(svc),
	}
}

func (s *iamGateway) V2() IIamGatewayV2 {
	return s.iamGatewayV2
}
