package v2

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

func NewGetAllAddressPairByVirtualSubnetIdRequest(psubnetId string) IGetAllAddressPairByVirtualSubnetIdRequest {
	opt := new(GetAllAddressPairByVirtualSubnetIdRequest)
	opt.VirtualSubnetId = psubnetId
	return opt
}

type GetAllAddressPairByVirtualSubnetIdRequest struct {
	lscommon.UserAgent
	VirtualSubnetId string
}

func (s *GetAllAddressPairByVirtualSubnetIdRequest) GetVirtualSubnetId() string {
	return s.VirtualSubnetId
}

func NewSetAddressPairInVirtualSubnetRequest(psubnetId, networkInterfaceID, CIDR string) ISetAddressPairInVirtualSubnetRequest {
	opt := new(SetAddressPairInVirtualSubnetRequest)
	opt.VirtualSubnetId = psubnetId
	opt.AddressPairRequest = AddressPairRequest{
		CIDR:                       CIDR,
		InternalNetworkInterfaceId: networkInterfaceID,
	}
	return opt
}

type SetAddressPairInVirtualSubnetRequest struct {
	lscommon.UserAgent
	VirtualSubnetId    string
	AddressPairRequest AddressPairRequest
}

func (s *SetAddressPairInVirtualSubnetRequest) GetVirtualSubnetId() string {
	return s.VirtualSubnetId
}

func (s *SetAddressPairInVirtualSubnetRequest) ToRequestBody() interface{} {
	return s.AddressPairRequest
}

type AddressPairRequest struct {
	CIDR                       string `json:"cidr"`
	InternalNetworkInterfaceId string `json:"internalNetworkInterfaceId"`
}
