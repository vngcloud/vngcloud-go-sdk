package objects

type Snapshot struct {
	ID        string
	CreatedAt string
	VolumeID  string
	Size      int64
	Status    string
	Name      string
}

type SnapshotList struct {
	Items []Snapshot
}
