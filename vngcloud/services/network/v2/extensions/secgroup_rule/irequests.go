package secgroup_rule

type ICreateOptsBuilder interface {
	GetProjectID() string
	ToRequestBody() interface{}
	GetSecgroupUUID() string
}
