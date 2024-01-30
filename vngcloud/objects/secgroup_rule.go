package objects

type ClusterSecgroup struct {
	UUID         string
	ClusterID    string
	SecGroupID   string
	Master       bool
	SecgroupName string
}

type SecgroupRule struct {
	UUID           string
	SecgroupUUID   string
	Direction      string
	EtherType      string
	PortRangeMax   int
	PortRangeMin   int
	Protocol       string
	Description    string
	RemoteIPPrefix string
}
