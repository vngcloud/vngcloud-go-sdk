package policy

type IListOptsBuilder interface {
	GetProjectID() string
	GetLoadBalancerID() string
	GetListenerID() string
}

type ICreateOptsBuilder interface {
	GetProjectID() string
	GetLoadBalancerID() string
	GetListenerID() string
	ToRequestBody() interface{}
}

type IDeleteOptsBuilder interface {
	GetProjectID() string
	GetLoadBalancerID() string
	GetListenerID() string
	GetPolicyID() string
}

type IUpdateOptsBuilder interface {
	GetProjectID() string
	GetLoadBalancerID() string
	GetListenerID() string
	GetPolicyID() string
	ToRequestBody() interface{}
}

type IGetOptsBuilder interface {
	GetProjectID() string
	GetLoadBalancerID() string
	GetListenerID() string
	GetPolicyID() string
}