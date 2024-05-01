package subnet

type SubnetCommonOpts struct {
	CommonSubnetUUID string
}

func (s *SubnetCommonOpts) GetSubnetUUID() string {
	return s.CommonSubnetUUID
}
