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

func NewCreateTagsRequest(pserverId string) ICreateTagsRequest {
	opts := new(CreateTagsRequest)
	opts.ServerId = pserverId
	opts.ResourceID = pserverId
	opts.ResourceType = serverResourceType
	opts.TagRequestList = make([]lscommon.Tag, 0)
	return opts
}

type CreateTagsRequest struct {
	ResourceID     string         `json:"resourceId"`
	ResourceType   string         `json:"resourceType"`
	TagRequestList []lscommon.Tag `json:"tagRequestList"`

	lscommon.UserAgent
	lscommon.ServerCommon
}

func (s *CreateTagsRequest) AddUserAgent(pagent ...string) ICreateTagsRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *CreateTagsRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateTagsRequest) WithTags(ptags ...string) ICreateTagsRequest {
	if len(ptags)%2 != 0 {
		ptags = append(ptags, "none")
	}

	for i := 0; i < len(ptags); i += 2 {
		s.TagRequestList = append(s.TagRequestList, lscommon.Tag{Key: ptags[i], Value: ptags[i+1]})
	}
	return s
}
