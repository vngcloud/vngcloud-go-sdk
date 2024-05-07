package loadbalancer

func NewCreateResponse() ICommonResponse {
	return &CreateResponse{}
}

func NewGetResponse() ICommonResponse {
	return &GetResponse{}
}

func NewListBySubnetIDResponse() []*ListBySubnetIDResponse {
	var response []*ListBySubnetIDResponse
	return response
}

func NewListResponse() IListResponse {
	return &ListResponse{}
}

func NewUpdateResponse() IUpdateResponse {
	return &UpdateResponse{}
}

func NewListTagsResponse() []*ListTagsResponse {
	var resp []*ListTagsResponse
	return resp
}
