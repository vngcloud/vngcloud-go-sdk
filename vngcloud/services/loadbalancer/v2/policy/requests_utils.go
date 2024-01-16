package policy

func NewListOpts(pProjectID string) IListOptsBuilder {
	opts := new(ListOptsBuilder)
	opts.ProjectID = pProjectID
	return opts
}
