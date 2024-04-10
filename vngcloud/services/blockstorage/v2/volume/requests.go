package volume

import (
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

func (s *ListOpts) ToListQueryWithParams(pParams *map[string]interface{}) (string, error) {
	if pParams != nil {
		if value, ok := (*pParams)["name"]; ok {
			s.Name = value.(string)
		}

		if value, ok := (*pParams)["page"]; ok {
			s.Page = value.(int)
		}

		if value, ok := (*pParams)["size"]; ok {
			s.Size = value.(int)
		}

		parser, _ := lParser.GetParser()
		url, err := parser.UrlMe(s)
		if err != nil {
			return "", err
		}

		return url.String(), err
	}

	return "", nil
}

// ************************************************* CreateOptsBuilder *************************************************

type CreateOpts struct {
	BackupVolumePointId string               `json:"backupVolumePointId,omitempty"`
	CreatedFrom         CreateOptsCreateFrom `json:"createdFrom,omitempty"`
	EncryptionType      string               `json:"encryptionType,omitempty"`
	MultiAttach         bool                 `json:"multiAttach,omitempty"`
	Name                string               `json:"name"`
	Size                uint64               `json:"size"`
	VolumeTypeId        string               `json:"volumeTypeId"`
	Tags                []CreateOptsTag      `json:"tags,omitempty"`
	IsPoc               bool                 `json:"isPoc,omitempty"`
	IsEnableAutoRenew   bool                 `json:"isEnableAutoRenew,omitempty"`

	common.CommonOpts
}

type CreateOptsCreateFrom string

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
