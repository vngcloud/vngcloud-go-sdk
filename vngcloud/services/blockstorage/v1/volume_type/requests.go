package volume_type

import "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"

type GetVolumeTypeOptsRequest struct {
	*common.CommonOpts
	VolumeTypeID string
}

func (s *GetVolumeTypeOptsRequest) GetVolumeTypeID() string {
	return s.VolumeTypeID
}
