package subnet

func NewGetResponse() IGetResponse {
	return &GetResponse{}
}

func NewListByNetworkIDResponse() []*ListByNetworkIDResponse {
	var resp []*ListByNetworkIDResponse
	return resp
}
