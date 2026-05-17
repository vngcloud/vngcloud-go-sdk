package v2

import (
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
)

const (
	volumeResourceType = "VOLUME"
)

func NewListTagsRequest(pvolumeId string) IListTagsRequest {
	o := new(ListTagsRequest)
	o.BlockVolumeId = pvolumeId
	return o
}

func (s *ListTagsRequest) AddUserAgent(pagent ...string) IListTagsRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

type ListTagsRequest struct {
	lscommon.UserAgent
	lscommon.BlockVolumeCommon
}
