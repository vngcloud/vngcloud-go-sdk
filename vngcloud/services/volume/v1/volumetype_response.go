package v1

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type GetVolumeTypeByIdResponse struct {
	VolumeTypes []VolumeType `json:"volumeTypes"`
}

type GetDefaultVolumeTypeResponse struct {
	Id     string `json:"volumeTypeId"`
	ZoneId string `json:"volumeTypeZoneId"`
}

type (
	VolumeType struct {
		Id         string `json:"id"`
		Iops       int    `json:"iops"`
		MaxSize    int    `json:"maxSize"`
		MinSize    int    `json:"minSize"`
		Name       string `json:"name"`
		ThroughPut int    `json:"throughPut,omitempty"`
		ZoneId     string `json:"zoneId,omitempty"`
	}
)

func (s *GetVolumeTypeByIdResponse) ToEntityVolumeType() *lsentity.VolumeType {
	if len(s.VolumeTypes) == 0 {
		return nil
	}

	return s.VolumeTypes[0].toEntityVolumeType()
}

func (s *GetDefaultVolumeTypeResponse) ToEntityVolumeType() *lsentity.VolumeType {
	return &lsentity.VolumeType{
		Id:     s.Id,
		ZoneId: s.ZoneId,
	}
}

func (s VolumeType) toEntityVolumeType() *lsentity.VolumeType {
	return &lsentity.VolumeType{
		Id:         s.Id,
		Iops:       s.Iops,
		MaxSize:    s.MaxSize,
		MinSize:    s.MinSize,
		Name:       s.Name,
		ThroughPut: s.ThroughPut,
		ZoneId:     s.ZoneId,
	}
}
