package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type ICreateLoadBalancerRequest interface {
	ToRequestBody() interface{}
	AddUserAgent(pagent ...string) ICreateLoadBalancerRequest
	WithListener(plistener ICreateListenerRequest) ICreateLoadBalancerRequest
	WithPool(ppool ICreatePoolRequest) ICreateLoadBalancerRequest
	WithTags(ptags ...string) ICreateLoadBalancerRequest
	WithScheme(pscheme LoadBalancerScheme) ICreateLoadBalancerRequest
	WithAutoScalable(pautoScalable bool) ICreateLoadBalancerRequest
	WithType(ptype LoadBalancerType) ICreateLoadBalancerRequest
	ParseUserAgent() string
}

type IResizeLoadBalancerRequest interface {
	ToRequestBody() interface{}
	AddUserAgent(pagent ...string) IResizeLoadBalancerRequest
	WithPackageId(ppackageId string) IResizeLoadBalancerRequest
	ParseUserAgent() string

	GetLoadBalancerId() string
}

type IListLoadBalancerPackagesRequest interface {
	AddUserAgent(pagent ...string) IListLoadBalancerPackagesRequest
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
	WithTimeoutClient(ptoc int) ICreateListenerRequest
	WithTimeoutConnection(ptoc int) ICreateListenerRequest
	WithTimeoutMember(ptom int) ICreateListenerRequest
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

	WithCertificateAuthorities(pca ...string) IUpdateListenerRequest
	WithClientCertificate(pclientCert string) IUpdateListenerRequest
	WithDefaultCertificateAuthority(pdefaultCA string) IUpdateListenerRequest
}

type IGetPoolHealthMonitorByIdRequest interface {
	GetLoadBalancerId() string
	GetPoolId() string
	ParseUserAgent() string
}

type ICreatePoolRequest interface {
	ToRequestBody() interface{}
	WithHealthMonitor(pmonitor IHealthMonitorRequest) ICreatePoolRequest
	WithMembers(pmembers ...IMemberRequest) ICreatePoolRequest
	WithLoadBalancerId(plbId string) ICreatePoolRequest
	WithAlgorithm(palgorithm PoolAlgorithm) ICreatePoolRequest
	ToMap() map[string]interface{}
	GetLoadBalancerId() string
	ParseUserAgent() string
}

type IUpdatePoolRequest interface {
	GetPoolId() string
	ToRequestBody() interface{}
	WithHealthMonitor(pmonitor IHealthMonitorRequest) IUpdatePoolRequest
	WithLoadBalancerId(plbId string) IUpdatePoolRequest
	WithAlgorithm(palgorithm PoolAlgorithm) IUpdatePoolRequest
	WithStickiness(v *bool) IUpdatePoolRequest
	WithTLSEncryption(v *bool) IUpdatePoolRequest
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

type IListTagsRequest interface {
	GetLoadBalancerId() string
	ParseUserAgent() string
}

type ICreateTagsRequest interface {
	GetLoadBalancerId() string
	ToRequestBody() interface{}
	ParseUserAgent() string
	WithTags(ptags ...string) ICreateTagsRequest
}

type IUpdateTagsRequest interface {
	GetLoadBalancerId() string
	ToRequestBody(plstTags *lsentity.ListTags) interface{}
	ParseUserAgent() string
	WithTags(ptags ...string) IUpdateTagsRequest
	ToMap() map[string]interface{}
}

// --------------------------------------------------------

type IListPoliciesRequest interface {
	ParseUserAgent() string
	GetLoadBalancerId() string
	GetListenerId() string
}

type ICreatePolicyRequest interface {
	ToRequestBody() interface{}
	ParseUserAgent() string
	GetLoadBalancerId() string
	GetListenerId() string
	ToMap() map[string]interface{}

	WithName(pname string) ICreatePolicyRequest
	WithRules(prules ...L7RuleRequest) ICreatePolicyRequest
	WithAction(paction PolicyAction) ICreatePolicyRequest

	// only for action redirect to pool
	WithRedirectPoolId(predirectPoolId string) ICreatePolicyRequest

	// only for action redirect to url
	WithRedirectURL(predirectURL string) ICreatePolicyRequest
	// only for action redirect to url
	WithRedirectHTTPCode(predirectHTTPCode int) ICreatePolicyRequest
	// only for action redirect to url
	WithKeepQueryString(pkeepQueryString bool) ICreatePolicyRequest
}

type IGetPolicyByIdRequest interface {
	ParseUserAgent() string
	GetLoadBalancerId() string
	GetListenerId() string
	GetPolicyId() string
}

type IUpdatePolicyRequest interface {
	ToRequestBody() interface{}
	ParseUserAgent() string
	GetLoadBalancerId() string
	GetListenerId() string
	GetPolicyId() string

	WithAction(paction PolicyAction) IUpdatePolicyRequest
	WithRules(prules ...L7RuleRequest) IUpdatePolicyRequest
	WithRedirectPoolID(predirectPoolId string) IUpdatePolicyRequest
	WithRedirectURL(predirectURL string) IUpdatePolicyRequest
	WithRedirectHTTPCode(predirectHTTPCode int) IUpdatePolicyRequest
	WithKeepQueryString(pkeepQueryString bool) IUpdatePolicyRequest
}

type IDeletePolicyByIdRequest interface {
	ParseUserAgent() string
	GetLoadBalancerId() string
	GetListenerId() string
	GetPolicyId() string
}
