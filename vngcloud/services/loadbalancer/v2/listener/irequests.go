package listener

type ICreateOptsBuilder interface {
	GetLoadBalancerID() string
	GetProjectID() string
	ToRequestBody() interface{}
}

type IGetBasedLoadBalancerOptsBuilder interface {
	GetLoadBalancerID() string
	GetProjectID() string
}

type IDeleteOptsBuilder interface {
	GetLoadBalancerID() string
	GetProjectID() string
	GetListenerID() string
}

type IUpdateOptsBuilder interface {
	GetLoadBalancerID() string
	GetProjectID() string
	GetListenerID() string
	ToRequestBody() interface{}
}
