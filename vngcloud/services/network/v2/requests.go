package v2

func NewCreateVirtualAddressCrossProjectRequest(pname, pprojectId, psubnetId string) ICreateVirtualAddressCrossProjectRequest {
	opts := new(CreateVirtualAddressCrossProjectRequest)
	opts.Name = pname
	opts.CrossProjectRequest.ProjectId = pprojectId
	opts.CrossProjectRequest.SubnetId = psubnetId
	return opts
}
