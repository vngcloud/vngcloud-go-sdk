package inter

import "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

func NewCreateLoadBalancerRequest(puserId, pname, ppackageId, pbeSubnetId, psubnetId string) ICreateLoadBalancerRequest {
	opt := new(CreateLoadBalancerRequest)
	opt.SetPortalUserId(puserId)
	opt.Name = pname
	opt.PackageID = ppackageId
	opt.Scheme = InterVpcLoadBalancerScheme
	opt.BackEndSubnetId = pbeSubnetId
	opt.SubnetID = psubnetId
	opt.Type = CreateOptsTypeOptLayer4
	opt.ZoneId = common.HCM_03_1A_ZONE
	return opt
}
