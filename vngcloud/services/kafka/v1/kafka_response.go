package v1

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
)

// KafkaCluster response payload (wire format, matches vdb-api.json KafkaCluster schema).
type KafkaClusterResponse struct {
	Id                   string                 `json:"id"`
	Name                 string                 `json:"name"`
	Tags                 map[string]string      `json:"tags"`
	KafkaVersion         string                 `json:"kafkaVersion"`
	ServerFlavorId       string                 `json:"serverFlavorId"`
	KafkaBrokerCount     int32                  `json:"kafkaBrokerCount"`
	KafkaStorageType     string                 `json:"kafkaStorageType"`
	KafkaStorageSize     int32                  `json:"kafkaStorageSize"`
	KafkaStorageUsage    []int64                `json:"kafkaStorageUsage"`
	VserverProjectId     string                 `json:"vserverProjectId"`
	NetworkId            string                 `json:"networkId"`
	SubnetId             string                 `json:"subnetId"`
	FixedIps             []string               `json:"fixedIps"`
	FloatingIps          []string               `json:"floatingIps"`
	PublicAccess         bool                   `json:"publicAccess"`
	MtlsAuthen           bool                   `json:"mtlsAuthen"`
	SaslAuthen           bool                   `json:"saslAuthen"`
	SecurityGroupRules   []SecurityGroupRuleDto `json:"securityGroupRules"`
	ConfigGroupVersionId string                 `json:"configGroupVersionId"`
	PortalUserId         int32                  `json:"portalUserId"`
	Status               string                 `json:"status"`
	ErrorMessage         string                 `json:"errorMessage"`
	CreatedAt            string                 `json:"createdAt"`
	EncryptionVolume     bool                   `json:"encryptionVolume"`
	Iops                 int32                  `json:"iops"`
	VolumeType           string                 `json:"volumeType"`
	VolumeTypeZoneId     string                 `json:"volumeTypeZoneId"`
	Ram                  int32                  `json:"ram"`
	Vcpus                int32                  `json:"vcpus"`
	InstanceType         string                 `json:"instanceType"`
}

type SecurityGroupRuleDto struct {
	Id             string `json:"id"`
	Direction      string `json:"direction"`
	Protocol       string `json:"protocol"`
	PortRangeMin   int32  `json:"portRangeMin"`
	PortRangeMax   int32  `json:"portRangeMax"`
	RemoteIpPrefix string `json:"remoteIpPrefix"`
}

func (s *KafkaClusterResponse) ToEntity() *lsentity.KafkaCluster {
	if s == nil {
		return nil
	}
	rules := make([]lsentity.KafkaSecurityGroupRule, 0, len(s.SecurityGroupRules))
	for _, r := range s.SecurityGroupRules {
		rules = append(rules, lsentity.KafkaSecurityGroupRule{
			Id:             r.Id,
			Direction:      r.Direction,
			Protocol:       r.Protocol,
			PortRangeMin:   r.PortRangeMin,
			PortRangeMax:   r.PortRangeMax,
			RemoteIpPrefix: r.RemoteIpPrefix,
		})
	}
	return &lsentity.KafkaCluster{
		Id:                   s.Id,
		Name:                 s.Name,
		Tags:                 s.Tags,
		KafkaVersion:         s.KafkaVersion,
		ServerFlavorId:       s.ServerFlavorId,
		KafkaBrokerCount:     s.KafkaBrokerCount,
		KafkaStorageType:     s.KafkaStorageType,
		KafkaStorageSize:     s.KafkaStorageSize,
		KafkaStorageUsage:    s.KafkaStorageUsage,
		VserverProjectId:     s.VserverProjectId,
		NetworkId:            s.NetworkId,
		SubnetId:             s.SubnetId,
		FixedIps:             s.FixedIps,
		FloatingIps:          s.FloatingIps,
		PublicAccess:         s.PublicAccess,
		MtlsAuthen:           s.MtlsAuthen,
		SaslAuthen:           s.SaslAuthen,
		SecurityGroupRules:   rules,
		ConfigGroupVersionId: s.ConfigGroupVersionId,
		PortalUserId:         s.PortalUserId,
		Status:               s.Status,
		ErrorMessage:         s.ErrorMessage,
		CreatedAt:            s.CreatedAt,
		EncryptionVolume:     s.EncryptionVolume,
		Iops:                 s.Iops,
		VolumeType:           s.VolumeType,
		VolumeTypeZoneId:     s.VolumeTypeZoneId,
		Ram:                  s.Ram,
		Vcpus:                s.Vcpus,
		InstanceType:         s.InstanceType,
	}
}

