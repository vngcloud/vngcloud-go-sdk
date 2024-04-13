package snapshot

func NewCreateOpts(pprojectID, pVolumeID string, popts *CreateOpts) ICreateOptsBuilder {
	popts.ProjectID = pprojectID
	popts.VolumeID = pVolumeID
	return popts
}

func NewDeleteOpts(pProjectID, pSnapshotID string) IDeleteOptsBuilder {
	opts := new(DeleteOpts)
	opts.ProjectID = pProjectID
	opts.SnapshotID = pSnapshotID
	opts.VolumeID = "undefined"
	return opts
}

func NewListVolumeOpts(pprojectId, pVolumeId string, ppage, psize int) IListVolumeOptsBuilder {
	opts := new(ListVolumeOpts)
	opts.ProjectID = pprojectId
	opts.VolumeID = pVolumeId
	opts.Page = ppage
	opts.Size = psize
	return opts
}
