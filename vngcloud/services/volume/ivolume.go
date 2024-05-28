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
	ResizeBlockVolumeById(popts lsvolumeSvcV2.IResizeBlockVolumeByIdRequest) (*lsentity.Volume, lserr.ISdkError)

	// Snapshot
	ListSnapshotByVolumeId(popts lsvolumeSvcV2.IListSnapshotsByBlockVolumeIdRequest) (*lsentity.ListSnapshots, lserr.ISdkError)
	CreateSnapshotByVolumeId(popts lsvolumeSvcV2.ICreateSnapshotByVolumeIdRequest) (*lsentity.Snapshot, lserr.ISdkError)
	DeleteSnapshotById(popts lsvolumeSvcV2.IDeleteSnapshotByIdRequest) lserr.ISdkError
}
