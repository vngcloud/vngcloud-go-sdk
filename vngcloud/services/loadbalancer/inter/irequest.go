package inter

import "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

type ICreateLoadBalancerRequest interface {
	ToRequestBody() interface{}
	AddUserAgent(pagent ...string) ICreateLoadBalancerRequest
	WithListener(plistener ICreateListenerRequest) ICreateLoadBalancerRequest
	WithPool(ppool ICreatePoolRequest) ICreateLoadBalancerRequest
	WithProjectId(pprojectId string) ICreateLoadBalancerRequest
	WithTags(ptags ...string) ICreateLoadBalancerRequest
	WithZoneId(zoneID common.Zone) ICreateLoadBalancerRequest
	GetMapHeaders() map[string]string
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type ICreateListenerRequest interface {
	ToRequestBody() interface{}
	WithAllowedCidrs(pcidrs ...string) ICreateListenerRequest
	WithLoadBalancerId(plbid string) ICreateListenerRequest
	WithDefaultPoolId(ppoolId string) ICreateListenerRequest
	WithTimeoutClient(ptoc int) ICreateListenerRequest
	WithTimeoutConnection(ptoc int) ICreateListenerRequest
	WithTimeoutMember(ptom int) ICreateListenerRequest
	AddCidrs(pcidrs ...string) ICreateListenerRequest
	ParseUserAgent() string
	GetLoadBalancerId() string
	ToMap() map[string]interface{}
}

type ICreatePoolRequest interface {
	ToRequestBody() interface{}
	WithHealthMonitor(pmonitor IHealthMonitorRequest) ICreatePoolRequest
	WithMembers(pmembers ...IMemberRequest) ICreatePoolRequest
	WithAlgorithm(palgorithm PoolAlgorithm) ICreatePoolRequest
}

type IHealthMonitorRequest interface {
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
	WithHealthyThreshold(pht int) IHealthMonitorRequest
	WithUnhealthyThreshold(puht int) IHealthMonitorRequest
	WithInterval(pinterval int) IHealthMonitorRequest
	WithTimeout(pto int) IHealthMonitorRequest
	WithHealthCheckMethod(pmethod HealthCheckMethod) IHealthMonitorRequest
	WithHttpVersion(pversion HealthCheckHttpVersion) IHealthMonitorRequest
	WithHealthCheckPath(ppath string) IHealthMonitorRequest
	WithSuccessCode(pcode string) IHealthMonitorRequest
	WithDomainName(pdomain string) IHealthMonitorRequest
}

type IMemberRequest interface {
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}
