package server

func NewGetOpts(pProjectID, pServerID string) IGetOptsBuilder {
	opts := new(GetOpts)
	opts.ProjectID = pProjectID
	opts.ServerID = pServerID
	return opts
}

func NewDeleteOpts(pProjectID, pServerID string, pDeleteAllVolume bool) IDeleteOptsBuilder {
	opts := new(DeleteOpts)
	opts.ProjectID = pProjectID
	opts.ServerID = pServerID
	opts.DeleteAllVolume = pDeleteAllVolume
	return opts
}

func NewCreateOpts(pProjectID string, pOpts *CreateOpts) ICreateOptsBuilder {
	opts := new(CreateOpts)
	opts.ProjectID = pProjectID
	return opts
}
