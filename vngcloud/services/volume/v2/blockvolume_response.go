package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type CreateBlockVolumeResponse struct {
	Data BlockVolume `json:"data"`
}

type (
	BlockVolume struct {
		UUID               string   `json:"uuid"`
		ProjectID          string   `json:"projectId"`
		Name               string   `json:"name"`
		Size               uint64   `json:"size"`
		Status             string   `json:"status"`
		VolumeTypeID       string   `json:"volumeTypeId"`
		VolumeTypeZoneName string   `json:"volumeTypeZoneName"`
		IOPS               string   `json:"iops"`
		ServerID           *string  `json:"serverId"`
		CreatedAt          string   `json:"createdAt"`
		UpdatedAt          *string  `json:"updatedAt"`
		Bootable           bool     `json:"bootable"`
		EncryptionType     *string  `json:"encryptionType"`
		BootIndex          int      `json:"bootIndex"`
		MultiAttach        bool     `json:"multiAttach"`
		ServerIDList       []string `json:"serverIdList"`
		Location           *string  `json:"location"`
		Product            string   `json:"product"`
		PersistentVolume   bool     `json:"persistentVolume"`
	}
)

func (s *CreateBlockVolumeResponse) ToEntityVolume() *lsentity.Volume {
	if s == nil {
		return nil
	}

	return &lsentity.Volume{
		Id:        s.Data.UUID,
		Name:      s.Data.Name,
		Size:      s.Data.Size,
		Status:    s.Data.Status,
		CreatedAt: s.Data.CreatedAt,
		UpdatedAt: s.Data.UpdatedAt,
		VmId:      s.Data.ServerID,
	}
}
