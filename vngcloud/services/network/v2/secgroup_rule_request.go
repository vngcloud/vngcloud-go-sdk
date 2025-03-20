package v2

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

func NewCreateSecgroupRuleRequest(
	pdirection SecgroupRuleDirection,
	petherType SecgroupRuleEtherType,
	pprotocol SecgroupRuleProtocol,
	portMinRange, portMaxRange int,
	premoteIpPrefix, psecurityGroupId, pdescription string) ICreateSecgroupRuleRequest {

	opt := &CreateSecgroupRuleRequest{
		Description:     pdescription,
		Direction:       pdirection,
		EtherType:       petherType,
		PortRangeMax:    portMaxRange,
		PortRangeMin:    portMinRange,
		Protocol:        pprotocol,
		RemoteIPPrefix:  premoteIpPrefix,
		SecurityGroupID: psecurityGroupId,
	}
	opt.SecgroupId = psecurityGroupId
	return opt
}

func NewDeleteSecgroupRuleByIdRequest(psecgroupRuleId string) IDeleteSecgroupRuleByIdRequest {
	opt := new(DeleteSecgroupRuleByIdRequest)
	opt.SecgroupId = "undefined"
	opt.SecgroupRuleId = psecgroupRuleId
	return opt
}

func (s *DeleteSecgroupRuleByIdRequest) AddUserAgent(pagent ...string) IDeleteSecgroupRuleByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func NewListSecgroupRulesBySecgroupIdRequest(psecurityGroupId string) IListSecgroupRulesBySecgroupIdRequest {
	opt := new(ListSecgroupRulesBySecgroupIdRequest)
	opt.SecgroupId = psecurityGroupId
	return opt
}

const (
	SecgroupRuleProtocolTCP    SecgroupRuleProtocol = "tcp"
	SecgroupRuleProtocolUDP    SecgroupRuleProtocol = "udp"
	SecgroupRuleProtocolICMP   SecgroupRuleProtocol = "icmp"
	SecgroupRuleProtocolAll    SecgroupRuleProtocol = "any"
	SecgroupRuleProtocolIpInIp SecgroupRuleProtocol = "4"
)

const (
	SecgroupRuleEtherTypeIPv4 SecgroupRuleEtherType = "IPv4"
	SecgroupRuleEtherTypeIPv6 SecgroupRuleEtherType = "IPv6"
)

const (
	SecgroupRuleDirectionIngress SecgroupRuleDirection = "ingress"
	SecgroupRuleDirectionEgress  SecgroupRuleDirection = "egress"
)

type ( //_______________________________________________________________________________________________________________
	CreateSecgroupRuleRequest struct {
		Description     string                `json:"description"`
		Direction       SecgroupRuleDirection `json:"direction"`
		EtherType       SecgroupRuleEtherType `json:"etherType"`
		PortRangeMax    int                   `json:"portRangeMax"`
		PortRangeMin    int                   `json:"portRangeMin"`
		Protocol        SecgroupRuleProtocol  `json:"protocol"`
		RemoteIPPrefix  string                `json:"remoteIpPrefix"`
		SecurityGroupID string                `json:"securityGroupId"`

		lscommon.SecgroupCommon
		lscommon.UserAgent
	}

	SecgroupRuleDirection string
	SecgroupRuleEtherType string
	SecgroupRuleProtocol  string
)

func (s *CreateSecgroupRuleRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateSecgroupRuleRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"description":     s.Description,
		"direction":       s.Direction,
		"etherType":       s.EtherType,
		"portRangeMax":    s.PortRangeMax,
		"portRangeMin":    s.PortRangeMin,
		"protocol":        s.Protocol,
		"remoteIpPrefix":  s.RemoteIPPrefix,
		"securityGroupId": s.SecurityGroupID,
	}
}

func (s *CreateSecgroupRuleRequest) AddUserAgent(pagent ...string) ICreateSecgroupRuleRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

type DeleteSecgroupRuleByIdRequest struct { //__________________________________________________________________________
	SecgroupRuleId string

	lscommon.UserAgent
	lscommon.SecgroupCommon
}

func (s *DeleteSecgroupRuleByIdRequest) GetSecgroupRuleId() string {
	return s.SecgroupRuleId
}

type ListSecgroupRulesBySecgroupIdRequest struct { //___________________________________________________________________
	lscommon.SecgroupCommon
	lscommon.UserAgent
}

func (s *ListSecgroupRulesBySecgroupIdRequest) AddUserAgent(pagent ...string) IListSecgroupRulesBySecgroupIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}
