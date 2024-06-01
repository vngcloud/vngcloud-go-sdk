package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type ListSnapshotsByBlockVolumeIdResponse struct {
	Items      []Snapshot `json:"items"`
	Page       int        `json:"page"`
	PageSize   int        `json:"pageSize"`
	TotalPages int        `json:"totalPages"`
	TotalItems int        `json:"totalItems"`
}

type CreateSnapshotByBlockVolumeIdResponse struct {
	Snapshot
}

type (
	Snapshot struct {
		BackendID           string         `json:"backendId"`
		BackendSnapshotID   *string        `json:"backendSnapshotId"`
		BackendSnapshotName *string        `json:"backendSnapshotName"`
		CreatedAt           string         `json:"createdAt"`
		DeletedAt           *string        `json:"deletedAt"`
		ID                  string         `json:"id"`
		ParentID            *string        `json:"parentId"`
		ParentType          string         `json:"parentType"`
		ProjectID           string         `json:"projectId"`
		Size                int64          `json:"size"`
		SnapshotConfig      Config         `json:"snapshotConfig"`
		SnapshotTime        int64          `json:"snapshotTime"`
		SnapshotVolumeId    string         `json:"snapshotVolumeId"`
		Status              string         `json:"status"`
		UpdatedAt           string         `json:"updatedAt"`
		UserId              int64          `json:"userId"`
		VolumeId            string         `json:"volumeId"`
		VolumeSnapshot      VolumeSnapshot `json:"volumeSnapshot"`
		Name                string         `json:"name"`
		Description         string         `json:"description"`
	}

	Config struct {
		IsPermanently bool  `json:"isPermanently"`
		RetainedDays  int64 `json:"retainedDays"`
	}

	VolumeSnapshot struct {
		BackendPool      string  `json:"backendPool"`
		BackendPrefix    string  `json:"backendPrefix"`
		BackendStatus    string  `json:"backendStatus"`
		BackendUuid      string  `json:"backendUuid"`
		BootIndex        int     `json:"bootIndex"`
		Bootable         bool    `json:"bootable"`
		CreatedAt        string  `json:"createdAt"`
		Id               string  `json:"id"`
		MultiAttach      bool    `json:"multiAttach"`
		Name             string  `json:"name"`
		Product          string  `json:"product"`
		ProjectId        string  `json:"projectId"`
		Size             int64   `json:"size"`
		Status           string  `json:"status"`
		Type             string  `json:"type"`
		UpdatedAt        string  `json:"updatedAt"`
		Uuid             string  `json:"uuid"`
		VolumeId         int64   `json:"volumeId"`
		VolumeTypeId     string  `json:"volumeTypeId"`
		VolumeTypeZoneId string  `json:"volumeTypeZoneId"`
		EncryptionType   *string `json:"encryptionType"`
	}
)

func (s *ListSnapshotsByBlockVolumeIdResponse) ToEntityListSnapshots() *lsentity.ListSnapshots {
	sl := new(lsentity.ListSnapshots)

	for _, item := range s.Items {
		sl.Items = append(sl.Items, item.toEntitySnapshot())
	}

	sl.TotalPages = s.TotalPages
	sl.Page = s.Page
	sl.PageSize = s.PageSize
	sl.TotalItems = s.TotalItems

	return sl
}

func (s *Snapshot) toEntitySnapshot() *lsentity.Snapshot {
	return &lsentity.Snapshot{
		Id:        s.ID,
		Name:      s.Name,
		CreatedAt: s.CreatedAt,
		VolumeId:  s.VolumeId,
		Status:    s.Status,
		Size:      s.Size,
	}
}

func (s *CreateSnapshotByBlockVolumeIdResponse) ToEntitySnapshot() *lsentity.Snapshot {
	return s.toEntitySnapshot()
}
