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
