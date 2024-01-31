package server

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
	lServerV2 "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/compute/v2"
)

// ************************************************** GetOptsBuilder ***************************************************

type GetOpts struct {
	common.CommonOpts
	lServerV2.ServerV2Common
}

// ************************************************* DeleteOptsBuilder *************************************************

type DeleteOpts struct {
	DeleteAllVolume bool `json:"deleteAllVolume"`

	common.CommonOpts
	lServerV2.ServerV2Common
}

func (s *DeleteOpts) ToRequestBody() interface{} {
	return s
}
