package gateway

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsopensearchSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/opensearch"
)

var _ IVDBOpenSearchGateway = &vdbOpenSearchGateway{}

type vdbOpenSearchGateway struct {
	endpoint               string
	vdbOpenSearchGatewayV1 IVDBOpenSearchGatewayV1
}

// NewVDBOpenSearchGateway wires the vDB OpenSearch service.
// Base URL example: https://vdb.console.vngcloud.vn — paths are layered under
// /vdb/open-search/v1/{projectId}/open-search/...
// puserId is the portal user id required as a request header (same as Kafka).
// pprojectId is required because OpenSearch paths embed the project id directly.
func NewVDBOpenSearchGateway(pendpoint, puserId, pprojectId string, phc lsclient.IHttpClient) IVDBOpenSearchGateway {
	openSearchSvc := lsclient.NewServiceClient().
		WithEndpoint(pendpoint + "vdb/open-search/v1").
		WithClient(phc).
		WithUserId(puserId).
		WithProjectId(pprojectId)

	return &vdbOpenSearchGateway{
		endpoint:               pendpoint,
		vdbOpenSearchGatewayV1: NewVDBOpenSearchGatewayV1(openSearchSvc),
	}
}

func (s *vdbOpenSearchGateway) V1() IVDBOpenSearchGatewayV1 {
	return s.vdbOpenSearchGatewayV1
}

func (s *vdbOpenSearchGateway) GetEndpoint() string {
	return s.endpoint
}

var _ IVDBOpenSearchGatewayV1 = &vdbOpenSearchGatewayV1{}

type vdbOpenSearchGatewayV1 struct {
	openSearchService lsopensearchSvc.IOpenSearchServiceV1
}

func NewVDBOpenSearchGatewayV1(psvcClient lsclient.IServiceClient) IVDBOpenSearchGatewayV1 {
	return &vdbOpenSearchGatewayV1{
		openSearchService: lsopensearchSvc.NewOpenSearchServiceV1(psvcClient),
	}
}

func (s *vdbOpenSearchGatewayV1) OpenSearchService() lsopensearchSvc.IOpenSearchServiceV1 {
	return s.openSearchService
}
