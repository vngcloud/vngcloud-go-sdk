package v2

import (
	lfmt "fmt"
	ljparser "github.com/cuongpiger/joat/parser"

	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
)

func NewListSnapshotsByBlockVolumeIdRequest(ppage, psize int, pblockVolumeId string) IListSnapshotsByBlockVolumeIdRequest {
	opt := new(ListSnapshotsByBlockVolumeIdRequest)
	opt.BlockVolumeId = pblockVolumeId
	opt.Page = ppage
	opt.Size = psize

	return opt
}

func NewCreateSnapshotByVolumeIdRequest(pname, pblockVolumeId string) ICreateSnapshotByVolumeIdRequest {
	opt := new(CreateSnapshotByVolumeIdRequest)
	opt.Name = pname
	opt.BlockVolumeId = pblockVolumeId

	return opt
}

func NewDeleteSnapshotByIdRequest(psnapshotId string) IDeleteSnapshotByIdRequest {
	opt := new(DeleteSnapshotByIdRequest)
	opt.BlockVolumeId = "undefined"
	opt.SnapshotId = psnapshotId

	return opt
}

type ListSnapshotsByBlockVolumeIdRequest struct {
	Page int `q:"page"`
	Size int `q:"size"`

	lscommon.BlockVolumeCommon
}

type CreateSnapshotByVolumeIdRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Permanently bool   `json:"isPermanently"`
	RetainedDay uint64 `json:"retainedDay"`

	lscommon.BlockVolumeCommon
}

type DeleteSnapshotByIdRequest struct {
	lscommon.BlockVolumeCommon
	lscommon.SnapshotCommon
}

func (s *ListSnapshotsByBlockVolumeIdRequest) GetDefaultQuery() string {
	return lfmt.Sprintf("page=%d&size=%d", defaultPageListSnapshotsByBlockVolumeIdRequest, defaultSizeListSnapshotsByBlockVolumeIdRequest)
}

func (s *ListSnapshotsByBlockVolumeIdRequest) ToQuery() (string, error) {
	parser, _ := ljparser.GetParser()
	url, err := parser.UrlMe(s)

	if err != nil {
		return "", err
	}

	return url.String(), err
}

func (s *CreateSnapshotByVolumeIdRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateSnapshotByVolumeIdRequest) WithDescription(pdesc string) ICreateSnapshotByVolumeIdRequest {
	s.Description = pdesc
	return s
}

func (s *CreateSnapshotByVolumeIdRequest) WithPermanently(pval bool) ICreateSnapshotByVolumeIdRequest {
	s.Permanently = pval
	return s
}

func (s *CreateSnapshotByVolumeIdRequest) WithRetainedDay(pval uint64) ICreateSnapshotByVolumeIdRequest {
	s.RetainedDay = pval
	return s
}
