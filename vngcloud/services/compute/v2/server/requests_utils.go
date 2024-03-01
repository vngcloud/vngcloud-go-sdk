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
	pOpts.ProjectID = pProjectID
	return pOpts
}

func NewListOpts(pProjectID, pName string, pPage, pSize int) IListOptsBuilder {
	opts := new(ListOpts)
	opts.ProjectID = pProjectID
	opts.Name = pName
	opts.Page = func(page int) int {
		if page < 1 {
			return 1
		}
		return page
	}(pPage)

	opts.Size = func(size int) int {
		if size < 1 {
			return 10
		}
		return size
	}(pSize)

	return opts
}
