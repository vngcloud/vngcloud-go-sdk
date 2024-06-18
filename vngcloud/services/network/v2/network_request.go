package v2

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

func NewGetNetworkByIdRequest(pnetworkId string) IGetNetworkByIdRequest {
	opt := new(GetNetworkByIdRequest)
	opt.NetworkId = pnetworkId
	return opt
}

type GetNetworkByIdRequest struct {
	lscommon.NetworkCommon
}
