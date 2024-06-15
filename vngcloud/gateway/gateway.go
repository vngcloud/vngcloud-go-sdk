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

type vlbGateway struct {
	endpoint           string // Hold the endpoint of the vLB service
	vlbGatewayInternal IVLBGatewayInternal
	vlbGatewayV2       IVLBGatewayV2
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

func NewVLBGateway(pendpoint, pprojectId string, phc lsclient.IHttpClient) IVLBGateway {
	vlbSvcV2 := lsclient.NewServiceClient().
		WithEndpoint(pendpoint + "v2").
		WithClient(phc).
		WithProjectId(pprojectId)

	vlbSvcIn := lsclient.NewServiceClient().
		WithEndpoint(pendpoint + "internal").
		WithClient(phc).
		WithProjectId(pprojectId)

	return &vlbGateway{
		endpoint:           pendpoint,
		vlbGatewayV2:       NewVLBGatewayV2(vlbSvcV2),
		vlbGatewayInternal: NewVLBGatewayInternal(vlbSvcIn),
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

func (s *vlbGateway) Internal() IVLBGatewayInternal {
	return s.vlbGatewayInternal
}

func (s *vlbGateway) V2() IVLBGatewayV2 {
	return s.vlbGatewayV2
}

func (s *vserverGateway) GetEndpoint() string {
	return s.endpoint
}

func (s *vlbGateway) GetEndpoint() string {
	return s.endpoint
}
