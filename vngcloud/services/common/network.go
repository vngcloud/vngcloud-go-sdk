package common

type NetworkCommon struct {
	NetworkId string
}

func (s *NetworkCommon) GetNetworkId() string {
	return s.NetworkId
}

type InternalNetworkInterfaceCommon struct {
	InternalNetworkInterfaceId string
}

func (s *InternalNetworkInterfaceCommon) GetInternalNetworkInterfaceId() string {
	return s.InternalNetworkInterfaceId
}

type WanCommon struct {
	WanId string
}

func (s *WanCommon) GetWanId() string {
	return s.WanId
}
