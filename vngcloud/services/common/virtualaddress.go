package common

type VirtualAddressCommon struct {
	VirtualAddressId string
}

func (s *VirtualAddressCommon) GetVirtualAddressId() string {
	return s.VirtualAddressId
}
