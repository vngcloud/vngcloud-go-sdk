package v2

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
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

func NewCreateTagsRequest(pvolumeId string) ICreateTagsRequest {
	opts := new(CreateTagsRequest)
	opts.BlockVolumeId = pvolumeId
	opts.ResourceID = pvolumeId
	opts.ResourceType = volumeResourceType
	opts.TagRequestList = make([]lscommon.Tag, 0)
	return opts
}

type CreateTagsRequest struct {
	ResourceID     string         `json:"resourceId"`
	ResourceType   string         `json:"resourceType"`
	TagRequestList []lscommon.Tag `json:"tagRequestList"`

	lscommon.UserAgent
	lscommon.BlockVolumeCommon
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

func NewUpdateTagsRequest(pvolumeId string) IUpdateTagsRequest {
	opts := new(UpdateTagsRequest)
	opts.BlockVolumeId = pvolumeId
	opts.ResourceID = pvolumeId
	opts.ResourceType = volumeResourceType
	opts.TagRequestList = make([]lscommon.Tag, 0)
	return opts
}

type UpdateTagsRequest struct {
	ResourceID     string         `json:"resourceId"`
	ResourceType   string         `json:"resourceType"`
	TagRequestList []lscommon.Tag `json:"tagRequestList"`

	lscommon.UserAgent
	lscommon.BlockVolumeCommon
}

func (s *UpdateTagsRequest) AddUserAgent(pagent ...string) IUpdateTagsRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *UpdateTagsRequest) ToRequestBody(plstTags *lsentity.ListTags) interface{} {
	st := map[string]lscommon.Tag{}
	for _, tag := range plstTags.Items {
		st[tag.Key] = lscommon.Tag{
			IsEdited: false,
			Key:      tag.Key,
			Value:    tag.Value,
		}
	}

	for _, tag := range s.TagRequestList {
		st[tag.Key] = tag
	}

	s.TagRequestList = make([]lscommon.Tag, 0)
	for _, tag := range st {
		s.TagRequestList = append(s.TagRequestList, tag)
	}

	return s
}

func (s *UpdateTagsRequest) WithTags(ptags ...string) IUpdateTagsRequest {
	if len(ptags)%2 != 0 {
		ptags = append(ptags, "none")
	}

	for i := 0; i < len(ptags); i += 2 {
		s.TagRequestList = append(s.TagRequestList, lscommon.Tag{Key: ptags[i], Value: ptags[i+1]})
	}

	return s
}

func (s *UpdateTagsRequest) ToMap() map[string]interface{} {
	res := make(map[string]interface{})
	for _, tag := range s.TagRequestList {
		res[tag.Key] = tag.Value
	}
	return res
}
