package entity

type VolumeType struct {
	Id         string
	Name       string
	Iops       int
	MaxSize    int
	MinSize    int
	ThroughPut int
	ZoneId     string
}

type ListVolumeType struct {
	VolumeTypes []*VolumeType
}
