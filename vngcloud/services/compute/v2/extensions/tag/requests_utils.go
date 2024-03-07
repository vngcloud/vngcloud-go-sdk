package tag

func NewCreateOpts(pProjectID, pInstanceID string, popts *CreateOpts) ICreateOptsBuilder {
	popts.ProjectID = pProjectID
	popts.ServerID = pInstanceID
	return popts
}

func NewGetOpts(pProjectID, pInstanceID string) IGetOptsBuilder {
	opts := new(GetOpts)
	opts.ProjectID = pProjectID
	opts.ServerID = pInstanceID
	return opts
}
