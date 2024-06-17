package inter

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

func NewHealthMonitor(pcheckProtocol HealthCheckProtocol) IHealthMonitorRequest {
	switch pcheckProtocol {
	default:
		return newHealthMonitorTCPRequest()
	}
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

func newHealthMonitorTCPRequest() IHealthMonitorTCPRequest {
	return &HealthMonitor{
		HealthCheckProtocol: HealthCheckProtocolTCP,
		HealthyThreshold:    3,
		UnhealthyThreshold:  3,
		Interval:            30,
		Timeout:             5,
		HealthCheckPath:     nil,
		HttpVersion:         nil,
		SuccessCode:         nil,
		HealthCheckMethod:   nil,
		DomainName:          nil,
	}
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
}

type Member struct {
	Backup      bool   `json:"backup"`
	IpAddress   string `json:"ipAddress"`
	MonitorPort int    `json:"monitorPort"`
	Name        string `json:"name"`
	Port        int    `json:"port"`
	Weight      int    `json:"weight"`
}

func (s *CreatePoolRequest) ToRequestBody() interface{} {
	return s
}

func (s *HealthMonitor) validate() {
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
}

func (s *CreatePoolRequest) WithHealthMonitor(pmonitor IHealthMonitorRequest) ICreatePoolRequest {
	s.HealthMonitor = pmonitor
	return s
}

func (s *CreatePoolRequest) WithMembers(pmembers ...IMemberRequest) ICreatePoolRequest {
	s.Members = append(s.Members, pmembers...)
	return s
}

func (s *HealthMonitor) ToRequestBody() interface{} {
	return s
}

func (s *Member) ToRequestBody() interface{} {
	return s
}
