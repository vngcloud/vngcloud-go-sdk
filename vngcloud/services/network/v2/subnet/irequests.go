package subnet

type IGetOptsBuilder interface {
	GetSubnetUUID() string
	GetProjectID() string
	GetNetworkUUID() string
}

type IListByNetworkIDOptsBuilder interface {
	GetNetworkUUID() string
	GetProjectID() string
}
