package v2

func NewResizeLoadBalancerByIdRequest(plbId, ppackageId string) IResizeLoadBalancerByIdRequest {
	opts := new(ResizeLoadBalancerByIdRequest)
	opts.LoadBalancerId = plbId
	opts.PackageId = ppackageId
	return opts
}
