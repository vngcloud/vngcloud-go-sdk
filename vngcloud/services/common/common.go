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
