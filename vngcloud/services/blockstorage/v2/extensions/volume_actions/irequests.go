package volume_actions

type IResizeOptsBuilder interface {
	GetProjectID() string
	GetVolumeID() string
	ToRequestBody() interface{}
}

type IMappingOptsBuilder interface {
	GetProjectID() string
	GetVolumeID() string
}
