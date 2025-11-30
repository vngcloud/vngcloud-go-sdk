package v2

import (
	lfmt "fmt"

	ljparser "github.com/cuongpiger/joat/parser"
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
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

func NewGetBlockVolumeByIdRequest(pvolumeId string) IGetBlockVolumeByIdRequest {
	opt := new(GetBlockVolumeByIdRequest)
	opt.BlockVolumeId = pvolumeId
	return opt
}

func NewResizeBlockVolumeByIdRequest(pvolumeId, pvolumeType string, psize int) IResizeBlockVolumeByIdRequest {
	opt := new(ResizeBlockVolumeByIdRequest)
	opt.BlockVolumeId = pvolumeId
	opt.NewSize = psize
	opt.VolumeTypeID = pvolumeType
	return opt
}

func NewGetUnderVolumeIdRequest(pvolumeId string) IGetUnderBlockVolumeIdRequest {
	opt := new(GetUnderBlockVolumeIdRequest)
	opt.BlockVolumeId = pvolumeId
	return opt
}

func NewMigrateBlockVolumeByIdRequest(pvolumeId, pvolumeType string) IMigrateBlockVolumeByIdRequest {
	opt := new(MigrateBlockVolumeByIdRequest)
	opt.BlockVolumeId = pvolumeId
	opt.VolumeTypeId = pvolumeType
	opt.Action = InitMigrateAction
	return opt
}

const (
	CreateFromNew      = CreateVolumeFrom("NEW")
	CreateFromSnapshot = CreateVolumeFrom("SNAPSHOT")

	AesXtsPlain64_128 = EncryptType("aes-xts-plain64_128")
	AesXtsPlain64_256 = EncryptType("aes-xts-plain64_256")

	InitMigrateAction    = MigrateAction("INIT-MIGRATE")
	ProcessMigrateAction = MigrateAction("MIGRATE")
	ConfirmMigrateAction = MigrateAction("CONFIRM-MIGRATE")
)

type CreateBlockVolumeRequest struct {
	BackupVolumePointId    string                  `json:"backupVolumePointId,omitempty"`
	CreatedFrom            CreateVolumeFrom        `json:"createdFrom,omitempty"`
	EncryptionType         EncryptType             `json:"encryptionType,omitempty"`
	MultiAttach            bool                    `json:"multiAttach,omitempty"`
	Name                   string                  `json:"name"`
	Size                   int64                   `json:"size"`
	VolumeTypeId           string                  `json:"volumeTypeId"`
	Tags                   []VolumeTag             `json:"tags,omitempty"`
	IsPoc                  bool                    `json:"isPoc,omitempty"`
	IsEnableAutoRenew      bool                    `json:"isEnableAutoRenew,omitempty"`
	ConfigureVolumeRestore *ConfigureVolumeRestore `json:"configVolumeRestore,omitempty"`
	Zone                   string                  `json:"zoneId,omitempty"`
	PoolName               string                  `json:"poolName,omitempty"`
}

type DeleteBlockVolumeByIdRequest struct {
	lscommon.BlockVolumeCommon
}

type ResizeBlockVolumeByIdRequest struct {
	NewSize      int    `json:"newSize"`         // NewSize is the new size of the volume, in GB
	VolumeTypeID string `json:"newVolumeTypeId"` // VolumeTypeID is the type of the volume
	lscommon.BlockVolumeCommon
}

type ListBlockVolumesRequest struct {
	Name string `q:"name,beempty"`
	Page int    `q:"page"`
	Size int    `q:"size"`
}

type AttachBlockVolumeRequest struct {
	lscommon.BlockVolumeCommon
	lscommon.ServerCommon
}

type GetBlockVolumeByIdRequest struct {
	lscommon.BlockVolumeCommon
}

type GetUnderBlockVolumeIdRequest struct {
	lscommon.BlockVolumeCommon
}

type MigrateBlockVolumeByIdRequest struct {
	Action         MigrateAction `json:"action"`
	ConfirmMigrate bool
	Tags           []lscommon.Tag `json:"tags"`
	VolumeTypeId   string         `json:"volumeTypeId"`
	Auto           bool

	lscommon.BlockVolumeCommon
}

type (
	MigrateAction    string
	CreateVolumeFrom string
	EncryptType      string

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

func (s *CreateBlockVolumeRequest) WithZone(pzone string) ICreateBlockVolumeRequest {
	s.Zone = pzone
	return s
}

func (s *CreateBlockVolumeRequest) WithPoolName(poolName string) ICreateBlockVolumeRequest {
	s.PoolName = poolName
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

func (s *CreateBlockVolumeRequest) GetListParameters() []interface{} {
	return []interface{}{
		"backupVolumePointId", s.BackupVolumePointId,
		"createdFrom", s.CreatedFrom,
		"encryptionType", s.EncryptionType,
		"multiAttach", s.MultiAttach,
		"name", s.Name,
		"size", s.Size,
		"volumeTypeId", s.VolumeTypeId,
		"tags", s.Tags,
		"isPoc", s.IsPoc,
		"isEnableAutoRenew", s.IsEnableAutoRenew,
		"configVolumeRestore", s.ConfigureVolumeRestore,
	}
}

func (s *CreateBlockVolumeRequest) GetVolumeName() string {
	return s.Name
}

func (s *CreateBlockVolumeRequest) GetVolumeType() string {
	return s.VolumeTypeId
}

func (s *CreateBlockVolumeRequest) GetZone() string {
	return s.Zone
}

func (s *CreateBlockVolumeRequest) GetPoolName() string {
	return s.PoolName
}

func (s *CreateBlockVolumeRequest) GetSize() int64 {
	return s.Size
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

func (s *CreateBlockVolumeRequest) WithEncryptionType(pet EncryptType) ICreateBlockVolumeRequest {
	s.EncryptionType = pet
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

func (s *ResizeBlockVolumeByIdRequest) ToRequestBody() interface{} {
	return s
}

func (s *ResizeBlockVolumeByIdRequest) GetSize() int {
	return s.NewSize
}

func (s *ResizeBlockVolumeByIdRequest) GetVolumeTypeId() string {
	return s.VolumeTypeID
}

func (s *MigrateBlockVolumeByIdRequest) ToRequestBody() interface{} {
	return s
}

func (s *MigrateBlockVolumeByIdRequest) WithTags(ptags ...string) IMigrateBlockVolumeByIdRequest {
	if s.Tags == nil {
		s.Tags = make([]lscommon.Tag, 0)
	}

	if len(ptags)%2 != 0 {
		ptags = append(ptags, "none")
	}

	for i := 0; i < len(ptags); i += 2 {
		s.Tags = append(s.Tags, lscommon.Tag{Key: ptags[i], Value: ptags[i+1]})
	}

	return s
}

func (s *MigrateBlockVolumeByIdRequest) WithAction(paction MigrateAction) IMigrateBlockVolumeByIdRequest {
	switch paction {
	case InitMigrateAction, ProcessMigrateAction, ConfirmMigrateAction:
		s.Action = paction
	default:
		s.Action = InitMigrateAction
	}

	return s
}

func (s *MigrateBlockVolumeByIdRequest) WithConfirm(pconfirm bool) IMigrateBlockVolumeByIdRequest {
	s.ConfirmMigrate = pconfirm
	return s
}

func (s *MigrateBlockVolumeByIdRequest) IsConfirm() bool {
	return s.ConfirmMigrate
}
