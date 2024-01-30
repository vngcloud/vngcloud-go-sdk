package cluster

import "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"

// ****************************************************** GetOpts ******************************************************

type GetOpts struct {
	ClusterCommonOpts
	common.CommonOpts
}

// ************************************************* UpdateSecgroupOpts ************************************************

type UpdateSecgroupOpts struct {
	ClusterID   string   `json:"clusterId"`
	Master      bool     `json:"master"`
	SecGroupIds []string `json:"secGroupIds"`

	common.CommonOpts
}

func (s *UpdateSecgroupOpts) ToRequestBody() interface{} {
	return s
}

func (s *UpdateSecgroupOpts) GetClusterID() string {
	return s.ClusterID
}
