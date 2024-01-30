package volume_actions

import "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"

type IResizeResponse interface {
	ToResizeVolumeObject() *objects.ResizeVolume
}

type IMappingResponse interface {
	ToMappingVolumeObject() *objects.MappingVolume
}
