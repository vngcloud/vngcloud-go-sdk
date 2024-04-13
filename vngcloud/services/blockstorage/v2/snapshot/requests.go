package snapshot

import (
	lfmt "fmt"
	ljparser "github.com/cuongpiger/joat/parser"

	lsbsV2Cm "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/blockstorage/v2"
	lssvCm "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
)

/**
 * Define the request for the API CREATE Snapshot
 */

type CreateOpts struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Permanently bool   `json:"isPermanently"`
	RetainedDay uint64 `json:"retainedDay"`

	lssvCm.CommonOpts
	lsbsV2Cm.BlockStorageV2Common
}

func (s *CreateOpts) ToRequestBody() interface{} {
	return s
}

/**
 * Define the request for the API DELETE Snapshot
 */

type DeleteOpts struct {
	SnapshotID string
	lssvCm.CommonOpts
	lsbsV2Cm.BlockStorageV2Common
}

func (s *DeleteOpts) GetSnapshotID() string {
	return s.SnapshotID
}

/**
 * Define the request for the API List Snapshot
 */

type ListVolumeOpts struct {
	VolumeID string `q:"volumeId"`
	Page     int    `q:"page"`
	Size     int    `q:"size"`

	lssvCm.CommonOpts
}

// ToListQuery formats a ListOpts into a query string
func (s *ListVolumeOpts) ToListQuery() (string, error) {
	parser, _ := ljparser.GetParser()
	url, err := parser.UrlMe(s)
	if err != nil {
		return "", err
	}

	return url.String(), err
}

func (s *ListVolumeOpts) GetDefaultQuery() string {
	return lfmt.Sprintf("page=%d&size=%d", defaultPageListVolumeSnapshot, defaultSizeListVolumeSnapshot)
}
