package volume_attach

type AttachCommonOpts struct {
	PersistentVolume bool `json:"persistentVolume"`

	InstanceID string
	VolumeID   string
}

func (s *AttachCommonOpts) GetInstanceID() string {
	return s.InstanceID
}

func (s *AttachCommonOpts) GetVolumeID() string {
	return s.VolumeID
}

func (s *AttachCommonOpts) ToRequestBody() interface{} {
	return s
}
