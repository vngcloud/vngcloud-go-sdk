package snapshot

import (
	"fmt"

	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/pagination"
)

// ********************************************* List Snapshot API response ********************************************

type (
	ListResponse struct {
		Page      int64      `json:"page"`
		PageSize  int64      `json:"pageSize"`
		TotalPage int64      `json:"totalPage"`
		TotalItem int64      `json:"totalItem"`
		Items     []Snapshot `json:"items"`
		pagination.IPage
	}
)

func (s *ListResponse) IsEmpty() (bool, error) {
	if len(s.Items) < 1 {
		return true, nil
	}

	return false, nil
}

func (s *ListResponse) NextPageURL(pOpts pagination.IPageOpts) (string, error) {
	currentPage := s.Page
	totalPages := s.TotalPage
	defaultOpts := pOpts.(*ListOpts)

	if totalPages > currentPage {
		query, err := pOpts.ToListQueryWithParams(&map[string]interface{}{
			"page": defaultOpts.Page + 1,
			"size": defaultOpts.Size,
		})
		if err != nil {
			return "", err
		}

		return query, nil
	}

	return "", nil
}

func (s *ListResponse) GetBody() interface{} {
	return s
}

func (s *ListResponse) NextPage() string {
	currentPage := s.Page
	totalPages := s.TotalPage
	if totalPages > currentPage {
		return fmt.Sprintf("%d", currentPage+1)
	}

	return ""
}

func (s *ListResponse) ToListSnapshotObjects() []*objects.Snapshot {
	var volumes []*objects.Snapshot
	for i, _ := range s.Items {
		vol := s.ToSnapshotObject(i)
		if vol != nil {
			volumes = append(volumes, vol)
		}
	}

	return volumes
}

func (s *ListResponse) ToSnapshotObject(pIdx int) *objects.Snapshot {
	if s == nil {
		return nil
	}

	if pIdx >= 0 && pIdx < len(s.Items) {
		snapshot := s.Items[pIdx]
		return &objects.Snapshot{
			ID:        snapshot.ID,
			CreatedAt: snapshot.CreatedAt,
			Size:      snapshot.Size,
			VolumeID:  snapshot.VolumeId,
			Status:    snapshot.Status,
		}
	}

	return nil
}

// ******************************************** Create Snapshot API Response *******************************************

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
