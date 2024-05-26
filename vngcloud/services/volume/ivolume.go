package volume

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lsvolumeSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/volume/v2"
)

type IVolumeServiceV2 interface {
	CreateBlockVolume(popts lsvolumeSvcV2.ICreateBlockVolumeRequest) (*lsentity.Volume, lserr.ISdkError)
	DeleteBlockVolumeById(popts lsvolumeSvcV2.IDeleteBlockVolumeByIdRequest) lserr.ISdkError
	ListBlockVolumes(popts lsvolumeSvcV2.IListBlockVolumesRequest) (*lsentity.ListVolumes, lserr.ISdkError)
	GetBlockVolumeById(popts lsvolumeSvcV2.IGetBlockVolumeByIdRequest) (*lsentity.Volume, lserr.ISdkError)
}
