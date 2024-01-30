package common

type CommonOpts struct {
	ProjectID string
}

func (s *CommonOpts) GetProjectID() string {
	return s.ProjectID
}
