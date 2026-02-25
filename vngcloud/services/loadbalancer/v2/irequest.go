package v2

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	"github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
)

type ICreateLoadBalancerRequest interface {
	ToRequestBody() interface{}
	AddUserAgent(pagent ...string) ICreateLoadBalancerRequest
	WithListener(plistener ICreateListenerRequest) ICreateLoadBalancerRequest
	WithPool(ppool ICreatePoolRequest) ICreateLoadBalancerRequest
	WithTags(ptags ...string) ICreateLoadBalancerRequest
	WithScheme(pscheme LoadBalancerScheme) ICreateLoadBalancerRequest
	WithAutoScalable(pautoScalable bool) ICreateLoadBalancerRequest
	WithPackageId(ppackageId string) ICreateLoadBalancerRequest
	WithSubnetId(psubnetId string) ICreateLoadBalancerRequest
	WithType(ptype LoadBalancerType) ICreateLoadBalancerRequest
	WithPoc(poc bool) ICreateLoadBalancerRequest
	WithZoneId(pzoneId common.Zone) ICreateLoadBalancerRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IResizeLoadBalancerRequest interface {
	ToRequestBody() interface{}
	AddUserAgent(pagent ...string) IResizeLoadBalancerRequest
	WithPackageId(ppackageId string) IResizeLoadBalancerRequest
	ParseUserAgent() string

	GetLoadBalancerId() string
}

type IListLoadBalancerPackagesRequest interface {
	WithZoneId(pzoneId common.Zone) IListLoadBalancerPackagesRequest
	GetZoneId() string
	AddUserAgent(pagent ...string) IListLoadBalancerPackagesRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
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
	AddUserAgent(pagent ...string) IListLoadBalancersRequest
}

type ICreateListenerRequest interface {
	ToRequestBody() interface{}
	WithAllowedCidrs(pcidrs ...string) ICreateListenerRequest
	WithLoadBalancerId(plbid string) ICreateListenerRequest
	WithDefaultPoolId(ppoolId string) ICreateListenerRequest
	WithTimeoutClient(ptoc int) ICreateListenerRequest
	WithTimeoutConnection(ptoc int) ICreateListenerRequest
	WithTimeoutMember(ptom int) ICreateListenerRequest
	WithInsertHeaders(pheaders ...string) ICreateListenerRequest
	WithDefaultCertificateAuthority(pdefaultCA *string) ICreateListenerRequest
	WithCertificateAuthorities(pca *[]string) ICreateListenerRequest
	WithClientCertificate(pclientCert *string) ICreateListenerRequest
	AddCidrs(pcidrs ...string) ICreateListenerRequest
	ParseUserAgent() string
	GetLoadBalancerId() string
	ToMap() map[string]interface{}
	AddUserAgent(pagent ...string) ICreateListenerRequest
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
	WithInsertHeaders(pheaders ...string) IUpdateListenerRequest
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IUpdateListenerRequest

	WithCertificateAuthorities(pca ...string) IUpdateListenerRequest
	WithClientCertificate(pclientCert string) IUpdateListenerRequest
	WithDefaultCertificateAuthority(pdefaultCA string) IUpdateListenerRequest
}

type IGetPoolHealthMonitorByIdRequest interface {
	GetLoadBalancerId() string
	GetPoolId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IGetPoolHealthMonitorByIdRequest
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
	AddUserAgent(pagent ...string) ICreatePoolRequest
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
	AddUserAgent(pagent ...string) IUpdatePoolRequest
}

type IListListenersByLoadBalancerIdRequest interface {
	GetLoadBalancerId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IListListenersByLoadBalancerIdRequest
}

type IListPoolsByLoadBalancerIdRequest interface {
	GetLoadBalancerId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IListPoolsByLoadBalancerIdRequest
}

type IUpdatePoolMembersRequest interface {
	WithMembers(pmembers ...IMemberRequest) IUpdatePoolMembersRequest
	ToRequestBody() interface{}
	GetLoadBalancerId() string
	GetPoolId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IUpdatePoolMembersRequest
}

type IListPoolMembersRequest interface {
	GetLoadBalancerId() string
	GetPoolId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IListPoolMembersRequest
}

type IDeletePoolByIdRequest interface {
	GetLoadBalancerId() string
	GetPoolId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IDeletePoolByIdRequest
}

type IDeleteListenerByIdRequest interface {
	GetLoadBalancerId() string
	GetListenerId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IDeleteListenerByIdRequest
}

type IDeleteLoadBalancerByIdRequest interface {
	GetLoadBalancerId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IDeleteLoadBalancerByIdRequest
}

type IHealthMonitorRequest interface {
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
	WithHealthCheckProtocol(pprotocol HealthCheckProtocol) IHealthMonitorRequest
	WithHealthyThreshold(pht int) IHealthMonitorRequest
	WithUnhealthyThreshold(puht int) IHealthMonitorRequest
	WithInterval(pinterval int) IHealthMonitorRequest
	WithTimeout(pto int) IHealthMonitorRequest
	WithHealthCheckMethod(pmethod *HealthCheckMethod) IHealthMonitorRequest
	WithHttpVersion(pversion *HealthCheckHttpVersion) IHealthMonitorRequest
	WithHealthCheckPath(ppath *string) IHealthMonitorRequest
	WithSuccessCode(pcode *string) IHealthMonitorRequest
	WithDomainName(pdomain *string) IHealthMonitorRequest
	AddUserAgent(pagent ...string) IHealthMonitorRequest
}

type IMemberRequest interface {
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

type IListTagsRequest interface {
	GetLoadBalancerId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IListTagsRequest
}

type ICreateTagsRequest interface {
	GetLoadBalancerId() string
	ToRequestBody() interface{}
	ParseUserAgent() string
	WithTags(ptags ...string) ICreateTagsRequest
	AddUserAgent(pagent ...string) ICreateTagsRequest
}

type IUpdateTagsRequest interface {
	GetLoadBalancerId() string
	ToRequestBody(plstTags *lsentity.ListTags) interface{}
	ParseUserAgent() string
	WithTags(ptags ...string) IUpdateTagsRequest
	ToMap() map[string]interface{}
	AddUserAgent(pagent ...string) IUpdateTagsRequest
}

// --------------------------------------------------------

type IListPoliciesRequest interface {
	ParseUserAgent() string
	GetLoadBalancerId() string
	GetListenerId() string
	AddUserAgent(pagent ...string) IListPoliciesRequest
}

type ICreatePolicyRequest interface {
	ToRequestBody() interface{}
	ParseUserAgent() string
	GetLoadBalancerId() string
	GetListenerId() string
	ToMap() map[string]interface{}
	AddUserAgent(pagent ...string) ICreatePolicyRequest

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
	AddUserAgent(pagent ...string) IGetPolicyByIdRequest
}

type IUpdatePolicyRequest interface {
	ToRequestBody() interface{}
	ParseUserAgent() string
	GetLoadBalancerId() string
	GetListenerId() string
	GetPolicyId() string
	AddUserAgent(pagent ...string) IUpdatePolicyRequest

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
	AddUserAgent(pagent ...string) IDeletePolicyByIdRequest
}

type IReorderPoliciesRequest interface {
	ToRequestBody() interface{}
	ParseUserAgent() string
	GetLoadBalancerId() string
	GetListenerId() string
	AddUserAgent(pagent ...string) IReorderPoliciesRequest

	WithPoliciesOrder(plstPolicies []string) IReorderPoliciesRequest
}

type IGetPoolByIdRequest interface {
	GetLoadBalancerId() string
	GetPoolId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IGetPoolByIdRequest
}

type IGetListenerByIdRequest interface {
	GetLoadBalancerId() string
	GetListenerId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IGetListenerByIdRequest
}

type IResizeLoadBalancerByIdRequest interface {
	GetLoadBalancerId() string
	ToMap() map[string]interface{}
	ParseUserAgent() string
	ToRequestBody() interface{}
	AddUserAgent(pagent ...string) IResizeLoadBalancerByIdRequest
}

type IScaleLoadBalancerRequest interface {
	GetLoadBalancerId() string
	ToMap() map[string]interface{}
	ParseUserAgent() string
	ToRequestBody() interface{}
	AddUserAgent(pagent ...string) IScaleLoadBalancerRequest
	WithScaling(pscaling *ScalingConfig) IScaleLoadBalancerRequest
	WithNetworking(pnetworking *NetworkingConfig) IScaleLoadBalancerRequest
}

// --------------------------------------------------------

type IListCertificatesRequest interface {
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IListCertificatesRequest
}

type IGetCertificateByIdRequest interface {
	GetCertificateId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IGetCertificateByIdRequest
}

type ICreateCertificateRequest interface {
	ToRequestBody() interface{}
	ParseUserAgent() string
	ToMap() map[string]interface{}
	AddUserAgent(pagent ...string) ICreateCertificateRequest

	WithCertificateChain(pchain string) ICreateCertificateRequest
	WithPassphrase(ppassphrase string) ICreateCertificateRequest
	WithPrivateKey(pprivateKey string) ICreateCertificateRequest
}

type IDeleteCertificateByIdRequest interface {
	GetCertificateId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IDeleteCertificateByIdRequest
}
