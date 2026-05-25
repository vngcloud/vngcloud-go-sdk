package v1

import (
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
)

// ListClustersRequest -----------------------------------------------------

type ListClustersRequest struct {
	lscommon.UserAgent
}

func NewListClustersRequest() IListClustersRequest {
	return &ListClustersRequest{}
}

// GetClusterRequest -------------------------------------------------------

type GetClusterRequest struct {
	ClusterId string
	lscommon.UserAgent
}

// NewGetClusterRequest takes the full UUID returned in the cluster's `id` field
// (NOT the truncated `clusterId` field — only `id` works as the URL segment).
func NewGetClusterRequest(pclusterId string) IGetClusterRequest {
	return &GetClusterRequest{ClusterId: pclusterId}
}

func (s *GetClusterRequest) GetClusterId() string {
	return s.ClusterId
}

// ListEndpointsRequest ----------------------------------------------------

type ListEndpointsRequest struct {
	ClusterId string
	lscommon.UserAgent
}

func NewListEndpointsRequest(pclusterId string) IListEndpointsRequest {
	return &ListEndpointsRequest{ClusterId: pclusterId}
}

func (s *ListEndpointsRequest) GetClusterId() string {
	return s.ClusterId
}
