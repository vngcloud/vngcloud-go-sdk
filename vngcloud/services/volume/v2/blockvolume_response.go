package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type CreateBlockVolumeResponse struct {
	Data BlockVolume `json:"data"`
}

type ListBlockVolumesResponse struct {
	Page      int64         `json:"page"`
	PageSize  int64         `json:"pageSize"`
	TotalPage int64         `json:"totalPage"`
	TotalItem int64         `json:"totalItem"`
	ListData  []BlockVolume `json:"listData"`
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
	return s.Data.toEntityVolume()
}

func (s *ListBlockVolumesResponse) ToEntityListVolumes() *lsentity.ListVolumes {
	lstVolumes := new(lsentity.ListVolumes)
	for _, vol := range s.ListData {
		lstVolumes.Items = append(lstVolumes.Items, vol.toEntityVolume())
	}

	return lstVolumes
}

func (s *ListBlockVolumesResponse) toEntityVolume(pIdx int) *lsentity.Volume {
	if s == nil {
		return nil
	}

	if pIdx >= 0 && pIdx < len(s.ListData) {
		vol := s.ListData[pIdx]
		return &lsentity.Volume{
			Id:        vol.UUID,
			Name:      vol.Name,
			Size:      vol.Size,
			Status:    vol.Status,
			CreatedAt: vol.CreatedAt,
			UpdatedAt: vol.UpdatedAt,
			VmId:      vol.ServerID,
		}
	}

	return nil
}

func (s *BlockVolume) toEntityVolume() *lsentity.Volume {
	return &lsentity.Volume{
		Id:        s.UUID,
		Name:      s.Name,
		Size:      s.Size,
		Status:    s.Status,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
		VmId:      s.ServerID,
	}
}