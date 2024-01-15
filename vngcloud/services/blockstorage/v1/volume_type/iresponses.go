package volume_type

import "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"

type IGetVolumeTypeResponse interface {
	ToVolumeTypeObject() (*objects.VolumeType, error)
}
