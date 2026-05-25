package v1

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

// ListClusters returns every Kafka cluster the authenticated portal user owns.
func (s *KafkaServiceV1) ListClusters(popts IListClustersRequest) (*[]lsentity.KafkaCluster, lserr.IError) {
	url := listClustersUrl(s.VkafkaClient)
	resp := new(ListClustersResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithUserId(s.VkafkaClient.GetUserId()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VkafkaClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp)
	}

	out := resp.ToEntity()
	return &out, nil
}

// GetCluster fetches detail of a specific Kafka cluster.
func (s *KafkaServiceV1) GetCluster(popts IGetClusterRequest) (*lsentity.KafkaCluster, lserr.IError) {
	url := getClusterUrl(s.VkafkaClient, popts)
	resp := new(KafkaClusterResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithUserId(s.VkafkaClient.GetUserId()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VkafkaClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("clusterId", popts.GetClusterId())
	}

	return resp.ToEntity(), nil
}

// ListUsers returns Kafka users registered on the given cluster.
func (s *KafkaServiceV1) ListUsers(popts IListUsersRequest) (*[]lsentity.KafkaUser, lserr.IError) {
	url := listUsersUrl(s.VkafkaClient, popts)
	resp := new(ListUsersResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithUserId(s.VkafkaClient.GetUserId()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VkafkaClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("clusterId", popts.GetClusterId())
	}

	out := resp.ToEntity()
	return &out, nil
}

// GetUserAuthenCreds downloads the authen-creds bundle as a raw ZIP byte slice.
// The zip contains both mtls/ and sasl/ subfolders — use ParseAuthenCredsZip to extract PEM bytes.
func (s *KafkaServiceV1) GetUserAuthenCreds(popts IGetUserAuthenCredsRequest) ([]byte, lserr.IError) {
	url := getUserAuthenCredsUrl(s.VkafkaClient, popts)
	var zipBytes []byte
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithUserId(s.VkafkaClient.GetUserId()).
		WithOkCodes(200).
		WithBytesResponse(&zipBytes).
		WithJsonError(errResp)

	if _, sdkErr := s.VkafkaClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"clusterId", popts.GetClusterId(),
				"userId", popts.GetUserId(),
			)
	}

	return zipBytes, nil
}

// ListTopics returns the topics present on the cluster.
func (s *KafkaServiceV1) ListTopics(popts IListTopicsRequest) (*[]lsentity.KafkaTopic, lserr.IError) {
	url := listTopicsUrl(s.VkafkaClient, popts)
	resp := new(ListTopicsResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithUserId(s.VkafkaClient.GetUserId()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VkafkaClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("clusterId", popts.GetClusterId())
	}

	out := resp.ToEntity()
	return &out, nil
}

// CreateTopic provisions a new topic on the cluster. Returns the TopicDto.
// Worker uses this for lazy create when sink=Kafka.
func (s *KafkaServiceV1) CreateTopic(popts ICreateTopicRequest) (*lsentity.KafkaTopic, lserr.IError) {
	url := createTopicUrl(s.VkafkaClient, popts)
	resp := new(KafkaTopicResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithUserId(s.VkafkaClient.GetUserId()).
		WithOkCodes(200, 202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VkafkaClient.Post(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"clusterId", popts.GetClusterId(),
				"topicName", popts.GetName(),
			)
	}

	return resp.ToEntity(), nil
}
