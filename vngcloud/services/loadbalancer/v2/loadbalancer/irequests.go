package loadbalancer

import lsobj "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"

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
	ToListQuery() (string, error)
	GetDefaultQuery() string
}

type IUpdateOptsBuilder interface {
	GetProjectID() string
	GetLoadBalancerID() string
	ToRequestBody() interface{}
}

type ICreateTagOptsBuilder interface {
	GetProjectID() string
	GetLoadBalancerID() string
	ToRequestBody() interface{}
}

type IListTagsOptsBuilder interface {
	GetProjectID() string
	GetLoadBalancerID() string
}

type IUpdateTagOptsBuilder interface {
	GetProjectID() string
	GetLoadBalancerID() string
	ToRequestBody(ptags []*lsobj.LoadBalancerTag) interface{}
}
