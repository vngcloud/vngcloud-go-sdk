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
	WithServerGroupId(pserverGroupId string) ICreateServerRequest
	WithPoc(pisPoc bool) ICreateServerRequest
	WithType(ptype string) ICreateServerRequest
	WithNetwork(pnetworkId, psubnetId string) ICreateServerRequest
	WithProduct(pproduct string) ICreateServerRequest
	WithServerNetworkInterface(pprojectId, pnetworkId, psubnetId string, pattachFloating bool) ICreateServerRequest
	WithZone(pzone string) ICreateServerRequest
	AddUserAgent(pagent ...string) ICreateServerRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IGetServerByIdRequest interface {
	GetServerId() string
	AddUserAgent(pagent ...string) IGetServerByIdRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IDeleteServerByIdRequest interface {
	GetServerId() string
	WithDeleteAllVolume(pok bool) IDeleteServerByIdRequest
	ToRequestBody() interface{}
	AddUserAgent(pagent ...string) IDeleteServerByIdRequest
	ParseUserAgent() string
}

type IUpdateServerSecgroupsByServerIdRequest interface {
	GetServerId() string
	ToRequestBody() interface{}
	GetListSecgroupsIds() []string
	AddUserAgent(pagent ...string) IUpdateServerSecgroupsByServerIdRequest
	ParseUserAgent() string
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
	ToMap() map[string]interface{}
}

type IDetachFloatingIpRequest interface {
	GetInternalNetworkInterfaceId() string
	GetWanId() string
	GetServerId() string
	ToRequestBody() interface{}
	AddUserAgent(pagent ...string) IDetachFloatingIpRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IListServerGroupPoliciesRequest interface {
	AddUserAgent(pagent ...string) IListServerGroupPoliciesRequest
	ParseUserAgent() string
}

type IDeleteServerGroupByIdRequest interface {
	AddUserAgent(pagent ...string) IDeleteServerGroupByIdRequest
	ParseUserAgent() string
	GetServerGroupId() string
	ToMap() map[string]interface{}
}

type IListServerGroupsRequest interface {
	WithName(pname string) IListServerGroupsRequest
	AddUserAgent(pagent ...string) IListServerGroupsRequest
	ToListQuery() (string, error)
	ParseUserAgent() string
	GetDefaultQuery() string
	ToMap() map[string]interface{}
}

type ICreateServerGroupRequest interface {
	ParseUserAgent() string
	AddUserAgent(pagent ...string) ICreateServerGroupRequest
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}
