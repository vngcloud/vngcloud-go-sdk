package v2

func NewDeleteSecgroupRequest(psecgroupId string) IDeleteSecgroupRequest {
	return &DeleteSecgroupRequest{
		SecgroupId: psecgroupId,
	}
}

type DeleteSecgroupRequest struct {
	SecgroupId string
}

func (s *DeleteSecgroupRequest) GetSecgroupId() string {
	return s.SecgroupId
}
