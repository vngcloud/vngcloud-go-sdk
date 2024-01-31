package v2

type ServerV2Common struct {
	ServerID string
}

func (s *ServerV2Common) GetServerID() string {
	return s.ServerID
}
