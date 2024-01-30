package objects

type ResizeVolume struct {
	VolumeTypeID     string
	Size             int64
	Status           string
	UUID             string
	Name             string
	PersistentVolume bool
}

type MappingVolume struct {
	UUID string
}
