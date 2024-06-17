package inter

type ICreateLoadBalancerRequest interface {
	ToRequestBody() interface{}
	AddUserAgent(pagent ...string) ICreateLoadBalancerRequest
	WithListener(plistener ICreateListenerRequest) ICreateLoadBalancerRequest
	WithPool(ppool ICreatePoolRequest) ICreateLoadBalancerRequest
	WithProjectId(pprojectId string) ICreateLoadBalancerRequest
	GetMapHeaders() map[string]string
}

type ICreateListenerRequest interface {
	ToRequestBody() interface{}
	WithAllowedCidrs(pcidrs ...string) ICreateListenerRequest
	AddCidrs(pcidrs ...string) ICreateListenerRequest
}

type ICreatePoolRequest interface {
	ToRequestBody() interface{}
	WithHealthMonitor(pmonitor IHealthMonitorRequest) ICreatePoolRequest
	WithMembers(pmembers ...IMemberRequest) ICreatePoolRequest
}

type IHealthMonitorRequest interface {
	ToRequestBody() interface{}
}

type IHealthMonitorTCPRequest interface {
	ToRequestBody() interface{}
}

type IMemberRequest interface {
	ToRequestBody() interface{}
}
