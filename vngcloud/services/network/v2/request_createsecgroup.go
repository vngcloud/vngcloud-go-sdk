package v2

type CreateSecgroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (s *CreateSecgroupRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateSecgroupRequest) GetSecgroupName() string {
	return s.Name
}

func NewCreateSecgroupRequest(pname, pdescription string) ICreateSecgroupRequest {
	return &CreateSecgroupRequest{
		Name:        pname,
		Description: pdescription,
	}
}
