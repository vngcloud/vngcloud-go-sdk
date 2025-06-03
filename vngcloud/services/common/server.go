package common

type ServerCommon struct {
	ServerId string
}

func (s *ServerCommon) GetServerId() string {
	return s.ServerId
}

type ServerGroupCommon struct {
	ServerGroupId string
}

func (s *ServerGroupCommon) GetServerGroupId() string {
	return s.ServerGroupId
}
