package flavor

func NewGetOpts(pprojectId, pflavorId string) IGetOptsBuilder {
	opts := new(GetOpts)
	opts.ProjectID = pprojectId
	opts.FlavorID = pflavorId
	return opts
}
