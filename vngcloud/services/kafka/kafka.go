package kafka

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lskafkaSvcV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/kafka/v1"
)

func NewKafkaServiceV1(psvcClient lsclient.IServiceClient) IKafkaServiceV1 {
	return &lskafkaSvcV1.KafkaServiceV1{
		VkafkaClient: psvcClient,
	}
}
