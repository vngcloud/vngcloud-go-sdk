package volume

import (
	lfmt "fmt"
	lParser "github.com/cuongpiger/joat/parser"

	bstCm "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/blockstorage/v2"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
)

// ************************************************** ListOptsBuilder **************************************************

type ListOpts struct {
	Name string `q:"name,beempty"`
	Page int    `q:"page"`
	Size int    `q:"size"`

	common.CommonOpts
}

func (s *ListOpts) ToListQuery() (string, error) {
	parser, _ := lParser.GetParser()
	url, err := parser.UrlMe(s)

	if err != nil {
		return "", err
	}

	return url.String(), err
}

func (s *ListOpts) GetDefaultQuery() string {
	return lfmt.Sprintf("page=%d&size=%d&name=", defaultPageListVolume, defaultSizeListVolume)
}

// ************************************************* CreateOptsBuilder *************************************************

type CreateOpts struct {
	BackupVolumePointId    string                  `json:"backupVolumePointId,omitempty"`
	CreatedFrom            CreateOptsCreateFrom    `json:"createdFrom,omitempty"`
	EncryptionType         string                  `json:"encryptionType,omitempty"`
	MultiAttach            bool                    `json:"multiAttach,omitempty"`
	Name                   string                  `json:"name"`
	Size                   uint64                  `json:"size"`
	VolumeTypeId           string                  `json:"volumeTypeId"`
	Tags                   []CreateOptsTag         `json:"tags,omitempty"`
	IsPoc                  bool                    `json:"isPoc,omitempty"`
	IsEnableAutoRenew      bool                    `json:"isEnableAutoRenew,omitempty"`
	ConfigureVolumeRestore *ConfigureVolumeRestore `json:"configVolumeRestore,omitempty"`

	common.CommonOpts
}

type CreateOptsCreateFrom string

type ConfigureVolumeRestore struct {
	SnapshotVolumePointId string `json:"snapshotVolumePointId"`
	VolumeTypeId          string `json:"volumeTypeId"`
}

const (
	CreateFromNew      CreateOptsCreateFrom = "NEW"
	CreateFromSnapshot CreateOptsCreateFrom = "SNAPSHOT"
)

type CreateOptsTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (s *CreateOpts) ToRequestBody() interface{} {
	return s
}

// *************************************************** GetOptsBuilder **************************************************

type GetOpts struct {
	common.CommonOpts
	bstCm.BlockStorageV2Common
}

// ************************************************* DeleteOptsBuilder *************************************************

type DeleteOpts struct {
	common.CommonOpts
	bstCm.BlockStorageV2Common
}

// ************************************************* ListAllOptsBuilder ************************************************

type ListAllOpts struct {
	common.CommonOpts
}

func (s *ListAllOpts) ToListQuery() (string, error) {
	parser, _ := lParser.GetParser()
	url, err := parser.UrlMe(s)

	if err != nil {
		return "", err
	}

	return url.String(), err
}
