package entity

type Secgroup struct {
	Id          string
	Name        string
	Description string
	Status      string
}

type ListSecgroups struct {
	Items []*Secgroup
}
