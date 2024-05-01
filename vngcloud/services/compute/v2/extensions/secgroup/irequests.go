package secgroup

type IUpdateSecgroupOptsBuilder interface {
	GetProjectID() string
	GetServerID() string
	ToRequestBody() interface{}
}
