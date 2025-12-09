package v1

import (
	lfmt "fmt"

	ljparser "github.com/cuongpiger/joat/parser"
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
)

func NewGetVolumeTypeByIdRequest(pvolumeTypeId string) IGetVolumeTypeByIdRequest {
	opt := new(GetVolumeTypeByIdRequest)
	opt.VolumeTypeId = pvolumeTypeId
	return opt
}

func NewListVolumeTypeRequest(volumeTypeZoneId string) IGetListVolumeTypeRequest {
	opt := new(GetListVolumeTypeRequest)
	opt.VolumeTypeZoneId = volumeTypeZoneId
	return opt
}

func NewGetVolumeTypeZonesRequest(zoneId string) IGetVolumeTypeZonesRequest {
	opt := new(GetVolumeTypeZonesRequest)
	opt.ZoneId = zoneId
	return opt
}

type GetVolumeTypeByIdRequest struct {
	lscommon.VolumeTypeCommon
}

type GetVolumeTypeZonesRequest struct {
	ZoneId string `q:"zoneId"`
}

type GetListVolumeTypeRequest struct {
	VolumeTypeZoneId string
}

func (s *GetVolumeTypeZonesRequest) GetDefaultQuery() string {
	return lfmt.Sprintf("zoneId=%s", defaultZoneGetVolumeTypeZonesRequest)
}

func (s *GetVolumeTypeZonesRequest) ToQuery() (string, error) {
	parser, _ := ljparser.GetParser()
	url, err := parser.UrlMe(s)

	if err != nil {
		return "", err
	}

	return url.String(), err
}

func (s *GetListVolumeTypeRequest) GetVolumeTypeZoneId() string {
	return s.VolumeTypeZoneId
}
