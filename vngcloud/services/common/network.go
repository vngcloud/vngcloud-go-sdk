package common

type NetworkCommon struct {
	NetworkId string
}

func (s *NetworkCommon) GetNetworkId() string {
	return s.NetworkId
}
