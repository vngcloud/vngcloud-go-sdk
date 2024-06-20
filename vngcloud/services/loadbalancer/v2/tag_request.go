package v2

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

func NewListTagsRequest(plbId string) IListTagsRequest {
	opt := new(ListTagsRequest)
	opt.LoadBalancerId = plbId
	return opt
}

func NewCreateTagsRequest(plbId string) ICreateTagsRequest {
	opts := new(CreateTagsRequest)
	opts.LoadBalancerId = plbId
	opts.ResourceID = plbId
	opts.ResourceType = "LOAD-BALANCER"
	opts.TagRequestList = make([]lscommon.Tag, 0)

	return opts
}

type ListTagsRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
}

type CreateTagsRequest struct {
	ResourceID     string         `json:"resourceId"`
	ResourceType   string         `json:"resourceType"`
	TagRequestList []lscommon.Tag `json:"tagRequestList"`

	lscommon.UserAgent
	lscommon.LoadBalancerCommon
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
