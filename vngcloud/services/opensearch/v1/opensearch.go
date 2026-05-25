package v1

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

// ListClusters returns every OpenSearch cluster the authenticated portal user
// owns inside the configured project.
func (s *OpenSearchServiceV1) ListClusters(popts IListClustersRequest) (*[]lsentity.OpenSearchCluster, lserr.IError) {
	url := listClustersUrl(s.VopensearchClient)
	resp := new(ListClustersResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithUserId(s.VopensearchClient.GetUserId()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VopensearchClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.VopensearchClient.GetProjectId())
	}

	out := resp.ToEntity()
	return &out, nil
}

// ListEndpoints returns the public/private URLs of a cluster.
// Each cluster typically exposes two URLs:
//   - type="dashboard" :443   — OpenSearch Dashboards UI
//   - type="log"       :9200  — REST API (use this for bulk writes / _cluster/health probes)
func (s *OpenSearchServiceV1) ListEndpoints(popts IListEndpointsRequest) (*[]lsentity.OpenSearchEndpoint, lserr.IError) {
	url := listEndpointsUrl(s.VopensearchClient, popts)
	resp := new(ListEndpointsResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithUserId(s.VopensearchClient.GetUserId()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VopensearchClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"projectId", s.VopensearchClient.GetProjectId(),
				"clusterId", popts.GetClusterId(),
			)
	}

	out := resp.ToEntity()
	return &out, nil
}

// GetCluster fetches the detailed view of a single OpenSearch cluster.
// Pass the full UUID from the cluster's `id` field (not the truncated `clusterId`).
func (s *OpenSearchServiceV1) GetCluster(popts IGetClusterRequest) (*lsentity.OpenSearchCluster, lserr.IError) {
	url := getClusterUrl(s.VopensearchClient, popts)
	resp := new(GetClusterResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithUserId(s.VopensearchClient.GetUserId()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VopensearchClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"projectId", s.VopensearchClient.GetProjectId(),
				"clusterId", popts.GetClusterId(),
			)
	}

	return resp.Data.ToEntity(), nil
}
