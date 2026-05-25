package entity

const (
	KafkaTopicStatusActive   = "ACTIVE"
	KafkaTopicStatusCreating = "CREATING"
	KafkaTopicStatusError    = "ERROR"
)

type KafkaTopic struct {
	Id               string
	ClusterId        string
	Name             string
	Partitions       int32
	Replicas         int32
	RetentionSeconds int64
	RetentionBytes   int64
	PortalUserId     int32
	Status           string
	CreatedAt        string
}

func (s *KafkaTopic) IsActive() bool {
	return s.Status == KafkaTopicStatusActive
}
