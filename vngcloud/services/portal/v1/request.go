package v1

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
)

type IGetPortalInfoRequest interface {
	GetBackEndProjectId() string
	ToRequest() lsclient.IRequest
}

func NewGetPortalInfoRequest(pbackendProjectId string) IGetPortalInfoRequest {
	return &getPortalInfoRequest{
		BackEndProjectId: pbackendProjectId,
	}
}

type getPortalInfoRequest struct {
	BackEndProjectId string
}

func (s *getPortalInfoRequest) ToRequest() lsclient.IRequest {
	return lsclient.NewRequest().
		WithOkCodes(200)
}

func (s *getPortalInfoRequest) GetBackEndProjectId() string {
	return s.BackEndProjectId
}
