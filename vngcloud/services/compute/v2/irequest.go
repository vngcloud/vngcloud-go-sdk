package v2

type ICreateServerRequest interface {
	ToRequestBody() interface{}
	WithEncryptionVolume(pencryptionVolume bool) ICreateServerRequest
	WithUserData(puserData string, pbase64Encode bool) ICreateServerRequest
	WithAttachFloating(pattachFloating bool) ICreateServerRequest
	WithSecgroups(psecgroups ...string) ICreateServerRequest
	WithPoct(pisPoc bool) ICreateServerRequest
	WithType(ptype string) ICreateServerRequest
	WithProduct(pproduct string) ICreateServerRequest
}