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

func (s *GetSubnetByIdRequest) AddUserAgent(pagent ...string) IGetSubnetByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

// --------------------------------------------------------
type SecondarySubnetUpdateBody struct {
	Name string `json:"name"`
	CIDR string `json:"cidr"`
}
type UpdateSubnetBody struct {
	Name                    string                      `json:"name"`
	CIDR                    string                      `json:"cidr"`
	SecondarySubnetRequests []SecondarySubnetUpdateBody `json:"secondarySubnetRequests"`
}

func NewUpdateSubnetByIdRequest(pnetworkId, psubnetId string, updateBody *UpdateSubnetBody) IUpdateSubnetByIdRequest {
	opt := new(UpdateSubnetByIdRequest)
	opt.NetworkId = pnetworkId
	opt.SubnetId = psubnetId
	opt.UpdateSubnetBody = updateBody
	return opt
}

type UpdateSubnetByIdRequest struct {
	UpdateSubnetBody *UpdateSubnetBody `json:"subnet"`
	lscommon.UserAgent
	lscommon.SubnetCommon
	lscommon.NetworkCommon
}

func (s *UpdateSubnetByIdRequest) ToRequestBody() interface{} {
	return s.UpdateSubnetBody
}

func (s *UpdateSubnetByIdRequest) AddUserAgent(pagent ...string) IUpdateSubnetByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}
