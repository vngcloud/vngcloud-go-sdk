package volume

func NewGetOpts(pProjectID, pVolumeID string) IGetOptsBuilder {
	opts := new(GetOpts)
	opts.ProjectID = pProjectID
	opts.VolumeID = pVolumeID
	return opts
}

func NewDeleteOpts(pProjectID, pVolumeID string) IDeleteOptsBuilder {
	opts := new(DeleteOpts)
	opts.ProjectID = pProjectID
	opts.VolumeID = pVolumeID
	return opts
}

func NewCreateOpts(pProjectID string, pOpts *CreateOpts) ICreateOptsBuilder {
	pOpts.ProjectID = pProjectID
	return pOpts
}

func NewListOpts(pProjectID, pName string, pPage, pSize int) IListOptsBuilder {
	opts := new(ListOpts)
	opts.Page = pPage
	opts.Size = pSize
	opts.Name = pName
	opts.ProjectID = pProjectID
	return opts
}

func NewListAllOpts(pProjectID string) IListAllOptsBuilder {
	opts := new(ListAllOpts)
	opts.ProjectID = pProjectID
	return opts
}
