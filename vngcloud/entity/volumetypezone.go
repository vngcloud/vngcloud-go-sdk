package entity

type VolumeTypeZone struct {
	Id       string
	Name     string
	PoolName []string
}

type ListVolumeTypeZones struct {
	VolumeTypeZones []*VolumeTypeZone
}
