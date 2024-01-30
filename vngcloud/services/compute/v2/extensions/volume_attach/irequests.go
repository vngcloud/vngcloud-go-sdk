package volume_attach

type ICreateOptsBuilder interface {
	GetProjectID() string
	GetInstanceID() string
	GetVolumeID() string
	ToRequestBody() interface{}
}

type IDeleteOptsBuilder interface {
	GetProjectID() string
	GetInstanceID() string
	GetVolumeID() string
	ToRequestBody() interface{}
}
