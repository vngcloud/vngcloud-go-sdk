package common

type Project struct {
	Id string
}

func (s *Project) GetPorjectId() string {
	return s.Id
}

func (s *Project) SetProjectId(pid string) {
	s.Id = pid
}

type Paging struct {
	Page int
	Size int
}

func (s *Paging) GetPage() int {
	return s.Page
}

func (s *Paging) GetSize() int {
	return s.Size
}

func (s *Paging) SetPage(ppage int) *Paging {
	s.Page = ppage
	return s
}

func (s *Paging) SetSize(psize int) *Paging {
	s.Size = psize
	return s
}
