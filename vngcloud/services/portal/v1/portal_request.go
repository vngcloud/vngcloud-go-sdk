package v1

type GetPortalInfoRequest struct {
	BackEndProjectId string
}

func NewGetPortalInfoRequest(pbackendProjectId string) IGetPortalInfoRequest {
	return &GetPortalInfoRequest{
		BackEndProjectId: pbackendProjectId,
	}
}

func (s *GetPortalInfoRequest) GetBackEndProjectId() string {
	return s.BackEndProjectId
}
