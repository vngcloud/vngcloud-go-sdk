package entity

type Snapshot struct {
	Id         string
	CreatedAt  string
	VolumeId   string
	Size       int64
	VolumeSize int64
	Status     string
	Name       string
}

type ListSnapshots struct {
	Items      []*Snapshot
	TotalPages int
	Page       int
	PageSize   int
	TotalItems int
}
