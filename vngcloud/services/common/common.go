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

type Zone string

const (
	HCM_03_1A_ZONE     = "HCM03-1A"
	HCM_03_1B_ZONE     = "HCM03-1B"
	HCM_03_1C_ZONE     = "HCM03-1C"
	HCM_03_BKK_01_ZONE = "HCM03-BKK-01"
)

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

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
