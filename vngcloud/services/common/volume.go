package common

type BlockVolumeCommon struct {
	BlockVolumeId string
}

func (s *BlockVolumeCommon) GetBlockVolumeId() string {
	return s.BlockVolumeId
}

type VolumeTypeCommon struct {
	VolumeTypeId string
}

func (s *VolumeTypeCommon) GetVolumeTypeId() string {
	return s.VolumeTypeId
}
