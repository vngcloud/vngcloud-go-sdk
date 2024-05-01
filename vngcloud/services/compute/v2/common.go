package v2

type ServerV2Common struct {
	ServerID string
}

func (s *ServerV2Common) GetServerID() string {
	return s.ServerID
}

type ResourceV2Common struct {
	ResourceID string
}

func (s *ResourceV2Common) GetResourceID() string {
	return s.ResourceID
}
