package v1

type IListClustersRequest interface {
	ParseUserAgent() string
}

type IGetClusterRequest interface {
	GetClusterId() string
	ParseUserAgent() string
}

type IListUsersRequest interface {
	GetClusterId() string
	ParseUserAgent() string
}

type IGetUserAuthenCredsRequest interface {
	GetClusterId() string
	GetUserId() string
	ParseUserAgent() string
}

type IListTopicsRequest interface {
	GetClusterId() string
	ParseUserAgent() string
}

type ICreateTopicRequest interface {
	GetClusterId() string
	GetName() string
	ToRequestBody() interface{}
	WithName(pname string) ICreateTopicRequest
	WithPartitions(ppartitions int32) ICreateTopicRequest
	WithReplicas(preplicas int32) ICreateTopicRequest
	WithRetentionSeconds(pretentionSeconds int64) ICreateTopicRequest
	WithRetentionBytes(pretentionBytes int64) ICreateTopicRequest
	ParseUserAgent() string
}
