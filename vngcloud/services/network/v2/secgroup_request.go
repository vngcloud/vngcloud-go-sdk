package v2

func NewCreateSecgroupRequest(pname, pdescription string) ICreateSecgroupRequest {
	return &CreateSecgroupRequest{
		Name:        pname,
		Description: pdescription,
	}
}

func NewDeleteSecgroupRequest(psecgroupId string) IDeleteSecgroupRequest {
	return &DeleteSecgroupRequest{
		SecgroupId: psecgroupId,
	}
}

func NewGetSecgroupByIdRequest(psecgroupId string) IGetSecgroupByIdRequest {
	opt := new(GetSecgroupByIdRequest)
	opt.SecgroupId = psecgroupId
	return opt
}

type DeleteSecgroupRequest struct { //__________________________________________________________________________________
	SecgroupId string
}

func (s *DeleteSecgroupRequest) GetSecgroupId() string {
	return s.SecgroupId
}

type CreateSecgroupRequest struct { //__________________________________________________________________________________
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (s *CreateSecgroupRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateSecgroupRequest) GetSecgroupName() string {
	return s.Name
}

type GetSecgroupByIdRequest struct { //_________________________________________________________________________________
	SecgroupCommon
}
