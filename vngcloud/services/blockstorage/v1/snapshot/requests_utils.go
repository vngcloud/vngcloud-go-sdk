package snapshot

func NewListVolumeOpts(pprojectId, pVolumeId string, ppage, psize int) IListVolumeOptsBuilder {
	opts := new(ListVolumeOpts)
	opts.ProjectID = pprojectId
	opts.VolumeID = pVolumeId
	opts.Page = ppage
	opts.Size = psize
	return opts
}
