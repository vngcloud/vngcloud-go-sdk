package v2

import (
	lfmt "fmt"
	ljparser "github.com/cuongpiger/joat/parser"
)

func NewCreateBlockVolumeRequest(pvolumeName, pvolumeType string, psize int64) ICreateBlockVolumeRequest {
	opt := new(CreateBlockVolumeRequest)
	opt.VolumeTypeId = pvolumeType
	opt.CreatedFrom = CreateFromNew
	opt.Name = pvolumeName
	opt.Size = psize

	return opt
}

func NewDeleteBlockVolumeByIdRequest(pvolumeId string) IDeleteBlockVolumeByIdRequest {
	opt := new(DeleteBlockVolumeByIdRequest)
	opt.BlockVolumeId = pvolumeId
	return opt
}

func NewListBlockVolumesRequest(ppage, psize int) IListBlockVolumesRequest {
	opt := new(ListBlockVolumesRequest)
	opt.Page = ppage
	opt.Size = psize
	return opt
}

const (
	CreateFromNew      = CreateVolumeFrom("NEW")
	CreateFromSnapshot = CreateVolumeFrom("SNAPSHOT")
)

type CreateBlockVolumeRequest struct {
	BackupVolumePointId    string                  `json:"backupVolumePointId,omitempty"`
	CreatedFrom            CreateVolumeFrom        `json:"createdFrom,omitempty"`
	EncryptionType         string                  `json:"encryptionType,omitempty"`
	MultiAttach            bool                    `json:"multiAttach,omitempty"`
	Name                   string                  `json:"name"`
	Size                   int64                   `json:"size"`
	VolumeTypeId           string                  `json:"volumeTypeId"`
	Tags                   []VolumeTag             `json:"tags,omitempty"`
	IsPoc                  bool                    `json:"isPoc,omitempty"`
	IsEnableAutoRenew      bool                    `json:"isEnableAutoRenew,omitempty"`
	ConfigureVolumeRestore *ConfigureVolumeRestore `json:"configVolumeRestore,omitempty"`
}

type DeleteBlockVolumeByIdRequest struct {
	BlockVolumeCommon
}

type ListBlockVolumesRequest struct {
	Name string `q:"name,beempty"`
	Page int    `q:"page"`
	Size int    `q:"size"`
}

type (
	CreateVolumeFrom string

	VolumeTag struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	ConfigureVolumeRestore struct {
		SnapshotVolumePointId string `json:"snapshotVolumePointId"`
		VolumeTypeId          string `json:"volumeTypeId"`
	}
)

func (s *CreateBlockVolumeRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateBlockVolumeRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"backupVolumePointId": s.BackupVolumePointId,
		"createdFrom":         s.CreatedFrom,
		"encryptionType":      s.EncryptionType,
		"multiAttach":         s.MultiAttach,
		"name":                s.Name,
		"size":                s.Size,
		"volumeTypeId":        s.VolumeTypeId,
		"tags":                s.Tags,
		"isPoc":               s.IsPoc,
		"isEnableAutoRenew":   s.IsEnableAutoRenew,
		"configVolumeRestore": s.ConfigureVolumeRestore,
	}
}

func (s *CreateBlockVolumeRequest) WithPoc(pisPoc bool) ICreateBlockVolumeRequest {
	s.IsPoc = pisPoc
	return s
}

func (s *CreateBlockVolumeRequest) WithAutoRenew(pval bool) ICreateBlockVolumeRequest {
	s.IsEnableAutoRenew = pval
	return s
}

func (s *CreateBlockVolumeRequest) WithMultiAttach(pmultiAttach bool) ICreateBlockVolumeRequest {
	s.MultiAttach = pmultiAttach
	return s
}

func (s *CreateBlockVolumeRequest) WithSize(psize int64) ICreateBlockVolumeRequest {
	s.Size = psize
	return s
}

func (s *CreateBlockVolumeRequest) WithVolumeType(pvolumeTypeId string) ICreateBlockVolumeRequest {
	s.VolumeTypeId = pvolumeTypeId
	return s
}

func (s *CreateBlockVolumeRequest) WithVolumeRestoreFromSnapshot(psnapshotID, pvolumeTypeID string) ICreateBlockVolumeRequest {
	s.CreatedFrom = CreateFromSnapshot
	s.ConfigureVolumeRestore = &ConfigureVolumeRestore{
		SnapshotVolumePointId: psnapshotID,
		VolumeTypeId:          pvolumeTypeID,
	}

	return s
}

func (s *CreateBlockVolumeRequest) WithTags(ptags ...string) ICreateBlockVolumeRequest {
	if s.Tags == nil {
		s.Tags = make([]VolumeTag, 0)
	}

	if len(ptags)%2 != 0 {
		ptags = append(ptags, "none")
	}

	for i := 0; i < len(ptags); i += 2 {
		s.Tags = append(s.Tags, VolumeTag{Key: ptags[i], Value: ptags[i+1]})
	}

	return s
}

func (s *ListBlockVolumesRequest) ToQuery() (string, error) {
	parser, _ := ljparser.GetParser()
	url, err := parser.UrlMe(s)

	if err != nil {
		return "", err
	}

	return url.String(), err
}

func (s *ListBlockVolumesRequest) GetDefaultQuery() string {
	return lfmt.Sprintf("page=%d&size=%d&name=", defaultPageListBlockVolumesRequest, defaultSizeListBlockVolumesRequest)
}

func (s *ListBlockVolumesRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name": s.Name,
		"page": s.Page,
		"size": s.Size,
	}
}

func (s *ListBlockVolumesRequest) WithName(pname string) IListBlockVolumesRequest {
	s.Name = pname
	return s
}
