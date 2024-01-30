package volume_attach

func NewCreateOpts(pProjectID, pInstanceID, pVolumeID string) ICreateOptsBuilder {
	opts := new(CreateOpts)
	opts.ProjectID = pProjectID
	opts.InstanceID = pInstanceID
	opts.VolumeID = pVolumeID
	opts.PersistentVolume = true
	return opts
}

func NewDeleteOpts(pProjectID, pInstanceID, pVolumeID string) IDeleteOptsBuilder {
	opts := new(CreateOpts)
	opts.ProjectID = pProjectID
	opts.InstanceID = pInstanceID
	opts.VolumeID = pVolumeID
	opts.PersistentVolume = true
	return opts
}
