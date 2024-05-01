package pool

func NewCreateResponse() ICreateResponse {
	return &CreateResponse{}
}

func NewListPoolsBasedLoadBalancerResponse() IListPoolsBasedLoadBalancerResponse {
	return &ListPoolsBasedLoadBalancerResponse{}
}

func NewGetResponse() IGetResponse {
	return &GetResponse{}
}

func NewGetMemberResponse() IGetMemberResponse {
	return &GetMemberResponse{}
}

func NewGetHealthMonitorResponse() IGetHealthMonitorResponse {
	return &GetHealthMonitorResponse{}
}
