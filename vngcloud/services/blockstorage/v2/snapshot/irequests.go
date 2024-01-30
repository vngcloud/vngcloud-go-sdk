package snapshot

type ICreateOptsBuilder interface {
	GetProjectID() string
	GetVolumeID() string
	ToRequestBody() interface{}
}

type IDeleteOptsBuilder interface {
	GetProjectID() string
	GetVolumeID() string
	GetSnapshotID() string
}
