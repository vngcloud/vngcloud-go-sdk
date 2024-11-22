package v1

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

type GetPortalInfoRequest struct {
	BackEndProjectId string
}

func (s *GetPortalInfoRequest) GetBackEndProjectId() string {
	return s.BackEndProjectId
}

type ListProjectsRequest struct {
	lscommon.UserAgent
}

func (s *ListProjectsRequest) AddUserAgent(pagent ...string) IListProjectsRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}
