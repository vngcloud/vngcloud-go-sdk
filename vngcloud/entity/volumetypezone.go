package entity

type VolumeTypeZone struct {
	Id   string
	Name string
}

type ListVolumeTypeZones struct {
	VolumeTypeZones []*VolumeTypeZone
}
