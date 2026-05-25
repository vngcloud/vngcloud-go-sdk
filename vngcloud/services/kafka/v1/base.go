package v1

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
)

type KafkaServiceV1 struct {
	VkafkaClient lsclient.IServiceClient
}
