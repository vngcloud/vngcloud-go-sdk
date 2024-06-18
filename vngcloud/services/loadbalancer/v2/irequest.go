package v2

type ICreateLoadBalancerRequest interface {
	ToRequestBody() interface{}
	AddUserAgent(pagent ...string) ICreateLoadBalancerRequest
	WithListener(plistener ICreateListenerRequest) ICreateLoadBalancerRequest
	WithPool(ppool ICreatePoolRequest) ICreateLoadBalancerRequest
	WithTags(ptags ...string) ICreateLoadBalancerRequest
	ParseUserAgent() string
}

type IGetLoadBalancerByIdRequest interface {
	AddUserAgent(pagent ...string) IGetLoadBalancerByIdRequest
	ParseUserAgent() string
	GetLoadBalancerId() string
}

type IListLoadBalancersRequest interface {
	WithName(pname string) IListLoadBalancersRequest
	WithTags(ptags ...string) IListLoadBalancersRequest
	ToListQuery() (string, error)
	ParseUserAgent() string
	GetDefaultQuery() string
}

type ICreateListenerRequest interface {
	ToRequestBody() interface{}
	WithAllowedCidrs(pcidrs ...string) ICreateListenerRequest
	WithLoadBalancerId(plbid string) ICreateListenerRequest
	WithDefaultPoolId(ppoolId string) ICreateListenerRequest
	AddCidrs(pcidrs ...string) ICreateListenerRequest
	ParseUserAgent() string
	GetLoadBalancerId() string
	ToMap() map[string]interface{}
}

type IUpdateListenerRequest interface {
	GetLoadBalancerId() string
	GetListenerId() string
	ToRequestBody() interface{}
	WithCidrs(pcidrs ...string) IUpdateListenerRequest
	WithTimeoutClient(ptoc int) IUpdateListenerRequest
	WithTimeoutConnection(ptoc int) IUpdateListenerRequest
	WithTimeoutMember(ptom int) IUpdateListenerRequest
	WithDefaultPoolId(ppoolId string) IUpdateListenerRequest
	WithHeaders(pheaders ...string) IUpdateListenerRequest
	ParseUserAgent() string
}

type ICreatePoolRequest interface {
	ToRequestBody() interface{}
	WithHealthMonitor(pmonitor IHealthMonitorRequest) ICreatePoolRequest
	WithMembers(pmembers ...IMemberRequest) ICreatePoolRequest
	WithLoadBalancerId(plbId string) ICreatePoolRequest
	ToMap() map[string]interface{}
	GetLoadBalancerId() string
	ParseUserAgent() string
}

type IListListenersByLoadBalancerIdRequest interface {
	GetLoadBalancerId() string
	ParseUserAgent() string
}

type IListPoolsByLoadBalancerIdRequest interface {
	GetLoadBalancerId() string
	ParseUserAgent() string
}

type IUpdatePoolMembersRequest interface {
	WithMembers(pmembers ...IMemberRequest) IUpdatePoolMembersRequest
	ToRequestBody() interface{}
	GetLoadBalancerId() string
	GetPoolId() string
	ParseUserAgent() string
}

type IListPoolMembersRequest interface {
	GetLoadBalancerId() string
	GetPoolId() string
	ParseUserAgent() string
}

type IDeletePoolByIdRequest interface {
	GetLoadBalancerId() string
	GetPoolId() string
	ParseUserAgent() string
}

type IDeleteListenerByIdRequest interface {
	GetLoadBalancerId() string
	GetListenerId() string
	ParseUserAgent() string
}

type IDeleteLoadBalancerByIdRequest interface {
	GetLoadBalancerId() string
	ParseUserAgent() string
}

type IHealthMonitorRequest interface {
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

type IHealthMonitorTCPRequest interface {
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

type IMemberRequest interface {
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}
