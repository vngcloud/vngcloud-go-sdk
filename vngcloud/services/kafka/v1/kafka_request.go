package v1

import (
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
)

// ListClustersRequest -----------------------------------------------------

type ListClustersRequest struct {
	lscommon.UserAgent
}

func NewListClustersRequest() IListClustersRequest {
	return &ListClustersRequest{}
}

// GetClusterRequest -------------------------------------------------------

type GetClusterRequest struct {
	ClusterId string
	lscommon.UserAgent
}

func NewGetClusterRequest(pclusterId string) IGetClusterRequest {
	return &GetClusterRequest{ClusterId: pclusterId}
}

func (s *GetClusterRequest) GetClusterId() string {
	return s.ClusterId
}

// ListUsersRequest --------------------------------------------------------

type ListUsersRequest struct {
	ClusterId string
	lscommon.UserAgent
}

func NewListUsersRequest(pclusterId string) IListUsersRequest {
	return &ListUsersRequest{ClusterId: pclusterId}
}

func (s *ListUsersRequest) GetClusterId() string {
	return s.ClusterId
}

// GetUserAuthenCredsRequest -----------------------------------------------

type GetUserAuthenCredsRequest struct {
	ClusterId string
	UserId    string
	lscommon.UserAgent
}

func NewGetUserAuthenCredsRequest(pclusterId, puserId string) IGetUserAuthenCredsRequest {
	return &GetUserAuthenCredsRequest{ClusterId: pclusterId, UserId: puserId}
}

func (s *GetUserAuthenCredsRequest) GetClusterId() string {
	return s.ClusterId
}

func (s *GetUserAuthenCredsRequest) GetUserId() string {
	return s.UserId
}

// ListTopicsRequest -------------------------------------------------------

type ListTopicsRequest struct {
	ClusterId string
	lscommon.UserAgent
}

func NewListTopicsRequest(pclusterId string) IListTopicsRequest {
	return &ListTopicsRequest{ClusterId: pclusterId}
}

func (s *ListTopicsRequest) GetClusterId() string {
	return s.ClusterId
}

// CreateTopicRequest ------------------------------------------------------

// CreateTopicRequest maps to TopicCreateRequest in the vDB OpenAPI spec.
// Constraints (per spec): Partitions 1-2048; Replicas ≤ broker count;
// RetentionSeconds 3600-7776000 (1h-90d) or empty for cluster default;
// RetentionBytes 1-1099511627776 (1 TiB) or -1 for unlimited or empty for default.
// ClusterId is a path param — exposed via ToRequestBody() returning a body view that strips it.
type CreateTopicRequest struct {
	ClusterId        string `json:"-"`
	Name             string `json:"name"`
	Partitions       int32  `json:"partitions,omitempty"`
	Replicas         int32  `json:"replicas,omitempty"`
	RetentionSeconds int64  `json:"retentionSeconds,omitempty"`
	RetentionBytes   int64  `json:"retentionBytes,omitempty"`
	lscommon.UserAgent
}

// createTopicBody is the wire-level body shape — no embedded UserAgent so that
// json marshalling stays clean ({name, partitions, replicas, retentionSeconds, retentionBytes}).
type createTopicBody struct {
	Name             string `json:"name"`
	Partitions       int32  `json:"partitions,omitempty"`
	Replicas         int32  `json:"replicas,omitempty"`
	RetentionSeconds int64  `json:"retentionSeconds,omitempty"`
	RetentionBytes   int64  `json:"retentionBytes,omitempty"`
}

func NewCreateTopicRequest(pclusterId string) ICreateTopicRequest {
	return &CreateTopicRequest{ClusterId: pclusterId}
}

func (s *CreateTopicRequest) GetClusterId() string {
	return s.ClusterId
}

func (s *CreateTopicRequest) GetName() string {
	return s.Name
}

func (s *CreateTopicRequest) ToRequestBody() interface{} {
	return &createTopicBody{
		Name:             s.Name,
		Partitions:       s.Partitions,
		Replicas:         s.Replicas,
		RetentionSeconds: s.RetentionSeconds,
		RetentionBytes:   s.RetentionBytes,
	}
}

func (s *CreateTopicRequest) WithName(pname string) ICreateTopicRequest {
	s.Name = pname
	return s
}

func (s *CreateTopicRequest) WithPartitions(ppartitions int32) ICreateTopicRequest {
	s.Partitions = ppartitions
	return s
}

func (s *CreateTopicRequest) WithReplicas(preplicas int32) ICreateTopicRequest {
	s.Replicas = preplicas
	return s
}

func (s *CreateTopicRequest) WithRetentionSeconds(pretentionSeconds int64) ICreateTopicRequest {
	s.RetentionSeconds = pretentionSeconds
	return s
}

func (s *CreateTopicRequest) WithRetentionBytes(pretentionBytes int64) ICreateTopicRequest {
	s.RetentionBytes = pretentionBytes
	return s
}
