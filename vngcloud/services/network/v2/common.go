package v2

type SecgroupCommon struct {
	SecgroupId string
}

func (s *SecgroupCommon) GetSecgroupId() string {
	return s.SecgroupId
}

type NetworkCommon struct {
	NetworkId string
}

func (s *NetworkCommon) GetNetworkId() string {
	return s.NetworkId
}
