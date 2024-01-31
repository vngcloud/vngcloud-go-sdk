package server

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
	lServerV2 "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/compute/v2"
)

// ************************************************** GetOptsBuilder ***************************************************

type GetOpts struct {
	common.CommonOpts
	lServerV2.ServerV2Common
}
