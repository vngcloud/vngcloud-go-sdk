package volume_actions

import "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"

type ResizeResponse struct {
	Data struct {
		Uuid               string        `json:"uuid"`
		ProjectId          string        `json:"projectId"`
		Name               string        `json:"name"`
		Size               int64         `json:"size"`
		Status             string        `json:"status"`
		VolumeTypeId       string        `json:"volumeTypeId"`
		VolumeTypeZoneName string        `json:"volumeTypeZoneName"`
		Iops               string        `json:"iops"`
		CreatedAt          string        `json:"createdAt"`
		UpdatedAt          string        `json:"updatedAt"`
		Bootable           bool          `json:"bootable"`
		EncryptionType     interface{}   `json:"encryptionType"`
		BootIndex          int           `json:"bootIndex"`
		MultiAttach        bool          `json:"multiAttach"`
		ServerIdList       []interface{} `json:"serverIdList"`
		Location           string        `json:"location"`
		Product            string        `json:"product"`
		PersistentVolume   bool          `json:"persistentVolume"`
	} `json:"data"`
}

func (s *ResizeResponse) ToResizeVolumeObject() *objects.ResizeVolume {
	return &objects.ResizeVolume{
		UUID:             s.Data.Uuid,
		Name:             s.Data.Name,
		Size:             s.Data.Size,
		Status:           s.Data.Status,
		VolumeTypeID:     s.Data.VolumeTypeId,
		PersistentVolume: s.Data.PersistentVolume,
	}
}

type MappingResponse struct {
	Uuid               string      `json:"uuid"`
	ProjectId          interface{} `json:"projectId"`
	Name               interface{} `json:"name"`
	Size               interface{} `json:"size"`
	Status             interface{} `json:"status"`
	VolumeTypeId       interface{} `json:"volumeTypeId"`
	VolumeTypeZoneName interface{} `json:"volumeTypeZoneName"`
	Iops               interface{} `json:"iops"`
	ServerId           interface{} `json:"serverId"`
	CreatedAt          interface{} `json:"createdAt"`
	UpdatedAt          interface{} `json:"updatedAt"`
	Bootable           interface{} `json:"bootable"`
	EncryptionType     interface{} `json:"encryptionType"`
	BootIndex          interface{} `json:"bootIndex"`
	MultiAttach        interface{} `json:"multiAttach"`
	ServerIdList       interface{} `json:"serverIdList"`
	Location           interface{} `json:"location"`
	Product            interface{} `json:"product"`
	PersistentVolume   interface{} `json:"persistentVolume"`
}

func (s *MappingResponse) ToMappingVolumeObject() *objects.MappingVolume {
	return &objects.MappingVolume{
		UUID: s.Uuid,
	}
}
