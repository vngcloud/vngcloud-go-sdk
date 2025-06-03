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

type ListSecgroupRules struct {
	Items []*SecgroupRule
}

func (s *ListSecgroupRules) Len() int {
	return len(s.Items)
}

func (s *ListSecgroupRules) Get(i int) *SecgroupRule {
	if i < 0 || i >= len(s.Items) {
		return nil
	}
	return s.Items[i]
}
