package volume

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/pagination"
)

// ********************************************* List Volume API Response **********************************************

type (
	ListResponse struct {
		Page      int64    `json:"page"`
		PageSize  int64    `json:"pageSize"`
		TotalPage int64    `json:"totalPage"`
		TotalItem int64    `json:"totalItem"`
		ListData  []Common `json:"listData"`

		pagination.IPage
	}

	Common struct {
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

func (s *ListResponse) ToVolumeListObject() *objects.VolumeList {
	vl := new(objects.VolumeList)
	for i, _ := range s.ListData {
		vol := s.toVolumeObject(i)
		if vol != nil {
			vl.Items = append(vl.Items, *vol)
		}
	}

	return vl
}

func (s *ListResponse) toVolumeObject(pIdx int) *objects.Volume {
	if s == nil {
		return nil
	}

	if pIdx >= 0 && pIdx < len(s.ListData) {
		vol := s.ListData[pIdx]
		return &objects.Volume{
			VolumeId:  vol.UUID,
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

// ********************************************* Create Volume API Response ********************************************

type (
	CreateResponse struct {
		Data Common `json:"data"`
	}
)

func (s *CreateResponse) ToVolumeObject() *objects.Volume {
	if s == nil {
		return nil
	}

	return &objects.Volume{
		VolumeId:  s.Data.UUID,
		Name:      s.Data.Name,
		Size:      s.Data.Size,
		Status:    s.Data.Status,
		CreatedAt: s.Data.CreatedAt,
		UpdatedAt: s.Data.UpdatedAt,
		VmId:      s.Data.ServerID,
	}
}

// ********************************************** Get Volume API Response **********************************************

type (
	GetResponse struct {
		Data Common `json:"data"`
	}
)

func (s *GetResponse) ToVolumeObject() *objects.Volume {
	if s == nil {
		return nil
	}

	return &objects.Volume{
		VolumeId:         s.Data.UUID,
		VolumeTypeID:     s.Data.VolumeTypeID,
		Name:             s.Data.Name,
		Size:             s.Data.Size,
		Status:           s.Data.Status,
		CreatedAt:        s.Data.CreatedAt,
		UpdatedAt:        s.Data.UpdatedAt,
		VmId:             s.Data.ServerID,
		PersistentVolume: s.Data.PersistentVolume,
		AttachedMachine:  s.Data.ServerIDList,
	}
}

// ************************************************** ListAllResponse **************************************************

type (
	ListAllResponse struct {
		Page      int64    `json:"page"`
		PageSize  int64    `json:"pageSize"`
		TotalPage int64    `json:"totalPage"`
		TotalItem int64    `json:"totalItem"`
		ListData  []Common `json:"listData"`
	}
)

func (s *ListAllResponse) ToListVolumeObjects() []*objects.Volume {
	var volumes []*objects.Volume
	for i, _ := range s.ListData {
		vol := s.ToVolumeObject(i)
		if vol != nil {
			volumes = append(volumes, vol)
		}
	}

	return volumes
}

func (s *ListAllResponse) ToVolumeObject(pIdx int) *objects.Volume {
	if s == nil {
		return nil
	}

	if pIdx >= 0 && pIdx < len(s.ListData) {
		vol := s.ListData[pIdx]
		return &objects.Volume{
			VolumeId:  vol.UUID,
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
