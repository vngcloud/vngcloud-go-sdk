package tag

type ICreateOptsBuilder interface {
	GetProjectID() string
	GetResourceID() string
	ToRequestBody() interface{}
}

type IGetOptsBuilder interface {
	GetProjectID() string
	GetResourceID() string
}

type IUpdateOptsBuilder interface {
	GetProjectID() string
	GetResourceID() string
	ToRequestBody() interface{}
}
