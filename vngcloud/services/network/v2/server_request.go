package v2

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

func NewListAllServersBySecgroupIdRequest(psecgroupId string) IListAllServersBySecgroupIdRequest {
	opt := new(ListAllServersBySecgroupIdRequest)
	opt.SecgroupId = psecgroupId
	return opt
}

type ListAllServersBySecgroupIdRequest struct {
	lscommon.SecgroupCommon
}
