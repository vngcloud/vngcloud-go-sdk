package policy

type IListOptsBuilder interface {
	GetProjectID() string
	GetLoadBalancerID() string
	GetListenerID() string
}
