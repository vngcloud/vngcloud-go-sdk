package v2

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

type ListSecgroupRequest struct {
}

type DeleteSecgroupByIdRequest struct { //__________________________________________________________________________________
	SecgroupId string
}

func (s *DeleteSecgroupByIdRequest) GetSecgroupId() string {
	return s.SecgroupId
}

type CreateSecgroupRequest struct { //__________________________________________________________________________________
	Name        string `json:"name"`
	Description string `json:"description"`

	lscommon.UserAgent
}

func (s *CreateSecgroupRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateSecgroupRequest) AddUserAgent(pagent ...string) ICreateSecgroupRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *CreateSecgroupRequest) GetSecgroupName() string {
	return s.Name
}

type GetSecgroupByIdRequest struct { //_________________________________________________________________________________
	lscommon.SecgroupCommon
	lscommon.UserAgent
}

func (s *GetSecgroupByIdRequest) AddUserAgent(pagent ...string) IGetSecgroupByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}
