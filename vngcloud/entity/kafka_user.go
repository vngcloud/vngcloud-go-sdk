package entity

const (
	KafkaUserStatusActive   = "ACTIVE"
	KafkaUserStatusCreating = "CREATING"
	KafkaUserStatusError    = "ERROR"
)

type KafkaUser struct {
	Id                       string
	ClusterId                string
	Name                     string
	ProduceTopicNames        []string
	ProduceAll               bool
	ConsumeTopicNames        []string
	ConsumeAll               bool
	ProduceConsumeTopicNames []string
	ProduceConsumeAll        bool
	AdminTopicNames          []string
	AdminAll                 bool
	MtlsAuthen               bool
	SaslAuthen               bool
	PortalUserId             int32
	OldStatus                string
	Status                   string
	CreatedAt                string
}

func (s *KafkaUser) CanProduceTo(topic string) bool {
	if s.ProduceAll || s.ProduceConsumeAll || s.AdminAll {
		return true
	}
	for _, t := range s.ProduceTopicNames {
		if t == topic {
			return true
		}
	}
	for _, t := range s.ProduceConsumeTopicNames {
		if t == topic {
			return true
		}
	}
	for _, t := range s.AdminTopicNames {
		if t == topic {
			return true
		}
	}
	return false
}
