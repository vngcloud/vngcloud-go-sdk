package v2

type IGetSecgroupByIdRequest interface {
	GetSecgroupId() string
}

type ICreateSecgroupRequest interface {
	ToRequestBody() interface{}
	GetSecgroupName() string
}

type IDeleteSecgroupRequest interface {
	GetSecgroupId() string
}
