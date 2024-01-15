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
	ToListQuery() (string, error)
	ToListQueryWithParams(*map[string]interface{}) (string, error)
	GetProjectID() string
}

type IListAllOptsBuilder interface {
	ToListQuery() (string, error)
	GetProjectID() string
}
