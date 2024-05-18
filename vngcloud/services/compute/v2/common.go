package v2

type ServerCommon struct {
	ServerId string
}

func (s *ServerCommon) GetServerId() string {
	return s.ServerId
}
