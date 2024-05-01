package objects

type Volume struct {
	Name             string
	VolumeId         string
	VolumeTypeID     string
	ClusterId        *string
	VmId             *string
	Size             uint64
	IopsId           uint64
	Status           string
	CreatedAt        string
	UpdatedAt        *string
	PersistentVolume bool
	AttachedMachine  []string
}

type VolumeList struct {
	Items      []Volume
	Page       int
	PageSize   int
	TotalPages int
	TotalItems int
}
