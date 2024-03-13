package server

type IGetOptsBuilder interface {
	GetServerID() string
	GetProjectID() string
}

type IDeleteOptsBuilder interface {
	GetServerID() string
	GetProjectID() string
	ToRequestBody() interface{}
}

type ICreateOptsBuilder interface {
	GetProjectID() string
	ToRequestBody() interface{}
}

type IListOptsBuilder interface {
	GetProjectID() string
	ToListQuery() (string, error)
}
