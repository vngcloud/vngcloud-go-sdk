package v1

func NewGetEndpointByIdRequest(pendpointId string) IGetEndpointByIdRequest {
	opt := new(GetEndpointByIdRequest)
	opt.EndpointId = pendpointId
	return opt
}

func NewCreateEndpointRequest(pname, pserviceId, pvpcId, psubnetId string) ICreateEndpointRequest {
	opts := new(CreateEndpointRequest)
	opts.ResourceInfo.EndpointName = pname
	opts.ResourceInfo.ServiceUuid = pserviceId
	opts.ResourceInfo.VpcUuid = pvpcId
	opts.ResourceInfo.SubnetUuid = psubnetId
	opts.ResourceInfo.PackageUuid = defaultPackageId
	opts.ResourceInfo.EnableDnsName = false
	return opts
}

func NewDeleteEndpointByIdRequest(pendpointId, pvpcId, pendpointServiceId string) IDeleteEndpointByIdRequest {
	opt := new(DeleteEndpointByIdRequest)
	opt.EndpointId = pendpointId
	opt.EndpointUuid = pendpointId
	opt.VpcUuid = pvpcId
	opt.EndpointServiceUuid = pendpointServiceId

	return opt
}

func NewListEndpointsRequest(ppage, psize int) IListEndpointsRequest {
	return &ListEndpointsRequest{
		Page: ppage,
		Size: psize,
	}
}

func NewListTagsByEndpointIdRequest(puserId, pprojectId, pendpointId string) IListTagsByEndpointIdRequest {
	opt := new(ListTagsByEndpointIdRequest)
	opt.Id = pendpointId
	opt.EndpointId = pendpointId
	opt.ProjectId = pprojectId
	opt.SetPortalUserId(puserId)
	return opt
}

func NewCreateTagsWithEndpointIdRequest(puserId, pprojectId, pendpointId string) ICreateTagsWithEndpointIdRequest {
	opt := new(CreateTagsWithEndpointIdRequest)
	opt.ResourceUuid = pendpointId
	opt.EndpointId = pendpointId
	opt.SystemTag = true
	opt.ProjectId = pprojectId
	opt.SetPortalUserId(puserId)

	return opt
}

func NewDeleteTagOfEndpointRequest(puserId, pprojectId, ptagId string) IDeleteTagOfEndpointRequest {
	opt := new(DeleteTagOfEndpointRequest)
	opt.TagId = ptagId
	opt.ProjectId = pprojectId
	opt.SetPortalUserId(puserId)

	return opt
}

func NewUpdateTagValueOfEndpointRequest(puserId, pprojectId, ptagId, pvalue string) IUpdateTagValueOfEndpointRequest {
	opt := new(UpdateTagValueOfEndpointRequest)
	opt.TagId = ptagId
	opt.TagValue = pvalue
	opt.ProjectId = pprojectId
	opt.SetPortalUserId(puserId)

	return opt
}
