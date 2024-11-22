package v2

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

func NewListAllServersBySecgroupIdRequest(psecgroupId string) IListAllServersBySecgroupIdRequest {
	opt := new(ListAllServersBySecgroupIdRequest)
	opt.SecgroupId = psecgroupId
	return opt
}

type ListAllServersBySecgroupIdRequest struct {
	lscommon.SecgroupCommon
	lscommon.UserAgent
}

func (s *ListAllServersBySecgroupIdRequest) AddUserAgent(pagent ...string) IListAllServersBySecgroupIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}
