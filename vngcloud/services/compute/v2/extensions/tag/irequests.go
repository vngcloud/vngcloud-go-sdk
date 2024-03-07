package tag

type ICreateOptsBuilder interface {
	GetProjectID() string
	GetServerID() string
	ToRequestBody() interface{}
}

type IGetOptsBuilder interface {
	GetProjectID() string
	GetServerID() string
}
