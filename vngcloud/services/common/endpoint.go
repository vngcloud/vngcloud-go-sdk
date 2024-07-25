package common

type EndpointCommon struct {
	EndpointId string
}

func (s *EndpointCommon) GetEndpointId() string {
	return s.EndpointId
}
