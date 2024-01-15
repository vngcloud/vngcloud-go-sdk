package listener

func NewCreateResponse() ICreateResponse {
	return &CreateResponse{}
}

func NewGetBasedLoadBalancerResponse() IGetBasedLoadBalancerResponse {
	return &GetBasedLoadBalancerResponse{}
}
