package v2

type ICreateBlockVolumeRequest interface {
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
	WithPoc(pisPoc bool) ICreateBlockVolumeRequest
	WithAutoRenew(pval bool) ICreateBlockVolumeRequest
	WithMultiAttach(pmultiAttach bool) ICreateBlockVolumeRequest
	WithSize(psize int64) ICreateBlockVolumeRequest
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
