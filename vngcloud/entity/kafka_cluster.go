package entity

const (
	KafkaClusterStatusActive   = "ACTIVE"
	KafkaClusterStatusCreating = "CREATING"
	KafkaClusterStatusError    = "ERROR"
	KafkaClusterStatusStopped  = "STOPPED"
)

type (
	KafkaCluster struct {
		Id                   string
		Name                 string
		Tags                 map[string]string
		KafkaVersion         string
		ServerFlavorId       string
		KafkaBrokerCount     int32
		KafkaStorageType     string
		KafkaStorageSize     int32
		KafkaStorageUsage    []int64
		VserverProjectId     string
		NetworkId            string
		SubnetId             string
		FixedIps             []string
		FloatingIps          []string
		PublicAccess         bool
		MtlsAuthen           bool
		SaslAuthen           bool
		SecurityGroupRules   []KafkaSecurityGroupRule
		ConfigGroupVersionId string
		PortalUserId         int32
		Status               string
		ErrorMessage         string
		CreatedAt            string
		EncryptionVolume     bool
		Iops                 int32
		VolumeType           string
		VolumeTypeZoneId     string
		Ram                  int32
		Vcpus                int32
		InstanceType         string
	}

	KafkaSecurityGroupRule struct {
		Id             string
		Direction      string
		Protocol       string
		PortRangeMin   int32
		PortRangeMax   int32
		RemoteIpPrefix string
	}
)

func (s *KafkaCluster) IsActive() bool {
	return s.Status == KafkaClusterStatusActive
}

func (s *KafkaCluster) HasPublicEndpoint() bool {
	return s.PublicAccess && len(s.FloatingIps) > 0
}

func (s *KafkaCluster) HasPrivateEndpoint() bool {
	return len(s.FixedIps) > 0
}
