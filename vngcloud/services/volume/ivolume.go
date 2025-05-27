package volume

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lsvolumeSvcV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/volume/v1"
	lsvolumeSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/volume/v2"
)

type IVolumeServiceV2 interface {
	CreateBlockVolume(popts lsvolumeSvcV2.ICreateBlockVolumeRequest) (*lsentity.Volume, lserr.IError)
	DeleteBlockVolumeById(popts lsvolumeSvcV2.IDeleteBlockVolumeByIdRequest) lserr.IError
	ListBlockVolumes(popts lsvolumeSvcV2.IListBlockVolumesRequest) (*lsentity.ListVolumes, lserr.IError)
	GetBlockVolumeById(popts lsvolumeSvcV2.IGetBlockVolumeByIdRequest) (*lsentity.Volume, lserr.IError)
	ResizeBlockVolumeById(popts lsvolumeSvcV2.IResizeBlockVolumeByIdRequest) (*lsentity.Volume, lserr.IError)
	GetUnderBlockVolumeId(popts lsvolumeSvcV2.IGetUnderBlockVolumeIdRequest) (*lsentity.Volume, lserr.IError)
	MigrateBlockVolumeById(popts lsvolumeSvcV2.IMigrateBlockVolumeByIdRequest) lserr.IError

	// Snapshot
	ListSnapshotsByBlockVolumeId(popts lsvolumeSvcV2.IListSnapshotsByBlockVolumeIdRequest) (*lsentity.ListSnapshots, lserr.IError)
	CreateSnapshotByBlockVolumeId(popts lsvolumeSvcV2.ICreateSnapshotByBlockVolumeIdRequest) (*lsentity.Snapshot, lserr.IError)
	DeleteSnapshotById(popts lsvolumeSvcV2.IDeleteSnapshotByIdRequest) lserr.IError
}

type IVolumeServiceV1 interface {
	// Volume Type Api group
	GetVolumeTypeById(popts lsvolumeSvcV1.IGetVolumeTypeByIdRequest) (*lsentity.VolumeType, lserr.IError)
	GetListVolumeTypes(popts lsvolumeSvcV1.IGetListVolumeTypeRequest) (*lsentity.ListVolumeType, lserr.IError)
	GetVolumeTypeZones(popts lsvolumeSvcV1.IGetVolumeTypeZonesRequest) (*lsentity.ListVolumeTypeZones, lserr.IError)
	GetDefaultVolumeType() (*lsentity.VolumeType, lserr.IError)
}
