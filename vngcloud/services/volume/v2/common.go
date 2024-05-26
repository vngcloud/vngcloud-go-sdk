package v2

type BlockVolumeCommon struct {
	BlockVolumeId string
}

func (s *BlockVolumeCommon) GetBlockVolumeId() string {
	return s.BlockVolumeId
}
