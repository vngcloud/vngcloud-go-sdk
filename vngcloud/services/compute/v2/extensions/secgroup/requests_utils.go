package secgroup

func NewUpdateOpts(pprojectID, pinstanceID string, psecgroupIDs []string) IUpdateSecgroupOptsBuilder {
	opts := new(UpdateSecgroupOpts)
	opts.ProjectID = pprojectID
	opts.InstanceID = pinstanceID
	opts.ServerID = pinstanceID
	opts.SecurityGroup = psecgroupIDs

	return opts
}
