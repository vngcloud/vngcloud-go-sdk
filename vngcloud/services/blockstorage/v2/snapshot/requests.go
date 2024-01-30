package snapshot

import (
	bsCm "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/blockstorage/v2"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
)

type CreateOpts struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Permanently bool   `json:"isPermanently"`
	RetainedDay uint64 `json:"retainedDay"`

	common.CommonOpts
	bsCm.BlockStorageV2Common
}

func (s *CreateOpts) ToRequestBody() interface{} {
	return s
}

// ******************************************************* Delete ******************************************************

type DeleteOpts struct {
	SnapshotID string
	common.CommonOpts
	bsCm.BlockStorageV2Common
}

func (s *DeleteOpts) GetSnapshotID() string {
	return s.SnapshotID
}
