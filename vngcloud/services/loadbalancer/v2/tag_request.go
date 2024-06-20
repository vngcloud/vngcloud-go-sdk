package v2

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

func NewListTagsRequest(plbId string) *ListTagsRequest {
	opt := new(ListTagsRequest)
	opt.LoadBalancerId = plbId
	return opt
}

type ListTagsRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
}
