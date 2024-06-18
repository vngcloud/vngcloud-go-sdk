package v2

type SecgroupCommon struct {
	SecgroupId string
}

func (s *SecgroupCommon) GetSecgroupId() string {
	return s.SecgroupId
}
