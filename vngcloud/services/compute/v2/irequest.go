package v2

type ICreateServerRequest interface {
	ToRequestBody() interface{}
	WithRootDiskEncryptionType(pencryptionVolume DataDiskEncryptionType) ICreateServerRequest
	WithEncryptionVolume(penabled bool) ICreateServerRequest
	WithAutoRenew(pval bool) ICreateServerRequest
	WithUserData(puserData string, pbase64Encode bool) ICreateServerRequest
	WithTags(ptags ...string) ICreateServerRequest
	WithAttachFloating(pattachFloating bool) ICreateServerRequest
	WithSecgroups(psecgroups ...string) ICreateServerRequest
	WithPoc(pisPoc bool) ICreateServerRequest
	WithType(ptype string) ICreateServerRequest
	WithNetwork(pnetworkId, psubnetId string) ICreateServerRequest
	WithProduct(pproduct string) ICreateServerRequest
	WithServerNetworkInterface(pprojectId, pnetworkId, psubnetId string, pattachFloating bool) ICreateServerRequest
	AddUserAgent(pagent ...string) ICreateServerRequest
	ParseUserAgent() string
}

type IGetServerByIdRequest interface {
	GetServerId() string
	AddUserAgent(pagent ...string) IGetServerByIdRequest
	ParseUserAgent() string
}

type IDeleteServerByIdRequest interface {
	GetServerId() string
	WithDeleteAllVolume(pok bool) IDeleteServerByIdRequest
	ToRequestBody() interface{}
}

type IUpdateServerSecgroupsByServerIdRequest interface {
	GetServerId() string
	ToRequestBody() interface{}
	GetListSecgroupsIds() []string
}

type IAttachBlockVolumeRequest interface {
	GetServerId() string
	GetBlockVolumeId() string
}

type IDetachBlockVolumeRequest interface {
	GetServerId() string
	GetBlockVolumeId() string
}

type IAttachFloatingIpRequest interface {
	GetServerId() string
	GetInternalNetworkInterfaceId() string
	ToRequestBody() interface{}
	AddUserAgent(pagent ...string) IAttachFloatingIpRequest
	ParseUserAgent() string
}

type IDetachFloatingIpRequest interface {
	GetInternalNetworkInterfaceId() string
	GetWanId() string
	GetServerId() string
	ToRequestBody() interface{}
	AddUserAgent(pagent ...string) IDetachFloatingIpRequest
	ParseUserAgent() string
}
