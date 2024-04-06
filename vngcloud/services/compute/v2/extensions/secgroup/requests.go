package secgroup

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
	lcpV2Cm "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/compute/v2"
)

type UpdateSecgroupOpts struct {
	InstanceID    string   `json:"serverId"`
	SecurityGroup []string `json:"securityGroup"`

	common.CommonOpts
	lcpV2Cm.ServerV2Common
}

func (s *UpdateSecgroupOpts) ToRequestBody() interface{} {
	return s
}
