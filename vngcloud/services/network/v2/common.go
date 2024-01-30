package v2

type SecgroupV2Common struct {
	SecgroupUUID string
}

func (s *SecgroupV2Common) GetSecgroupUUID() string {
	return s.SecgroupUUID
}

type NetworkV2Common struct {
	CommonNetworkUUID string
}

func (s *NetworkV2Common) GetNetworkUUID() string {
	return s.CommonNetworkUUID
}
