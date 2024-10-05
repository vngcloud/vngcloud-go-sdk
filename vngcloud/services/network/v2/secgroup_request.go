package v2

func NewCreateSecgroupRequest(pname, pdescription string) ICreateSecgroupRequest {
	return &CreateSecgroupRequest{
		Name:        pname,
		Description: pdescription,
	}
}

func NewDeleteSecgroupByIdRequest(psecgroupId string) IDeleteSecgroupByIdRequest {
	return &DeleteSecgroupByIdRequest{
		SecgroupId: psecgroupId,
	}
}

func NewGetSecgroupByIdRequest(psecgroupId string) IGetSecgroupByIdRequest {
	opt := new(GetSecgroupByIdRequest)
	opt.SecgroupId = psecgroupId
	return opt
}

func NewListSecgroupRequest() IListSecgroupRequest {
	return &ListSecgroupRequest{}
}

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
