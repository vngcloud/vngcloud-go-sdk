package loadbalancer

type ICreateOptsBuilder interface {
	ToRequestBody() interface{}
	GetProjectID() string
}

type IGetOptsBuilder interface {
	GetLoadBalancerID() string
	GetProjectID() string
}

type IDeleteOptsBuilder interface {
	GetLoadBalancerID() string
	GetProjectID() string
}

type IListBySubnetIDOptsBuilder interface {
	GetSubnetID() string
	GetProjectID() string
}

type IListOptsBuilder interface {
	GetProjectID() string
}

type IUpdateOptsBuilder interface {
	GetProjectID() string
	GetLoadBalancerID() string
	ToRequestBody() interface{}
}
