package pool

func NewCreateResponse() ICreateResponse {
	return &CreateResponse{}
}

func NewListPoolsBasedLoadBalancerResponse() IListPoolsBasedLoadBalancerResponse {
	return &ListPoolsBasedLoadBalancerResponse{}
}
