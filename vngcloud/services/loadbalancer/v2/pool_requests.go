package v2

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

const (
	PoolAlgorithmRoundRobin PoolAlgorithm = "ROUND_ROBIN"
	PoolAlgorithmLeastConn  PoolAlgorithm = "LEAST_CONNECTIONS"
	PoolAlgorithmSourceIP   PoolAlgorithm = "SOURCE_IP"
)

const (
	PoolProtocolTCP   PoolProtocol = "TCP"
	PoolProtocolUDP   PoolProtocol = "UDP"
	PoolProtocolHTTP  PoolProtocol = "HTTP"
	PoolProtocolProxy PoolProtocol = "PROXY"
)

const (
	HealthCheckProtocolTCP     HealthCheckProtocol = "TCP"
	HealthCheckProtocolHTTP    HealthCheckProtocol = "HTTP"
	HealthCheckProtocolHTTPs   HealthCheckProtocol = "HTTPS"
	HealthCheckProtocolPINGUDP HealthCheckProtocol = "PING-UDP"
)

const (
	HealthCheckMethodGET  HealthCheckMethod = "GET"
	HealthCheckMethodPUT  HealthCheckMethod = "PUT"
	HealthCheckMethodPOST HealthCheckMethod = "POST"
)

const (
	HealthCheckHttpVersionHttp1       HealthCheckHttpVersion = "1.0"
	HealthCheckHttpVersionHttp1Minor1 HealthCheckHttpVersion = "1.1"
)

const (
	defaultFakeDomainName = "nip.io"
)

func NewCreatePoolRequest(pname string, pprotocol PoolProtocol) ICreatePoolRequest {
	opts := new(CreatePoolRequest)
	opts.PoolName = pname
	opts.Algorithm = PoolAlgorithmRoundRobin
	opts.PoolProtocol = pprotocol
	opts.Members = make([]IMemberRequest, 0)

	return opts
}

func NewUpdatePoolRequest(lbID, poolID string) IUpdatePoolRequest {
	opts := new(UpdatePoolRequest)
	opts.LoadBalancerId = lbID
	opts.PoolId = poolID

	return opts
}

func NewGetPoolHealthMonitorByIdRequest(lbID, poolID string) IGetPoolHealthMonitorByIdRequest {
	opts := new(GetPoolHealthMonitorByIdRequest)
	opts.LoadBalancerId = lbID
	opts.PoolId = poolID

	return opts
}

func (s *GetPoolHealthMonitorByIdRequest) AddUserAgent(pagent ...string) IGetPoolHealthMonitorByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewListPoolsByLoadBalancerIdRequest(plbId string) IListPoolsByLoadBalancerIdRequest {
	opts := new(ListPoolsByLoadBalancerIdRequest)
	opts.LoadBalancerId = plbId

	return opts
}

func NewUpdatePoolMembersRequest(plbId, ppoolId string) IUpdatePoolMembersRequest {
	opts := new(UpdatePoolMembersRequest)
	opts.LoadBalancerId = plbId
	opts.PoolId = ppoolId
	opts.Members = make([]IMemberRequest, 0)

	return opts
}

func NewListPoolMembersRequest(plbId, ppoolId string) IListPoolMembersRequest {
	opts := new(ListPoolMembersRequest)
	opts.LoadBalancerId = plbId
	opts.PoolId = ppoolId

	return opts
}

func NewDeletePoolByIdRequest(plbId, ppoolId string) IDeletePoolByIdRequest {
	opts := new(DeletePoolByIdRequest)
	opts.LoadBalancerId = plbId
	opts.PoolId = ppoolId

	return opts
}

func NewHealthMonitor(pcheckProtocol HealthCheckProtocol) IHealthMonitorRequest {
	opts := new(HealthMonitor)
	opts.HealthCheckProtocol = pcheckProtocol
	opts.HealthyThreshold = 3
	opts.UnhealthyThreshold = 3
	opts.Interval = 30
	opts.Timeout = 5

	return opts
}

func NewMember(pname, pipAddress string, pport int, pmonitorPort int) IMemberRequest {
	return &Member{
		Backup:      false,
		IpAddress:   pipAddress,
		MonitorPort: pmonitorPort,
		Name:        pname,
		Port:        pport,
		Weight:      1,
	}
}

func NewGetPoolByIdRequest(plbId, ppoolId string) IGetPoolByIdRequest {
	opts := new(GetPoolByIdRequest)
	opts.LoadBalancerId = plbId
	opts.PoolId = ppoolId

	return opts
}

type (
	PoolAlgorithm          string
	PoolProtocol           string
	HealthCheckProtocol    string
	HealthCheckMethod      string
	HealthCheckHttpVersion string
)

type CreatePoolRequest struct {
	Algorithm     PoolAlgorithm         `json:"algorithm"`
	PoolName      string                `json:"poolName"`
	PoolProtocol  PoolProtocol          `json:"poolProtocol"`
	Stickiness    *bool                 `json:"stickiness,omitempty"`    // only for l7, l4 doesn't have this field => nil
	TLSEncryption *bool                 `json:"tlsEncryption,omitempty"` // only for l7, l4 doesn't have this field => nil
	HealthMonitor IHealthMonitorRequest `json:"healthMonitor"`
	Members       []IMemberRequest      `json:"members"`

	lscommon.LoadBalancerCommon
	lscommon.UserAgent
}

