package gateway

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
)

type iamGateway struct {
	iamGatewayV2 IIamGatewayV2
}

type vserverGateway struct {
	endpoint         string // Hold the endpoint of the vServer service
	vserverGatewayV1 IVServerGatewayV1
	vserverGatewayV2 IVServerGatewayV2
}

func NewIamGateway(pendpoint, projectId string, phc lsclient.IHttpClient) IIamGateway {
	iamSvcV2 := lsclient.NewServiceClient().
		WithEndpoint(pendpoint + "v2").
		WithClient(phc).
		WithProjectId(projectId)

	return &iamGateway{
		iamGatewayV2: NewIamGatewayV2(iamSvcV2),
	}
}

func NewVServerGateway(pendpoint, pprojectId string, phc lsclient.IHttpClient) IVServerGateway {
	vserverSvcV1 := lsclient.NewServiceClient().
		WithEndpoint(pendpoint + "v1").
		WithClient(phc).
		WithProjectId(pprojectId)

	vserverSvcV2 := lsclient.NewServiceClient().
		WithEndpoint(pendpoint + "v2").
		WithClient(phc).
		WithProjectId(pprojectId)

	return &vserverGateway{
		endpoint:         pendpoint,
		vserverGatewayV1: NewVServerGatewayV1(vserverSvcV1),
		vserverGatewayV2: NewVServerGatewayV2(vserverSvcV2),
	}
}

func (s *iamGateway) V2() IIamGatewayV2 {
	return s.iamGatewayV2
}

func (s *vserverGateway) V1() IVServerGatewayV1 {
	return s.vserverGatewayV1
}

func (s *vserverGateway) V2() IVServerGatewayV2 {
	return s.vserverGatewayV2
}

func (s *vserverGateway) GetEndpoint() string {
	return s.endpoint
}
