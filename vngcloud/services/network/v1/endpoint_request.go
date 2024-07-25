package v1

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

func NewGetEndpointByIdRequest(pendpointId string) IGetEndpointByIdRequest {
	opt := new(GetEndpointByIdRequest)
	opt.EndpointId = pendpointId
	return opt
}

type GetEndpointByIdRequest struct {
	lscommon.UserAgent
	lscommon.EndpointCommon
}

func (s *GetEndpointByIdRequest) AddUserAgent(pagent ...string) IGetEndpointByIdRequest {
	s.UserAgent.Agent = append(s.UserAgent.Agent, pagent...)
	return s
}
