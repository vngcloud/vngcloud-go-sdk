package tag

func NewCreateOpts(pProjectID, pInstanceID string, popts *CreateOpts) ICreateOptsBuilder {
	popts.ProjectID = pProjectID
	popts.ResourceID = pInstanceID
	popts.ResourceV2Common.ResourceID = pInstanceID
	return popts
}

func NewGetOpts(pProjectID, pInstanceID string) IGetOptsBuilder {
	opts := new(GetOpts)
	opts.ProjectID = pProjectID
	opts.ResourceID = pInstanceID
	opts.ResourceV2Common.ResourceID = pInstanceID
	return opts
}

func NewUpdateOpts(pProjectID, pInstanceID string, popts *UpdateOpts) IUpdateOptsBuilder {
	popts.ProjectID = pProjectID
	popts.ResourceID = pInstanceID
	popts.ResourceV2Common.ResourceID = pInstanceID
	return popts
}
