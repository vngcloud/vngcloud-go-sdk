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

type vnetworkGateway struct {
	endpoint                  string
	vnetworkGatewayV1         IVNetworkGatewayV1
	vnetworkGatewayInternalV1 IVNetworkGatewayInternalV1
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

func NewVLBGateway(plbEndpoint, pserverEndpoint, pprojectId string, phc lsclient.IHttpClient) IVLBGateway {
	vlbSvcV2 := lsclient.NewServiceClient().
		WithEndpoint(plbEndpoint + "v2").
		WithClient(phc).
		WithProjectId(pprojectId)

	vlbSvcIn := lsclient.NewServiceClient().
		WithEndpoint(plbEndpoint + "internal").
		WithClient(phc).
		WithProjectId(pprojectId)

	vserverSvcV2 := lsclient.NewServiceClient().
		WithEndpoint(pserverEndpoint + "v2").
		WithClient(phc).
		WithProjectId(pprojectId)

	return &vlbGateway{
		endpoint:           plbEndpoint,
		vlbGatewayV2:       NewVLBGatewayV2(vlbSvcV2, vserverSvcV2),
		vlbGatewayInternal: NewVLBGatewayInternal(vlbSvcIn),
	}
}

func NewVNetworkGateway(pendpoint, pzoneId, projectId, puserId string, phc lsclient.IHttpClient) IVNetworkGateway {
	vnetworkSvcV1 := lsclient.NewServiceClient().
		WithEndpoint(pendpoint + "vnetwork/v1").
		WithClient(phc).
		WithZoneId(pzoneId).
		WithProjectId(projectId).
		WithUserId(puserId)

	vnetworkSvcInternalV1 := lsclient.NewServiceClient().
		WithEndpoint(pendpoint + "internal/v1").
		WithClient(phc).
		WithZoneId(pzoneId).
		WithProjectId(projectId).
		WithUserId(puserId)

	return &vnetworkGateway{
		endpoint:                  pendpoint,
		vnetworkGatewayV1:         NewVNetworkGatewayV1(vnetworkSvcV1),
		vnetworkGatewayInternalV1: NewVNetworkGatewayInternalV1(vnetworkSvcInternalV1),
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

func (s *vnetworkGateway) V1() IVNetworkGatewayV1 {
	return s.vnetworkGatewayV1
}

func (s *vnetworkGateway) GetEndpoint() string {
	return s.endpoint
}
