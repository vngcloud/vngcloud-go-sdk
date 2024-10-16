package v1

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type GetPortalInfoResponse struct {
	ProjectID string `json:"projectId"`
	UserID    int    `json:"userId"`
}

func (s *GetPortalInfoResponse) ToEntityPortal() *lsentity.Portal {
	return &lsentity.Portal{
		ProjectID: s.ProjectID,
		UserID:    s.UserID,
	}
}

type ListProjectsResponse struct {
	Projects []struct {
		ProjectId string `json:"projectId"`
		UserId    int    `json:"userId"`
	}
}

func (s *ListProjectsResponse) ToEntityListPortals() *lsentity.ListPortals {
	listPortals := lsentity.NewListPortals()
	for _, p := range s.Projects {
		listPortals.Items = append(listPortals.Items, &lsentity.Portal{
			ProjectID: p.ProjectId,
			UserID:    p.UserId,
		})
	}

	return listPortals
}
