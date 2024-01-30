package volume_attach

import "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"

// ***************************************** Create Volume Attachment Response *****************************************

type CreateResponse struct {
}

func (s *CreateResponse) ToVolumeAttachObject() *objects.VolumeAttach {
	return nil
}

// ***************************************** Delete Volume Attachment Response *****************************************

type DeleteResponse struct{}

func (s *DeleteResponse) ToVolumeAttachObject() *objects.VolumeAttach {
	return nil
}
