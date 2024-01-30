package cluster

type IGetClusterOptsBuilder interface{}

type IGetOptsBuilder interface {
	GetClusterID() string
	GetProjectID() string
}

type IUpdateSecgroupOptsBuilder interface {
	GetClusterID() string
	GetProjectID() string
	ToRequestBody() interface{}
}
