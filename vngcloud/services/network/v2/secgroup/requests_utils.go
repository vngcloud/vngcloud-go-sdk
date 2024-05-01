package secgroup

func NewCreateOpts(pProjectID string, pOpts *CreateOpts) ICreateOptsBuilder {
	pOpts.ProjectID = pProjectID
	return pOpts
}

func NewDeleteOpts(pProjectID, pSecgroupUUID string) IDeleteOptsBuilder {
	opts := new(DeleteOpts)
	opts.ProjectID = pProjectID
	opts.SecgroupUUID = pSecgroupUUID
	return opts
}

func NewGetOpts(pProjectID, pSecgroupUUID string) IGetOptsBuilder {
	opts := new(GetOpts)
	opts.ProjectID = pProjectID
	opts.SecgroupUUID = pSecgroupUUID
	return opts
}

func NewListOpts(pProjectID, pName string) IListOptsBuilder {
	opts := new(ListOpts)
	opts.ProjectID = pProjectID
	opts.Name = pName
	return opts
}
