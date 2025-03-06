package entity

type ServerGroupPolicy struct {
	Name         string
	UUID         string
	Status       string
	Descriptions map[string]string
}

type ListServerGroupPolicies struct {
	Items []*ServerGroupPolicy
}

func (s *ListServerGroupPolicies) Add(item *ServerGroupPolicy) {
	s.Items = append(s.Items, item)
}
