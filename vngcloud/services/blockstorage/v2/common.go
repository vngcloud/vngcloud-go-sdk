package v2

type BlockStorageV2Common struct {
	VolumeID string
}

func (s *BlockStorageV2Common) GetVolumeID() string {
	return s.VolumeID
}
