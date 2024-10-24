package v2

func NewCreateLoadBalancerRequest(pname string) ICreateLoadBalancerRequest {
	return &CreateLoadBalancerRequest{
		Name:   pname,
		Scheme: InternetLoadBalancerScheme,
		Type:   LoadBalancerTypeLayer4,
	}
}

func NewGetLoadBalancerByIdRequest(plbId string) IGetLoadBalancerByIdRequest {
	opts := new(GetLoadBalancerByIdRequest)
	opts.LoadBalancerId = plbId
	return opts
}

func NewListLoadBalancersRequest(ppage, psize int) IListLoadBalancersRequest {
	opts := new(ListLoadBalancersRequest)
	opts.Page = ppage
	opts.Size = psize
	return opts
}

func NewDeleteLoadBalancerByIdRequest(plbId string) IDeleteLoadBalancerByIdRequest {
	opts := new(DeleteLoadBalancerByIdRequest)
	opts.LoadBalancerId = plbId
	return opts
}
