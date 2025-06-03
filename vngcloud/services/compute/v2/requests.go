package v2

func NewCreateServerRequest(pname, pimageId, pflavorId, prootDiskType string, prootDiskSize int) ICreateServerRequest {
	opt := new(CreateServerRequest)
	opt.Name = pname
	opt.ImageId = pimageId
	opt.FlavorId = pflavorId
	opt.RootDiskTypeId = prootDiskType
	opt.RootDiskSize = prootDiskSize
	return opt
}

func NewGetServerByIdRequest(pserverId string) IGetServerByIdRequest {
	opt := new(GetServerByIdRequest)
	opt.ServerId = pserverId
	return opt
}

func NewDeleteServerByIdRequest(pserverId string) IDeleteServerByIdRequest {
	opt := new(DeleteServerByIdRequest)
	opt.ServerId = pserverId
	opt.DeleteAllVolume = false
	return opt
}

func NewUpdateServerSecgroupsRequest(pserverId string, psecgroups ...string) IUpdateServerSecgroupsByServerIdRequest {
	opt := new(UpdateServerSecgroupsByServerIdRequest)
	opt.ServerId = pserverId
	opt.Secgroups = psecgroups
	return opt
}

func NewAttachBlockVolumeRequest(pserverId, pvolumeId string) IAttachBlockVolumeRequest {
	opt := new(AttachBlockVolumeRequest)
	opt.ServerId = pserverId
	opt.BlockVolumeId = pvolumeId
	return opt
}

func NewDetachBlockVolumeRequest(pserverId, pvolumeId string) IDetachBlockVolumeRequest {
	opt := new(DetachBlockVolumeRequest)
	opt.ServerId = pserverId
	opt.BlockVolumeId = pvolumeId
	return opt
}

func NewAttachFloatingIpRequest(pserverId, pniid string) IAttachFloatingIpRequest {
	opt := new(AttachFloatingIpRequest)
	opt.ServerId = pserverId
	opt.InternalNetworkInterfaceId = pniid
	opt.NetworkInterfaceId = pniid
	return opt
}

func NewDetachFloatingIpRequest(pserverId, pwanId, pniid string) IDetachFloatingIpRequest {
	opt := new(DetachFloatingIpRequest)
	opt.ServerId = pserverId
	opt.InternalNetworkInterfaceId = pniid
	opt.NetworkInterfaceId = pniid
	opt.WanId = pwanId
	return opt
}

func NewListServerGroupPoliciesRequest() IListServerGroupPoliciesRequest {
	return new(ListServerGroupPoliciesRequest)
}

func NewDeleteServerGroupByIdRequest(pserverGroupId string) IDeleteServerGroupByIdRequest {
	opt := new(DeleteServerGroupByIdRequest)
	opt.ServerGroupId = pserverGroupId
	return opt
}

func NewListServerGroupsRequest(ppage, psize int) IListServerGroupsRequest {
	opt := new(ListServerGroupsRequest)
	opt.Page = ppage
	opt.Size = psize
	opt.Name = ""

	return opt
}

func NewCreateServerGroupRequest(pname, pdescription, policyId string) ICreateServerGroupRequest {
	opt := new(CreateServerGroupRequest)
	opt.Name = pname
	opt.Description = pdescription
	opt.PolicyId = policyId

	return opt
}
