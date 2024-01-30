package snapshot

func NewCreateOpts(pProjectID, pVolumeID, pDesc string, pPermanently bool, pName string, pRetainedDay uint64) ICreateOptsBuilder {
	opts := new(CreateOpts)
	opts.ProjectID = pProjectID
	opts.VolumeID = pVolumeID
	opts.Description = pDesc
	opts.Permanently = pPermanently
	opts.Name = pName
	opts.RetainedDay = pRetainedDay
	return opts
}

func NewDeleteOpts(pProjectID, pVolumeID, pSnapshotID string) IDeleteOptsBuilder {
	opts := new(DeleteOpts)
	opts.ProjectID = pProjectID
	opts.VolumeID = pVolumeID
	opts.SnapshotID = pSnapshotID
	return opts
}
