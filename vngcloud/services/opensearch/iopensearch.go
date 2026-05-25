package opensearch

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lsopensearchSvcV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/opensearch/v1"
)

type IOpenSearchServiceV1 interface {
	ListClusters(popts lsopensearchSvcV1.IListClustersRequest) (*[]lsentity.OpenSearchCluster, lserr.IError)
	GetCluster(popts lsopensearchSvcV1.IGetClusterRequest) (*lsentity.OpenSearchCluster, lserr.IError)
	ListEndpoints(popts lsopensearchSvcV1.IListEndpointsRequest) (*[]lsentity.OpenSearchEndpoint, lserr.IError)
}
