package secgroup_rule

type ICreateOptsBuilder interface {
	GetProjectID() string
	ToRequestBody() interface{}
	GetSecgroupUUID() string
}

type IDeleteOptsBuilder interface {
	GetProjectID() string
	GetSecgroupUUID() string
	GetRuleUUID() string
}

type IListRulesBySecgroupIDOptsBuilder interface {
	GetProjectID() string
	GetSecgroupUUID() string
}
