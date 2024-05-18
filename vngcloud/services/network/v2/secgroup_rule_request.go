package v2

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

type (
	CreateSecgroupRuleRequest struct {
		Description     string                `json:"description"`
		Direction       SecgroupRuleDirection `json:"direction"`
		EtherType       SecgroupRuleEtherType `json:"etherType"`
		PortRangeMax    int                   `json:"portRangeMax"`
		PortRangeMin    int                   `json:"portRangeMin"`
		Protocol        SecgroupRuleProtocol  `json:"protocol"`
		RemoteIPPrefix  string                `json:"remoteIpPrefix"`
		SecurityGroupID string                `json:"securityGroupId"`

		SecgroupCommon
	}

	SecgroupRuleDirection string
	SecgroupRuleEtherType string
	SecgroupRuleProtocol  string
)

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

func (s *CreateSecgroupRuleRequest) ToRequestBody() interface{} {
	return s
}
