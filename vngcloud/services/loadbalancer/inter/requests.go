package inter

func NewCreateLoadBalancerRequest(puserId, pname, ppackageId, pbeSubnetId, psubnetId string) ICreateLoadBalancerRequest {
	opt := new(CreateLoadBalancerRequest)
	opt.SetPortalUserId(puserId)
	opt.Name = pname
	opt.PackageID = ppackageId
	opt.Scheme = InterVpcLoadBalancerScheme
	opt.BackEndSubnetId = pbeSubnetId
	opt.SubnetID = psubnetId
	opt.Type = CreateOptsTypeOptLayer4
	return opt
}
