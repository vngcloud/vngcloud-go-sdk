package volume_actions

func NewResizeOpts(pProjectID, pVolumeTypeID, pVolumeID string, pNewSize uint64) IResizeOptsBuilder {
	opts := new(ResizeOpts)
	opts.ProjectID = pProjectID
	opts.VolumeID = pVolumeID
	opts.VolumeTypeID = pVolumeTypeID
	opts.NewSize = pNewSize
	return opts
}

func NewMappingOpts(pProjectID, pVolumeID string) IMappingOptsBuilder {
	opts := new(MappingOpts)
	opts.ProjectID = pProjectID
	opts.VolumeID = pVolumeID
	return opts
}
