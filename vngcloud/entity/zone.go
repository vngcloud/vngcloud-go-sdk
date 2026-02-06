package entity

type Zone struct {
	Uuid          string
	Name          string
	OpenstackZone string
}

type ListZones struct {
	Items []*Zone
}
