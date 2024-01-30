package secgroup_rule

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
	lSecgroupCommonV2 "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/network/v2"
)

type (
	CreateOptsDirectionOpt string
	CreateOptsEtherTypeOpt string
	CreateOptsProtocolOpt  string
)

const (
	CreateOptsDirectionOptIngress CreateOptsDirectionOpt = "ingress"
	CreateOptsDirectionOptEgress  CreateOptsDirectionOpt = "egress"
)

const (
	CreateOptsEtherTypeOptIPv4 CreateOptsEtherTypeOpt = "IPv4"
	CreateOptsEtherTypeOptIPv6 CreateOptsEtherTypeOpt = "IPv6"
)

const (
	CreateOptsProtocolOptTCP CreateOptsProtocolOpt = "tcp"
	CreateOptsProtocolOptUDP CreateOptsProtocolOpt = "udp"
)

type CreateOpts struct {
	Description    string                 `json:"description"`
	Direction      CreateOptsDirectionOpt `json:"direction"`
	EtherType      CreateOptsEtherTypeOpt `json:"etherType"`
	PortRangeMax   int                    `json:"portRangeMax"`
	PortRangeMin   int                    `json:"portRangeMin"`
	Protocol       CreateOptsProtocolOpt  `json:"protocol"`
	RemoteIPPrefix string                 `json:"remoteIpPrefix"`
	SecgroupUUID   string                 `json:"securityGroupId"`

	common.CommonOpts
	lSecgroupCommonV2.SecgroupV2Common
}

func (s *CreateOpts) ToRequestBody() interface{} {
	return s
}
