package v2

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

func NewGetSubnetByIdRequest(pnetworkId, psubnetId string) IGetSubnetByIdRequest {
	opt := new(GetSubnetByIdRequest)
	opt.NetworkId = pnetworkId
	opt.SubnetId = psubnetId
	return opt
}

type GetSubnetByIdRequest struct {
	lscommon.UserAgent
	lscommon.SubnetCommon
	lscommon.NetworkCommon
}
