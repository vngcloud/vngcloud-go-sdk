package server

type GetOptsBuilder interface {
	ExtractProjectID() string
	ExtractInstanceID() string
}

type GetServerOpts struct {
	InstanceID string
	ProjectID  string
}

func (s GetServerOpts) ExtractProjectID() string {
	return s.ProjectID
}

func (s GetServerOpts) ExtractInstanceID() string {
	return s.InstanceID
}
