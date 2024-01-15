package cluster

func NewGetOpts(pProjectID, pClusterID string) IGetOptsBuilder {
	opts := new(GetOpts)
	opts.ProjectID = pProjectID
	opts.ClusterID = pClusterID
	return opts
}

func NewUpdateSecgroupOpts(pProjectID string, pOpts *UpdateSecgroupOpts) IUpdateSecgroupOptsBuilder {
	pOpts.ProjectID = pProjectID
	return pOpts
}
