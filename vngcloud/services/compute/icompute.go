package compute

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lscomputeSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/compute/v2"
)

type IComputeServiceV2 interface {
	CreateServer(popts lscomputeSvcV2.ICreateServerRequest) (*lsentity.Server, lserr.ISdkError)
	GetServerById(popts lscomputeSvcV2.IGetServerByIdRequest) (*lsentity.Server, lserr.ISdkError)
	DeleteServerById(popts lscomputeSvcV2.IDeleteServerByIdRequest) lserr.ISdkError
	UpdateServerSecgroupsByServerId(popts lscomputeSvcV2.IUpdateServerSecgroupsByServerIdRequest) (*lsentity.Server, lserr.ISdkError)
	AttachBlockVolume(popts lscomputeSvcV2.IAttachBlockVolumeRequest) lserr.ISdkError
	DetachBlockVolume(popts lscomputeSvcV2.IDetachBlockVolumeRequest) lserr.ISdkError
}
