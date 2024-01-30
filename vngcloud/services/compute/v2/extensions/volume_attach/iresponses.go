package volume_attach

import "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"

// ***************************************** Response of Create Attachment API *****************************************

type ICreateResponse interface {
	ToVolumeAttachObject() *objects.VolumeAttach
}

type IDeleteResponse interface {
	ToVolumeAttachObject() *objects.VolumeAttach
}
