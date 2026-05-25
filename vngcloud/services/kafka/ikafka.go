package kafka

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lskafkaSvcV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/kafka/v1"
)

type IKafkaServiceV1 interface {
	ListClusters(popts lskafkaSvcV1.IListClustersRequest) (*[]lsentity.KafkaCluster, lserr.IError)
	GetCluster(popts lskafkaSvcV1.IGetClusterRequest) (*lsentity.KafkaCluster, lserr.IError)
	ListUsers(popts lskafkaSvcV1.IListUsersRequest) (*[]lsentity.KafkaUser, lserr.IError)
	GetUserAuthenCreds(popts lskafkaSvcV1.IGetUserAuthenCredsRequest) ([]byte, lserr.IError)
	ListTopics(popts lskafkaSvcV1.IListTopicsRequest) (*[]lsentity.KafkaTopic, lserr.IError)
	CreateTopic(popts lskafkaSvcV1.ICreateTopicRequest) (*lsentity.KafkaTopic, lserr.IError)
}
