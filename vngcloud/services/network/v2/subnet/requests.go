package subnet

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
	lNetworkCommonV2 "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/network/v2"
)

// ****************************************************** GetOpts ******************************************************

type GetOpts struct {
	SubnetCommonOpts
	common.CommonOpts
	lNetworkCommonV2.NetworkV2Common
}

// ************************************************** ListByNetworkID **************************************************

type ListByNetworkIDOpts struct {
	common.CommonOpts
	lNetworkCommonV2.NetworkV2Common
}
