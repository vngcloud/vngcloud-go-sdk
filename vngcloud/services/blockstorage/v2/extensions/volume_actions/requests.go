package volume_actions

import (
	bstCm "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/blockstorage/v2"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
)

type ResizeOpts struct {
	NewSize      uint64 `json:"newSize"`         // NewSize is the new size of the volume, in GB
	VolumeTypeID string `json:"newVolumeTypeId"` // VolumeTypeID is the type of the volume

	common.CommonOpts
	bstCm.BlockStorageV2Common
}

func (s *ResizeOpts) ToRequestBody() interface{} {
	return s
}

type MappingOpts struct {
	common.CommonOpts
	bstCm.BlockStorageV2Common
}
