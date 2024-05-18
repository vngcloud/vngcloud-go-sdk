package entity

type SecgroupRule struct {
	Id             string
	SecgroupId     string
	Direction      string
	EtherType      string
	Protocol       string
	Description    string
	RemoteIPPrefix string
	PortRangeMax   int
	PortRangeMin   int
}
