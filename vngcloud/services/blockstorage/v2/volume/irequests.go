package volume

type IGetOptsBuilder interface {
	GetVolumeID() string
	GetProjectID() string
}

type IDeleteOptsBuilder interface {
	GetProjectID() string
	GetVolumeID() string
}

type ICreateOptsBuilder interface {
	ToRequestBody() interface{}
	GetProjectID() string
}

type IListOptsBuilder interface {
	GetProjectID() string
	ToListQuery() (string, error)
	GetDefaultQuery() string
}

type IListAllOptsBuilder interface {
	ToListQuery() (string, error)
	GetProjectID() string
}
