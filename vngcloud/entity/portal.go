package entity

type Portal struct {
	ProjectID string
	UserID    int
}

type Quota struct {
	Description string
	Name        string
	Type        string
	Limit       int
	Used        int
}

type ListQuotas struct {
	Items []*Quota
}

func (s *ListQuotas) FindQuotaByName(name string) *Quota {
	for _, q := range s.Items {
		if q.Name == name {
			return q
		}
	}

	return nil
}

type ListPortals struct {
	Items []*Portal
}

func NewListPortals() *ListPortals {
	return &ListPortals{
		Items: make([]*Portal, 0),
	}
}

func (s *ListPortals) At(pindex int) *Portal {
	if pindex < 0 || pindex >= len(s.Items) {
		return nil
	}

	return s.Items[pindex]
}
