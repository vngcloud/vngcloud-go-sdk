package tag

import "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"

// ***************************************** Response of Create Attachment API *****************************************

type ICreateResponse interface {
	ToServerTagObject() *objects.ServerTag
}

type IGetResponse interface {
	ToServerTagObject() *objects.VolumeAttach
}
