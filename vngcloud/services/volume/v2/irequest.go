package v2

type ICreateBlockVolumeRequest interface {
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
	GetListParameters() []interface{}
	GetSize() int64
	GetVolumeType() string
	GetVolumeName() string
	GetZone() string
	GetPoolName() string
	WithPoc(pisPoc bool) ICreateBlockVolumeRequest
	WithAutoRenew(pval bool) ICreateBlockVolumeRequest
	WithMultiAttach(pmultiAttach bool) ICreateBlockVolumeRequest
	WithSize(psize int64) ICreateBlockVolumeRequest
	WithEncryptionType(pet EncryptType) ICreateBlockVolumeRequest
	WithVolumeType(pvolumeTypeId string) ICreateBlockVolumeRequest
	WithZone(pzone string) ICreateBlockVolumeRequest
	WithPoolName(poolName string) ICreateBlockVolumeRequest
	WithVolumeRestoreFromSnapshot(psnapshotID, pvolumeTypeID string) ICreateBlockVolumeRequest
	WithTags(ptags ...string) ICreateBlockVolumeRequest
}

type IDeleteBlockVolumeByIdRequest interface {
	GetBlockVolumeId() string
}

type IListBlockVolumesRequest interface {
	WithName(pname string) IListBlockVolumesRequest
	ToQuery() (string, error)
	GetDefaultQuery() string
	ToMap() map[string]interface{}
}

type IGetBlockVolumeByIdRequest interface {
	GetBlockVolumeId() string
}

type IResizeBlockVolumeByIdRequest interface {
	ToRequestBody() interface{}
	GetBlockVolumeId() string
	GetSize() int
	GetVolumeTypeId() string
}

type IListSnapshotsByBlockVolumeIdRequest interface {
	GetBlockVolumeId() string
	ToQuery() (string, error)
	GetDefaultQuery() string
}

type ICreateSnapshotByBlockVolumeIdRequest interface {
	GetBlockVolumeId() string
	ToRequestBody() interface{}
	WithDescription(pdesc string) ICreateSnapshotByBlockVolumeIdRequest
	WithPermanently(pval bool) ICreateSnapshotByBlockVolumeIdRequest
	WithRetainedDay(pval uint64) ICreateSnapshotByBlockVolumeIdRequest
}

type IDeleteSnapshotByIdRequest interface {
	GetSnapshotId() string
	GetBlockVolumeId() string
}

type IGetUnderBlockVolumeIdRequest interface {
	GetBlockVolumeId() string
}

type IMigrateBlockVolumeByIdRequest interface {
	GetBlockVolumeId() string
	ToRequestBody() interface{}
	WithTags(ptags ...string) IMigrateBlockVolumeByIdRequest
	WithAction(paction MigrateAction) IMigrateBlockVolumeByIdRequest
	WithConfirm(pconfirm bool) IMigrateBlockVolumeByIdRequest
	IsConfirm() bool
}