type ListClustersResponse []KafkaClusterResponse

func (s ListClustersResponse) ToEntity() []lsentity.KafkaCluster {
	out := make([]lsentity.KafkaCluster, 0, len(s))
	for i := range s {
		out = append(out, *s[i].ToEntity())
	}
	return out
}

// UserDto -----------------------------------------------------------------

type KafkaUserResponse struct {
	Id                       string   `json:"id"`
	ClusterId                string   `json:"clusterId"`
	Name                     string   `json:"name"`
	ProduceTopicNames        []string `json:"produceTopicNames"`
	ProduceAll               bool     `json:"produceAll"`
	ConsumeTopicNames        []string `json:"consumeTopicNames"`
	ConsumeAll               bool     `json:"consumeAll"`
	ProduceConsumeTopicNames []string `json:"produceConsumeTopicNames"`
	ProduceConsumeAll        bool     `json:"produceConsumeAll"`
	AdminTopicNames          []string `json:"adminTopicNames"`
	AdminAll                 bool     `json:"adminAll"`
	MtlsAuthen               bool     `json:"mtlsAuthen"`
	SaslAuthen               bool     `json:"saslAuthen"`
	PortalUserId             int32    `json:"portalUserId"`
	OldStatus                string   `json:"oldStatus"`
	Status                   string   `json:"status"`
	CreatedAt                string   `json:"createdAt"`
}

func (s *KafkaUserResponse) ToEntity() *lsentity.KafkaUser {
	if s == nil {
		return nil
	}
	return &lsentity.KafkaUser{
		Id:                       s.Id,
		ClusterId:                s.ClusterId,
		Name:                     s.Name,
		ProduceTopicNames:        s.ProduceTopicNames,
		ProduceAll:               s.ProduceAll,
		ConsumeTopicNames:        s.ConsumeTopicNames,
		ConsumeAll:               s.ConsumeAll,
		ProduceConsumeTopicNames: s.ProduceConsumeTopicNames,
		ProduceConsumeAll:        s.ProduceConsumeAll,
		AdminTopicNames:          s.AdminTopicNames,
		AdminAll:                 s.AdminAll,
		MtlsAuthen:               s.MtlsAuthen,
		SaslAuthen:               s.SaslAuthen,
		PortalUserId:             s.PortalUserId,
		OldStatus:                s.OldStatus,
		Status:                   s.Status,
		CreatedAt:                s.CreatedAt,
	}
}

type ListUsersResponse []KafkaUserResponse

func (s ListUsersResponse) ToEntity() []lsentity.KafkaUser {
	out := make([]lsentity.KafkaUser, 0, len(s))
	for i := range s {
		out = append(out, *s[i].ToEntity())
	}
	return out
}

// TopicDto ----------------------------------------------------------------

type KafkaTopicResponse struct {
	Id               string `json:"id"`
	ClusterId        string `json:"clusterId"`
	Name             string `json:"name"`
	Partitions       int32  `json:"partitions"`
	Replicas         int32  `json:"replicas"`
	RetentionSeconds int64  `json:"retentionSeconds"`
	RetentionBytes   int64  `json:"retentionBytes"`
	PortalUserId     int32  `json:"portalUserId"`
	Status           string `json:"status"`
	CreatedAt        string `json:"createdAt"`
}

func (s *KafkaTopicResponse) ToEntity() *lsentity.KafkaTopic {
	if s == nil {
		return nil
	}
	return &lsentity.KafkaTopic{
		Id:               s.Id,
		ClusterId:        s.ClusterId,
		Name:             s.Name,
		Partitions:       s.Partitions,
		Replicas:         s.Replicas,
		RetentionSeconds: s.RetentionSeconds,
		RetentionBytes:   s.RetentionBytes,
		PortalUserId:     s.PortalUserId,
		Status:           s.Status,
		CreatedAt:        s.CreatedAt,
	}
}

type ListTopicsResponse []KafkaTopicResponse

func (s ListTopicsResponse) ToEntity() []lsentity.KafkaTopic {
	out := make([]lsentity.KafkaTopic, 0, len(s))
	for i := range s {
		out = append(out, *s[i].ToEntity())
	}
	return out
}
