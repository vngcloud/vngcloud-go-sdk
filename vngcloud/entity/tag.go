package entity

type Tag struct {
	Key        string
	Value      string
	SystemTag  bool
	ResourceId string
	TagId      string
}

type ListTags struct {
	Items []*Tag
}

func (s *ListTags) Len() int {
	return len(s.Items)
}

func (s *ListTags) Empty() bool {
	return s.Len() < 1
}

func (s *ListTags) Add(tags ...*Tag) {
	s.Items = append(s.Items, tags...)
}
