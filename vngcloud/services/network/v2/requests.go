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

func NewListAddressPairsByVirtualAddressIdRequest(pvirtualAddressId string) IListAddressPairsByVirtualAddressIdRequest {
	opts := new(ListAddressPairsByVirtualAddressIdRequest)
	opts.VirtualAddressId = pvirtualAddressId
	return opts
}

func NewCreateAddressPairRequest(pvirtualAddressId, pinternalNicId string) ICreateAddressPairRequest {
	opts := new(CreateAddressPairRequest)
	opts.VirtualAddressId = pvirtualAddressId
	opts.InternalNetworkInterfaceId = pinternalNicId
	return opts
}

func NewCreateSecgroupRequest(pname, pdescription string) ICreateSecgroupRequest {
	return &CreateSecgroupRequest{
		Name:        pname,
		Description: pdescription,
	}
}

func NewDeleteSecgroupByIdRequest(psecgroupId string) IDeleteSecgroupByIdRequest {
	opts := new(DeleteSecgroupByIdRequest)
	opts.SecgroupId = psecgroupId
	return opts
}

func NewGetSecgroupByIdRequest(psecgroupId string) IGetSecgroupByIdRequest {
	opt := new(GetSecgroupByIdRequest)
	opt.SecgroupId = psecgroupId
	return opt
}

func NewListSecgroupRequest() IListSecgroupRequest {
	return &ListSecgroupRequest{}
}
