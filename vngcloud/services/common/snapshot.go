package common

type SnapshotCommon struct {
	SnapshotId string
}

func (s *SnapshotCommon) GetSnapshotId() string {
	return s.SnapshotId
}
