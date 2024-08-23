package v2

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

func NewGetAllAddressPairsByVirtualSubnetIdRequest(psubnetId string) IGetAllAddressPairByVirtualSubnetIdRequest {
	opt := new(GetAllAddressPairsByVirtualSubnetIdRequest)
	opt.VirtualSubnetId = psubnetId
	return opt
}

type GetAllAddressPairsByVirtualSubnetIdRequest struct {
	lscommon.UserAgent
	VirtualSubnetId string
}

func (s *GetAllAddressPairsByVirtualSubnetIdRequest) GetVirtualSubnetId() string {
	return s.VirtualSubnetId
}
