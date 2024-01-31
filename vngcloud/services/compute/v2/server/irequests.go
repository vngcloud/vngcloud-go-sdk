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
