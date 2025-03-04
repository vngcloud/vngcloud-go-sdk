package v1

func NewGetPortalInfoRequest(pbackendProjectId string) IGetPortalInfoRequest {
	return &GetPortalInfoRequest{
		BackEndProjectId: pbackendProjectId,
	}
}

func NewListProjectsRequest() IListProjectsRequest {
	return &ListProjectsRequest{}
}
