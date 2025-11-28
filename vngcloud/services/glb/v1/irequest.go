package v1

type IListGlobalPoolsRequest interface {
	WithLoadBalancerId(plbId string) IListGlobalPoolsRequest
	GetLoadBalancerId() string

	AddUserAgent(pagent ...string) IListGlobalPoolsRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type ICreateGlobalPoolRequest interface {
	WithAlgorithm(palgorithm GlobalPoolAlgorithm) ICreateGlobalPoolRequest
	WithDescription(pdesc string) ICreateGlobalPoolRequest
	WithName(pname string) ICreateGlobalPoolRequest
	WithProtocol(pprotocol GlobalPoolProtocol) ICreateGlobalPoolRequest
	WithHealthMonitor(pmonitor IGlobalHealthMonitorRequest) ICreateGlobalPoolRequest
	WithMembers(pmembers ...ICreateGlobalPoolMemberRequest) ICreateGlobalPoolRequest

	WithLoadBalancerId(plbId string) ICreateGlobalPoolRequest
	GetLoadBalancerId() string // to use in request url

	AddUserAgent(pagent ...string) ICreateGlobalPoolRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

type IGlobalHealthMonitorRequest interface {
	WithHealthyThreshold(pht int) IGlobalHealthMonitorRequest
	WithInterval(pinterval int) IGlobalHealthMonitorRequest
	WithProtocol(pprotocol GlobalPoolHealthCheckProtocol) IGlobalHealthMonitorRequest
	WithTimeout(pto int) IGlobalHealthMonitorRequest
	WithUnhealthyThreshold(puht int) IGlobalHealthMonitorRequest

	// http, https
	WithHealthCheckMethod(pmethod *GlobalPoolHealthCheckMethod) IGlobalHealthMonitorRequest
	WithHttpVersion(pversion *GlobalPoolHealthCheckHttpVersion) IGlobalHealthMonitorRequest
	WithPath(ppath *string) IGlobalHealthMonitorRequest
	WithSuccessCode(pcode *string) IGlobalHealthMonitorRequest
	WithDomainName(pdomain *string) IGlobalHealthMonitorRequest

	AddUserAgent(pagent ...string) IGlobalHealthMonitorRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}
type ICreateGlobalPoolMemberRequest interface {
	WithName(pname string) ICreateGlobalPoolMemberRequest
	WithDescription(pdesc string) ICreateGlobalPoolMemberRequest
	WithRegion(pregion string) ICreateGlobalPoolMemberRequest
	WithVPCID(pvpcID string) ICreateGlobalPoolMemberRequest
	WithTrafficDial(pdial int) ICreateGlobalPoolMemberRequest
	WithMembers(pmembers ...IGlobalMemberRequest) ICreateGlobalPoolMemberRequest
	WithType(ptype GlobalPoolMemberType) ICreateGlobalPoolMemberRequest

	WithLoadBalancerId(plbId string) ICreateGlobalPoolMemberRequest
	WithPoolId(ppoolId string) ICreateGlobalPoolMemberRequest
	GetLoadBalancerId() string // to use in request url
	GetPoolId() string         // to use in request url

	AddUserAgent(pagent ...string) ICreateGlobalPoolMemberRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

type IGlobalMemberRequest interface {
	WithAddress(paddr string) IGlobalMemberRequest
	WithBackupRole(pbackup bool) IGlobalMemberRequest
	WithDescription(pdesc string) IGlobalMemberRequest
	WithMonitorPort(pport int) IGlobalMemberRequest
	WithName(pname string) IGlobalMemberRequest
	WithPort(pport int) IGlobalMemberRequest
	WithSubnetID(psubnetID string) IGlobalMemberRequest
	WithWeight(pweight int) IGlobalMemberRequest

	AddUserAgent(pagent ...string) IGlobalMemberRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

// --------------------------------------------------------

type IUpdateGlobalPoolRequest interface {
	WithAlgorithm(palgorithm GlobalPoolAlgorithm) IUpdateGlobalPoolRequest
	WithHealthMonitor(pmonitor IGlobalHealthMonitorRequest) IUpdateGlobalPoolRequest

	WithLoadBalancerId(plbId string) IUpdateGlobalPoolRequest
	WithPoolId(ppoolId string) IUpdateGlobalPoolRequest
	GetLoadBalancerId() string // to use in request url
	GetPoolId() string

	AddUserAgent(pagent ...string) IUpdateGlobalPoolRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

// --------------------------------------------------------

type IDeleteGlobalPoolRequest interface {
	WithLoadBalancerId(plbId string) IDeleteGlobalPoolRequest
	WithPoolId(ppoolId string) IDeleteGlobalPoolRequest
	GetLoadBalancerId() string // to use in request url
	GetPoolId() string

	AddUserAgent(pagent ...string) IDeleteGlobalPoolRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type IListGlobalPoolMembersRequest interface {
	WithLoadBalancerId(plbId string) IListGlobalPoolMembersRequest
	WithPoolId(ppoolId string) IListGlobalPoolMembersRequest
	GetLoadBalancerId() string // to use in request url
	GetPoolId() string

	AddUserAgent(pagent ...string) IListGlobalPoolMembersRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type IGetGlobalPoolMemberRequest interface {
	WithLoadBalancerId(plbId string) IGetGlobalPoolMemberRequest
	WithPoolId(ppoolId string) IGetGlobalPoolMemberRequest
	WithPoolMemberId(ppoolMemberId string) IGetGlobalPoolMemberRequest
	GetLoadBalancerId() string // to use in request url
	GetPoolId() string
	GetPoolMemberId() string

	AddUserAgent(pagent ...string) IGetGlobalPoolMemberRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type IDeleteGlobalPoolMemberRequest interface {
	WithLoadBalancerId(plbId string) IDeleteGlobalPoolMemberRequest
	WithPoolId(ppoolId string) IDeleteGlobalPoolMemberRequest
	WithPoolMemberId(ppoolMemberId string) IDeleteGlobalPoolMemberRequest
	GetLoadBalancerId() string // to use in request url
	GetPoolId() string
	GetPoolMemberId() string

	AddUserAgent(pagent ...string) IDeleteGlobalPoolMemberRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type IPatchGlobalPoolMembersRequest interface {
	WithBulkAction(paction ...IBulkActionRequest) IPatchGlobalPoolMembersRequest

	WithLoadBalancerId(plbId string) IPatchGlobalPoolMembersRequest
	WithPoolId(ppoolId string) IPatchGlobalPoolMembersRequest
	GetLoadBalancerId() string // to use in request url
	GetPoolId() string

	AddUserAgent(pagent ...string) IPatchGlobalPoolMembersRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

type IBulkActionRequest interface {
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

// --------------------------------------------------------

type IListGlobalListenersRequest interface {
	WithLoadBalancerId(plbId string) IListGlobalListenersRequest
	GetLoadBalancerId() string // to use in request url

	AddUserAgent(pagent ...string) IListGlobalListenersRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type IGetGlobalListenerRequest interface {
	WithLoadBalancerId(plbId string) IGetGlobalListenerRequest
	WithListenerId(plistenerId string) IGetGlobalListenerRequest
	GetLoadBalancerId() string // to use in request url
	GetListenerId() string

	AddUserAgent(pagent ...string) IGetGlobalListenerRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type ICreateGlobalListenerRequest interface {
	WithAllowedCidrs(pcidrs ...string) ICreateGlobalListenerRequest
	WithDescription(pdesc string) ICreateGlobalListenerRequest
	WithHeaders(pheaders ...string) ICreateGlobalListenerRequest
	WithName(pname string) ICreateGlobalListenerRequest
	WithPort(pport int) ICreateGlobalListenerRequest
	WithProtocol(pprotocol GlobalListenerProtocol) ICreateGlobalListenerRequest
	WithTimeoutClient(ptoc int) ICreateGlobalListenerRequest
	WithTimeoutConnection(ptoc int) ICreateGlobalListenerRequest
	WithTimeoutMember(ptom int) ICreateGlobalListenerRequest
	WithGlobalPoolId(ppoolId string) ICreateGlobalListenerRequest

	WithLoadBalancerId(plbid string) ICreateGlobalListenerRequest
	GetLoadBalancerId() string
	// AddCidrs(pcidrs ...string) ICreateGlobalListenerRequest

	AddUserAgent(pagent ...string) ICreateGlobalListenerRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

// --------------------------------------------------------

type IUpdateGlobalListenerRequest interface {
	WithAllowedCidrs(pcidrs ...string) IUpdateGlobalListenerRequest
	WithTimeoutClient(ptoc int) IUpdateGlobalListenerRequest
	WithTimeoutMember(ptom int) IUpdateGlobalListenerRequest
	WithTimeoutConnection(ptoc int) IUpdateGlobalListenerRequest
	WithHeaders(pheaders ...string) IUpdateGlobalListenerRequest
	WithGlobalPoolId(ppoolId string) IUpdateGlobalListenerRequest

	WithLoadBalancerId(plbId string) IUpdateGlobalListenerRequest
	WithListenerId(plistenerId string) IUpdateGlobalListenerRequest
	GetLoadBalancerId() string // to use in request url
	GetListenerId() string

	AddUserAgent(pagent ...string) IUpdateGlobalListenerRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

// --------------------------------------------------------

type IDeleteGlobalListenerRequest interface {
	WithLoadBalancerId(plbId string) IDeleteGlobalListenerRequest
	WithListenerId(plistenerId string) IDeleteGlobalListenerRequest
	GetLoadBalancerId() string // to use in request url
	GetListenerId() string

	AddUserAgent(pagent ...string) IDeleteGlobalListenerRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type IListGlobalPackagesRequest interface {
	AddUserAgent(pagent ...string) IListGlobalPackagesRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type IListGlobalRegionsRequest interface {
	AddUserAgent(pagent ...string) IListGlobalRegionsRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type IGetGlobalLoadBalancerUsageHistoriesRequest interface {
	WithLoadBalancerId(plbId string) IGetGlobalLoadBalancerUsageHistoriesRequest
	WithFrom(pfrom string) IGetGlobalLoadBalancerUsageHistoriesRequest
	WithTo(pto string) IGetGlobalLoadBalancerUsageHistoriesRequest
	WithType(ptype string) IGetGlobalLoadBalancerUsageHistoriesRequest
	GetLoadBalancerId() string

	AddUserAgent(pagent ...string) IGetGlobalLoadBalancerUsageHistoriesRequest
	ParseUserAgent() string
	ToListQuery() (string, error)
	GetDefaultQuery() string
}

// --------------------------------------------------------

type IListGlobalLoadBalancersRequest interface {
	WithName(pname string) IListGlobalLoadBalancersRequest
	WithTags(ptags ...string) IListGlobalLoadBalancersRequest
	ToListQuery() (string, error)
	GetDefaultQuery() string

	AddUserAgent(pagent ...string) IListGlobalLoadBalancersRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type ICreateGlobalLoadBalancerRequest interface {
	WithDescription(pdesc string) ICreateGlobalLoadBalancerRequest
	WithName(pname string) ICreateGlobalLoadBalancerRequest
	WithType(ptype GlobalLoadBalancerType) ICreateGlobalLoadBalancerRequest
	WithGlobalListener(plistener ICreateGlobalListenerRequest) ICreateGlobalLoadBalancerRequest
	WithGlobalPool(ppool ICreateGlobalPoolRequest) ICreateGlobalLoadBalancerRequest
	WithPackage(ppackageId string) ICreateGlobalLoadBalancerRequest
	WithPaymentFlow(ppaymentFlow GlobalLoadBalancerPaymentFlow) ICreateGlobalLoadBalancerRequest

	// WithTags(ptags ...string) ICreateGlobalLoadBalancerRequest
	// WithScheme(pscheme LoadBalancerScheme) ICreateGlobalLoadBalancerRequest
	// WithAutoScalable(pautoScalable bool) ICreateGlobalLoadBalancerRequest
	// WithPackageId(ppackageId string) ICreateGlobalLoadBalancerRequest
	// WithSubnetId(psubnetId string) ICreateGlobalLoadBalancerRequest
	// WithPoc(poc bool) ICreateGlobalLoadBalancerRequest

	AddUserAgent(pagent ...string) ICreateGlobalLoadBalancerRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

// --------------------------------------------------------

type IDeleteGlobalLoadBalancerRequest interface {
	WithLoadBalancerId(plbId string) IDeleteGlobalLoadBalancerRequest
	GetLoadBalancerId() string // to use in request url

	AddUserAgent(pagent ...string) IDeleteGlobalLoadBalancerRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type IGetGlobalLoadBalancerByIdRequest interface {
	WithLoadBalancerId(plbId string) IGetGlobalLoadBalancerByIdRequest
	GetLoadBalancerId() string // to use in request url

	AddUserAgent(pagent ...string) IGetGlobalLoadBalancerByIdRequest
	ParseUserAgent() string
}
