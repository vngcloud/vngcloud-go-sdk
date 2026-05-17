package v2

import (
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
)

const (
	serverResourceType = "SERVER"
)

func NewListTagsRequest(pserverId string) IListTagsRequest {
	o := new(ListTagsRequest)
	o.ServerId = pserverId
	return o
}

func (s *ListTagsRequest) AddUserAgent(pagent ...string) IListTagsRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

type ListTagsRequest struct {
	lscommon.UserAgent
	lscommon.ServerCommon
}
