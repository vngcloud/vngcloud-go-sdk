package common

type SubnetCommon struct {
	SubnetId string
}

func (s *SubnetCommon) GetSubnetId() string {
	return s.SubnetId
}
