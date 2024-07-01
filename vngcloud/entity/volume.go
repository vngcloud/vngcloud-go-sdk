package entity

type Volume struct {
	Name             string
	Id               string
	VolumeTypeID     string
	ClusterId        *string
	VmId             string
	Size             uint64
	IopsId           uint64
	Status           string
	CreatedAt        string
	UpdatedAt        *string
	PersistentVolume bool
	AttachedMachine  []string
	UnderId          string
}

type ListVolumes struct {
	Items []*Volume
}

func (s *ListVolumes) Len() int {
	return len(s.Items)
}

func (s *Volume) AttachedTheInstance(pinstanceId string) bool {
	if s.VmId == pinstanceId {
		return true
	}

	for _, machineId := range s.AttachedMachine {
		if machineId == pinstanceId {
			return true
		}
	}

	return false
}

func (s *Volume) IsAvailable() bool {
	return s.Status == "AVAILABLE"
}

func (s *Volume) IsError() bool {
	return s.Status == "ERROR"
}

func (s *Volume) IsInUse() bool {
	return s.Status == "IN-USE"
}
