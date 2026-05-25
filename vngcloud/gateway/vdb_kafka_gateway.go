package gateway

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lskafkaSvc "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/kafka"
)

var _ IVDBKafkaGateway = &vdbKafkaGateway{}

type vdbKafkaGateway struct {
	endpoint          string
	vdbKafkaGatewayV1 IVDBKafkaGatewayV1
}

func NewVDBKafkaGateway(pendpoint, puserId string, phc lsclient.IHttpClient) IVDBKafkaGateway {
	// vDB Gateway base URL: https://vdb-gateway.vngcloud.vn
	// Kafka APIs live under /vdb-kafka — there is no API version in the path
	// (single implicit v1).
	// puserId is the tenant's portal-user-id (integer as string) which the
	// gateway requires in every request header even with a valid Bearer token.
	kafkaSvc := lsclient.NewServiceClient().
		WithEndpoint(pendpoint + "vdb-kafka").
		WithClient(phc).
		WithUserId(puserId)

	return &vdbKafkaGateway{
		endpoint:          pendpoint,
		vdbKafkaGatewayV1: NewVDBKafkaGatewayV1(kafkaSvc),
	}
}

func (s *vdbKafkaGateway) V1() IVDBKafkaGatewayV1 {
	return s.vdbKafkaGatewayV1
}

func (s *vdbKafkaGateway) GetEndpoint() string {
	return s.endpoint
}

var _ IVDBKafkaGatewayV1 = &vdbKafkaGatewayV1{}

type vdbKafkaGatewayV1 struct {
	kafkaService lskafkaSvc.IKafkaServiceV1
}

func NewVDBKafkaGatewayV1(psvcClient lsclient.IServiceClient) IVDBKafkaGatewayV1 {
	return &vdbKafkaGatewayV1{
		kafkaService: lskafkaSvc.NewKafkaServiceV1(psvcClient),
	}
}

func (s *vdbKafkaGatewayV1) KafkaService() lskafkaSvc.IKafkaServiceV1 {
	return s.kafkaService
}
