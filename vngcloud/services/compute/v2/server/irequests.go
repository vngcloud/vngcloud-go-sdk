package server

type IGetOptsBuilder interface {
	GetServerID() string
	GetProjectID() string
}
