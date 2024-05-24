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
