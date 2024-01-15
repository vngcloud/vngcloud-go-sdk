package volume

func NewCreateResponse() ICreateResponse {
	return &CreateResponse{}
}

func NewGetResponse() IGetResponse {
	return &GetResponse{}
}

func NewListResponse() IListResponse {
	return &ListResponse{}
}

func NewListAllResponse() IListAllResponse {
	return &ListAllResponse{}
}
