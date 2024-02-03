package network

type IGetOptsBuilder interface {
	GetProjectID() string
	GetNetworkUUID() string
}
