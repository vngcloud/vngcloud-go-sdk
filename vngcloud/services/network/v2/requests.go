package v2

func NewCreateVirtualAddressCrossProjectRequest(pname, pprojectId, psubnetId string) ICreateVirtualAddressCrossProjectRequest {
	opts := new(CreateVirtualAddressCrossProjectRequest)
	opts.Name = pname
	opts.CrossProjectRequest.ProjectId = pprojectId
	opts.CrossProjectRequest.SubnetId = psubnetId
	return opts
}

func NewDeleteVirtualAddressByIdRequest(pvirtualAddressId string) IDeleteVirtualAddressByIdRequest {
	opts := new(DeleteVirtualAddressByIdRequest)
	opts.VirtualAddressId = pvirtualAddressId
	return opts
}

func NewGetVirtualAddressByIdRequest(pvirtualAddressId string) IGetVirtualAddressByIdRequest {
	opts := new(GetVirtualAddressByIdRequest)
	opts.VirtualAddressId = pvirtualAddressId
	return opts
}
