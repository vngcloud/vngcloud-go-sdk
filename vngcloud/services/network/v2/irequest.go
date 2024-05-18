package v2

// Secgroup

type IGetSecgroupByIdRequest interface {
	GetSecgroupId() string
}

type ICreateSecgroupRequest interface {
	ToRequestBody() interface{}
	GetSecgroupName() string
}

type IDeleteSecgroupByIdRequest interface {
	GetSecgroupId() string
}

// Secgroup Rule
type ICreateSecgroupRuleRequest interface {
	GetSecgroupId() string
	ToRequestBody() interface{}
}

type IDeleteSecgroupRuleByIdRequest interface {
	GetSecgroupId() string
	GetSecgroupRuleId() string
}

type IListSecgroupRulesBySecgroupIdRequest interface {
	GetSecgroupId() string
}
