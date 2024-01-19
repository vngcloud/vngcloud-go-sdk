package policy

func NewListResponse() IListResponse {
	return &ListResponse{}
}

func NewCreateResponse() ICreateResponse {
	return &CreateResponse{}
}

func NewGetResponse() IGetResponse {
	return &GetResponse{}
}