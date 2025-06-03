package v2

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

type ListSecgroupRequest struct {
	lscommon.UserAgent
}

func (s *ListSecgroupRequest) AddUserAgent(pagent ...string) IListSecgroupRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

type DeleteSecgroupByIdRequest struct { //__________________________________________________________________________________
	lscommon.UserAgent
	lscommon.SecgroupCommon
}

func (s *DeleteSecgroupByIdRequest) AddUserAgent(pagent ...string) IDeleteSecgroupByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
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
