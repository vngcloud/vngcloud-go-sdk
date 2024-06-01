package volume

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lsvolumeSvcV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/volume/v1"
	lsvolumeSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/volume/v2"
)

type IVolumeServiceV2 interface {
	CreateBlockVolume(popts lsvolumeSvcV2.ICreateBlockVolumeRequest) (*lsentity.Volume, lserr.ISdkError)
	DeleteBlockVolumeById(popts lsvolumeSvcV2.IDeleteBlockVolumeByIdRequest) lserr.ISdkError
	ListBlockVolumes(popts lsvolumeSvcV2.IListBlockVolumesRequest) (*lsentity.ListVolumes, lserr.ISdkError)
	GetBlockVolumeById(popts lsvolumeSvcV2.IGetBlockVolumeByIdRequest) (*lsentity.Volume, lserr.ISdkError)
	ResizeBlockVolumeById(popts lsvolumeSvcV2.IResizeBlockVolumeByIdRequest) (*lsentity.Volume, lserr.ISdkError)
	GetUnderBlockVolumeId(popts lsvolumeSvcV2.IGetUnderBlockVolumeIdRequest) (*lsentity.Volume, lserr.ISdkError)

	// Snapshot
	ListSnapshotsByBlockVolumeId(popts lsvolumeSvcV2.IListSnapshotsByBlockVolumeIdRequest) (*lsentity.ListSnapshots, lserr.ISdkError)
	CreateSnapshotByBlockVolumeId(popts lsvolumeSvcV2.ICreateSnapshotByBlockVolumeIdRequest) (*lsentity.Snapshot, lserr.ISdkError)
	DeleteSnapshotById(popts lsvolumeSvcV2.IDeleteSnapshotByIdRequest) lserr.ISdkError
}

type IVolumeServiceV1 interface {
	// Volume Type Api group
	GetVolumeTypeById(popts lsvolumeSvcV1.IGetVolumeTypeByIdRequest) (*lsentity.VolumeType, lserr.ISdkError)
}
