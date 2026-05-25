package v1

type IListClustersRequest interface {
	ParseUserAgent() string
}

type IGetClusterRequest interface {
	GetClusterId() string
	ParseUserAgent() string
}

type IListEndpointsRequest interface {
	GetClusterId() string
	ParseUserAgent() string
}