func (s *CreatePoolRequest) AddUserAgent(pagent ...string) ICreatePoolRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

type UpdatePoolRequest struct {
	Algorithm     PoolAlgorithm         `json:"algorithm"`
	Stickiness    *bool                 `json:"stickiness,omitempty"`    // only for l7, l4 doesn't have this field => nil
	TLSEncryption *bool                 `json:"tlsEncryption,omitempty"` // only for l7, l4 doesn't have this field => nil
	HealthMonitor IHealthMonitorRequest `json:"healthMonitor"`

	lscommon.LoadBalancerCommon
	lscommon.PoolCommon
	lscommon.UserAgent
}

type GetPoolHealthMonitorByIdRequest struct {
	lscommon.LoadBalancerCommon
	lscommon.PoolCommon
	lscommon.UserAgent
}

type ListPoolMembersRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
	lscommon.PoolCommon
}

func (s *ListPoolMembersRequest) AddUserAgent(pagent ...string) IListPoolMembersRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

type DeletePoolByIdRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
	lscommon.PoolCommon
}

func (s *DeletePoolByIdRequest) AddUserAgent(pagent ...string) IDeletePoolByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

type GetPoolByIdRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
	lscommon.PoolCommon
}

func (s *GetPoolByIdRequest) AddUserAgent(pagent ...string) IGetPoolByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

type HealthMonitor struct {
	HealthCheckProtocol HealthCheckProtocol     `json:"healthCheckProtocol"`
	HealthyThreshold    int                     `json:"healthyThreshold"`
	UnhealthyThreshold  int                     `json:"unhealthyThreshold"`
	Interval            int                     `json:"interval"`
	Timeout             int                     `json:"timeout"`
	HealthCheckMethod   *HealthCheckMethod      `json:"healthCheckMethod,omitempty"`
	HttpVersion         *HealthCheckHttpVersion `json:"httpVersion,omitempty"`
	HealthCheckPath     *string                 `json:"healthCheckPath,omitempty"`
	DomainName          *string                 `json:"domainName,omitempty"`
	SuccessCode         *string                 `json:"successCode,omitempty"`

	lscommon.UserAgent
}

type Member struct {
	Backup      bool   `json:"backup"`
	IpAddress   string `json:"ipAddress"`
	MonitorPort int    `json:"monitorPort"`
	Name        string `json:"name"`
	Port        int    `json:"port"`
	Weight      int    `json:"weight"`
}

type ListPoolsByLoadBalancerIdRequest struct {
	lscommon.LoadBalancerCommon
	lscommon.UserAgent
}

func (s *ListPoolsByLoadBalancerIdRequest) AddUserAgent(pagent ...string) IListPoolsByLoadBalancerIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

type UpdatePoolMembersRequest struct {
	Members []IMemberRequest `json:"members"`

	lscommon.UserAgent
	lscommon.LoadBalancerCommon
	lscommon.PoolCommon
}

func (s *UpdatePoolMembersRequest) AddUserAgent(pagent ...string) IUpdatePoolMembersRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *CreatePoolRequest) ToRequestBody() interface{} {
	s.HealthMonitor = s.HealthMonitor.(*HealthMonitor).toRequestBody()
	return s
}

func (s *HealthMonitor) toRequestBody() IHealthMonitorRequest {
	switch s.HealthCheckProtocol {
	case HealthCheckProtocolPINGUDP, HealthCheckProtocolTCP:
		s.HealthCheckPath = nil
		s.HttpVersion = nil
		s.SuccessCode = nil
		s.HealthCheckMethod = nil
		s.DomainName = nil

	case HealthCheckProtocolHTTP, HealthCheckProtocolHTTPs:
		if s.HttpVersion != nil {
			switch opt := *s.HttpVersion; opt {
			case HealthCheckHttpVersionHttp1:
				s.DomainName = nil
			case HealthCheckHttpVersionHttp1Minor1:
				if s.DomainName == nil ||
					(s.DomainName != nil && len(*s.DomainName) < 1) {

					fakeDomainName := defaultFakeDomainName
					s.DomainName = &fakeDomainName
				}
			}
		}
	}

	return s
}

func (s *HealthMonitor) AddUserAgent(pagent ...string) IHealthMonitorRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *HealthMonitor) WithHealthCheckProtocol(pprotocol HealthCheckProtocol) IHealthMonitorRequest {
	s.HealthCheckProtocol = pprotocol
	return s
}

func (s *CreatePoolRequest) WithHealthMonitor(pmonitor IHealthMonitorRequest) ICreatePoolRequest {
	s.HealthMonitor = pmonitor
	return s
}

func (s *CreatePoolRequest) WithMembers(pmembers ...IMemberRequest) ICreatePoolRequest {
	s.Members = append(s.Members, pmembers...)
	return s
}

