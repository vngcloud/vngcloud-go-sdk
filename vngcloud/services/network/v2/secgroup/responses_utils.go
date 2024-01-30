package secgroup

func NewCreateResponse() ICreateResponse {
	return &CreateResponse{}
}

func NewGetResponse() IGetResponse {
	return &GetResponse{}
}

func NewListResponse() IListResponse {
	return &ListResponse{}
}
