package snapshot

import "strings"

func NewListOpts(pPage, pSize int, pVolumeID, pStatus, pName string) IListOptsBuilder {
	opts := new(ListOpts)
	opts.Page = pPage
	opts.Size = pSize

	if pVolumeID != "" {
		opts.VolumeID = pVolumeID
	}

	if pStatus != "" {
		opts.Status = strings.ToLower(pStatus)
	}

	if pName != "" {
		opts.Name = pName
	}

	return opts
}
