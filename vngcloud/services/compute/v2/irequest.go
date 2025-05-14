package v2

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

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
	WithZone(pzone lscommon.Zone) ICreateServerRequest
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

type IAttachBlockVolumeRequest interface {
	GetServerId() string
	GetBlockVolumeId() string
}

type IDetachBlockVolumeRequest interface {
	GetServerId() string
	GetBlockVolumeId() string
}