func (s *CreatePoolRequest) WithLoadBalancerId(plbId string) ICreatePoolRequest {
	s.LoadBalancerId = plbId
	return s
}

func (s *CreatePoolRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"algorithm":     s.Algorithm,
		"poolName":      s.PoolName,
		"poolProtocol":  s.PoolProtocol,
		"stickiness":    s.Stickiness,
		"tlsEncryption": s.TLSEncryption,
		"healthMonitor": s.HealthMonitor.ToMap(),
		"members": func() []map[string]interface{} {
			members := make([]map[string]interface{}, 0, len(s.Members))
			for _, member := range s.Members {
				members = append(members, member.ToMap())
			}
			return members
		}(),
	}
}

func (s *CreatePoolRequest) WithAlgorithm(palgorithm PoolAlgorithm) ICreatePoolRequest {
	s.Algorithm = palgorithm
	return s
}

func (s *UpdatePoolRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"algorithm":     s.Algorithm,
		"stickiness":    s.Stickiness,
		"tlsEncryption": s.TLSEncryption,
		"healthMonitor": s.HealthMonitor.ToMap(),
	}
}

func (s *UpdatePoolRequest) ToRequestBody() interface{} {
	s.HealthMonitor = s.HealthMonitor.(*HealthMonitor).toRequestBody()
	return s
}

func (s *UpdatePoolRequest) WithAlgorithm(palgorithm PoolAlgorithm) IUpdatePoolRequest {
	s.Algorithm = palgorithm
	return s
}

func (s *UpdatePoolRequest) WithHealthMonitor(pmonitor IHealthMonitorRequest) IUpdatePoolRequest {
	s.HealthMonitor = pmonitor
	return s
}

func (s *UpdatePoolRequest) WithLoadBalancerId(plbId string) IUpdatePoolRequest {
	s.LoadBalancerId = plbId
	return s
}

func (s *UpdatePoolRequest) WithTLSEncryption(v *bool) IUpdatePoolRequest {
	s.TLSEncryption = v
	return s
}

func (s *UpdatePoolRequest) WithStickiness(v *bool) IUpdatePoolRequest {
	s.Stickiness = v
	return s
}

func (s *UpdatePoolRequest) AddUserAgent(pagent ...string) IUpdatePoolRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *HealthMonitor) ToRequestBody() interface{} {
	return s
}

func (s *HealthMonitor) WithHealthyThreshold(pht int) IHealthMonitorRequest {
	if pht < 1 {
		pht = 3
	}

	s.HealthyThreshold = pht
	return s
}

func (s *HealthMonitor) WithUnhealthyThreshold(puht int) IHealthMonitorRequest {
	if puht < 1 {
		puht = 3
	}

	s.UnhealthyThreshold = puht
	return s
}

func (s *HealthMonitor) WithInterval(pinterval int) IHealthMonitorRequest {
	if pinterval < 1 {
		pinterval = 30
	}

	s.Interval = pinterval
	return s
}

func (s *HealthMonitor) WithTimeout(pto int) IHealthMonitorRequest {
	if pto < 1 {
		pto = 5
	}

	s.Timeout = pto
	return s
}

func (s *HealthMonitor) WithHealthCheckMethod(pmethod *HealthCheckMethod) IHealthMonitorRequest {
	s.HealthCheckMethod = pmethod
	return s
}

func (s *HealthMonitor) WithHttpVersion(pversion *HealthCheckHttpVersion) IHealthMonitorRequest {
	s.HttpVersion = pversion
	return s
}

func (s *HealthMonitor) WithHealthCheckPath(ppath *string) IHealthMonitorRequest {
	s.HealthCheckPath = ppath
	return s
}

func (s *HealthMonitor) WithDomainName(pdomain *string) IHealthMonitorRequest {
	s.DomainName = pdomain
	return s
}

func (s *HealthMonitor) WithSuccessCode(pcode *string) IHealthMonitorRequest {
	s.SuccessCode = pcode
	return s
}

func (s *HealthMonitor) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"healthCheckProtocol": s.HealthCheckProtocol,
		"healthyThreshold":    s.HealthyThreshold,
		"unhealthyThreshold":  s.UnhealthyThreshold,
		"interval":            s.Interval,
		"timeout":             s.Timeout,
		"healthCheckMethod":   s.HealthCheckMethod,
		"httpVersion":         s.HttpVersion,
		"healthCheckPath":     s.HealthCheckPath,
		"domainName":          s.DomainName,
		"successCode":         s.SuccessCode,
	}
}

func (s *Member) ToRequestBody() interface{} {
	return s
}

func (s *Member) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"backup":      s.Backup,
		"ipAddress":   s.IpAddress,
		"monitorPort": s.MonitorPort,
		"name":        s.Name,
		"port":        s.Port,
		"weight":      s.Weight,
	}
}

func (s *UpdatePoolMembersRequest) WithMembers(pmembers ...IMemberRequest) IUpdatePoolMembersRequest {
	s.Members = append(s.Members, pmembers...)
	return s
}

func (s *UpdatePoolMembersRequest) ToRequestBody() interface{} {
	return s
}
