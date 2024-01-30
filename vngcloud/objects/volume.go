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
}
