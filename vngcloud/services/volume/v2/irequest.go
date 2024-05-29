package v2

type ICreateBlockVolumeRequest interface {
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
	WithPoc(pisPoc bool) ICreateBlockVolumeRequest
	WithAutoRenew(pval bool) ICreateBlockVolumeRequest
	WithMultiAttach(pmultiAttach bool) ICreateBlockVolumeRequest
	WithSize(psize int64) ICreateBlockVolumeRequest
	WithEncryptionType(pet EncryptType) ICreateBlockVolumeRequest
	WithVolumeType(pvolumeTypeId string) ICreateBlockVolumeRequest
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

type ICreateSnapshotByVolumeIdRequest interface {
	GetBlockVolumeId() string
	ToRequestBody() interface{}
	WithDescription(pdesc string) ICreateSnapshotByVolumeIdRequest
	WithPermanently(pval bool) ICreateSnapshotByVolumeIdRequest
	WithRetainedDay(pval uint64) ICreateSnapshotByVolumeIdRequest
}

type IDeleteSnapshotByIdRequest interface {
	GetSnapshotId() string
	GetBlockVolumeId() string
}
