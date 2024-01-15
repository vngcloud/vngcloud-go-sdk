package subnet

type IGetOptsBuilder interface {
	GetSubnetUUID() string
	GetProjectID() string
	GetNetworkUUID() string
}
