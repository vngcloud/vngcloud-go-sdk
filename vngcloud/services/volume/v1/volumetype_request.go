package v1

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

func NewGetVolumeTypeByIdRequest(pvolumeTypeId string) IGetVolumeTypeByIdRequest {
	opt := new(GetVolumeTypeByIdRequest)
	opt.VolumeTypeId = pvolumeTypeId
	return opt
}

type GetVolumeTypeByIdRequest struct {
	lscommon.VolumeTypeCommon
}
