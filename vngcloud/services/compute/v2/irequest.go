package v2

type ICreateServerRequest interface {
	ToRequestBody() interface{}
	WithEncryptionVolume(pencryptionVolume bool) ICreateServerRequest
	WithAutoRenew(pval bool) ICreateServerRequest
	WithUserData(puserData string, pbase64Encode bool) ICreateServerRequest
	WithTags(ptags ...string) ICreateServerRequest
	WithAttachFloating(pattachFloating bool) ICreateServerRequest
	WithSecgroups(psecgroups ...string) ICreateServerRequest
	WithPoc(pisPoc bool) ICreateServerRequest
	WithType(ptype string) ICreateServerRequest
	WithProduct(pproduct string) ICreateServerRequest
}

type IGetServerByIdRequest interface {
	GetServerId() string
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
