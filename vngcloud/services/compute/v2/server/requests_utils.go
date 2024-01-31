package server

func NewGetOpts(pProjectID, pServerID string) IGetOptsBuilder {
	opts := new(GetOpts)
	opts.ProjectID = pProjectID
	opts.ServerID = pServerID
	return opts
}
