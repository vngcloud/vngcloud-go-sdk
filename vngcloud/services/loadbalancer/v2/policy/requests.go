package policy

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
	lbCm "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/loadbalancer/v2"
)

type ListOptsBuilder struct {
	common.CommonOpts
	lbCm.LoadBalancerV2Common
	lbCm.ListenerV2Common
}
