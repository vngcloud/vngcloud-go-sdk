package pool

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
	lbCm "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/loadbalancer/v2"
)

// **************************************************** CreateOpts *****************************************************

type (
	CreateOptsAlgorithmOpt              string
	CreateOptsProtocolOpt               string
	CreateOptsHealthCheckProtocolOpt    string
	CreateOptsHealthCheckMethodOpt      string
	CreateOptsHealthCheckHttpVersionOpt string
)

const (
	CreateOptsAlgorithmOptRoundRobin CreateOptsAlgorithmOpt = "ROUND_ROBIN"
	CreateOptsAlgorithmOptLeastConn  CreateOptsAlgorithmOpt = "LEAST_CONNECTIONS"
	CreateOptsAlgorithmOptSourceIP   CreateOptsAlgorithmOpt = "SOURCE_IP"

	CreateOptsProtocolOptTCP   CreateOptsProtocolOpt = "TCP"
	CreateOptsProtocolOptUDP   CreateOptsProtocolOpt = "UDP"
	CreateOptsProtocolOptHTTP  CreateOptsProtocolOpt = "HTTP"
	CreateOptsProtocolOptProxy CreateOptsProtocolOpt = "PROXY"

	CreateOptsHealthCheckProtocolOptTCP     CreateOptsHealthCheckProtocolOpt = "TCP"
	CreateOptsHealthCheckProtocolOptHTTP    CreateOptsHealthCheckProtocolOpt = "HTTP"
	CreateOptsHealthCheckProtocolOptHTTPs   CreateOptsHealthCheckProtocolOpt = "HTTPS"
	CreateOptsHealthCheckProtocolOptPINGUDP CreateOptsHealthCheckProtocolOpt = "PING-UDP"

	CreateOptsHealthCheckMethodOptGET  CreateOptsHealthCheckMethodOpt = "GET"
	CreateOptsHealthCheckMethodOptPUT  CreateOptsHealthCheckMethodOpt = "PUT"
	CreateOptsHealthCheckMethodOptPOST CreateOptsHealthCheckMethodOpt = "POST"

	CreateOptsHealthCheckHttpVersionOptHttp1       CreateOptsHealthCheckHttpVersionOpt = "1.0"
	CreateOptsHealthCheckHttpVersionOptHttp1Minor1 CreateOptsHealthCheckHttpVersionOpt = "1.1"

	defaultFakeDomainName = "nip.io"
)

type (
	CreateOpts struct {
		Algorithm     CreateOptsAlgorithmOpt `json:"algorithm"`
		PoolName      string                 `json:"poolName"`
		PoolProtocol  CreateOptsProtocolOpt  `json:"poolProtocol"`
		Stickiness    *bool                  `json:"stickiness,omitempty"`    // only for l7, l4 doesn't have this field => nil
		TLSEncryption *bool                  `json:"tlsEncryption,omitempty"` // only for l7, l4 doesn't have this field => nil
		HealthMonitor HealthMonitor          `json:"healthMonitor"`
		Members       []*Member              `json:"members"`

		common.CommonOpts
		lbCm.LoadBalancerV2Common
	}

	UpdateOpts struct {
		Algorithm     CreateOptsAlgorithmOpt `json:"algorithm"`
		Stickiness    *bool                  `json:"stickiness,omitempty"`    // only for l7, l4 doesn't have this field => nil
		TLSEncryption *bool                  `json:"tlsEncryption,omitempty"` // only for l7, l4 doesn't have this field => nil
		HealthMonitor HealthMonitor          `json:"healthMonitor"`

		common.CommonOpts
		lbCm.LoadBalancerV2Common
		lbCm.PoolV2Common
	}

	Member struct {
		Backup      bool   `json:"backup"`
		IpAddress   string `json:"ipAddress"`
		MonitorPort int    `json:"monitorPort"`
		Name        string `json:"name"`
		Port        int    `json:"port"`
		Weight      int    `json:"weight"`
	}

	HealthMonitor struct {
		HealthCheckProtocol CreateOptsHealthCheckProtocolOpt     `json:"healthCheckProtocol"`
		HealthyThreshold    int                                  `json:"healthyThreshold"`
		UnhealthyThreshold  int                                  `json:"unhealthyThreshold"`
		Interval            int                                  `json:"interval"`
		Timeout             int                                  `json:"timeout"`
		HealthCheckMethod   *CreateOptsHealthCheckMethodOpt      `json:"healthCheckMethod,omitempty"`
		HttpVersion         *CreateOptsHealthCheckHttpVersionOpt `json:"httpVersion,omitempty"`
		HealthCheckPath     *string                              `json:"healthCheckPath,omitempty"`
		DomainName          *string                              `json:"domainName,omitempty"`
		SuccessCode         *string                              `json:"successCode,omitempty"`
	}
)

func (s *HealthMonitor) validate() {
	switch s.HealthCheckProtocol {
	case CreateOptsHealthCheckProtocolOptPINGUDP, CreateOptsHealthCheckProtocolOptTCP:
		s.HealthCheckPath = nil
		s.HttpVersion = nil
		s.SuccessCode = nil
		s.HealthCheckMethod = nil
		s.DomainName = nil

	case CreateOptsHealthCheckProtocolOptHTTP, CreateOptsHealthCheckProtocolOptHTTPs:
		if s.HttpVersion != nil {
			switch opt := *s.HttpVersion; opt {
			case CreateOptsHealthCheckHttpVersionOptHttp1:
				s.DomainName = nil
			case CreateOptsHealthCheckHttpVersionOptHttp1Minor1:
				if s.DomainName == nil ||
					(s.DomainName != nil && len(*s.DomainName) < 1) {

					fakeDomainName := defaultFakeDomainName
					s.DomainName = &fakeDomainName
				}
			}
		}
	}
}
func (s *CreateOpts) ToRequestBody() interface{} {
	s.HealthMonitor.validate()
	return s
}

// ****************************************** ListPoolsBasedLoadBalancerOpts *******************************************

type ListPoolsBasedLoadBalancerOpts struct {
	common.CommonOpts
	lbCm.LoadBalancerV2Common
}

// ***************************************************** DeleteOpts ****************************************************

type DeleteOpts struct {
	common.CommonOpts
	lbCm.LoadBalancerV2Common
	lbCm.PoolV2Common
}

// *********************************************** UpdatePoolMembersOpts ***********************************************

type UpdatePoolMembersOpts struct {
	Members []Member `json:"members"`

	common.CommonOpts
	lbCm.LoadBalancerV2Common
	lbCm.PoolV2Common
}

func (s *UpdatePoolMembersOpts) ToRequestBody() interface{} {
	return s
}

type GetOpts struct {
	common.CommonOpts
	lbCm.LoadBalancerV2Common
	lbCm.PoolV2Common
}

type GetMemberOpts struct {
	common.CommonOpts
	lbCm.LoadBalancerV2Common
	lbCm.PoolV2Common
}

func (s *UpdateOpts) ToRequestBody() interface{} {
	s.HealthMonitor.validate()
	return s
}
