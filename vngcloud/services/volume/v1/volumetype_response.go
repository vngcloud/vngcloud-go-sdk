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
	VolumeTypeZone struct {
		Id       string   `json:"id"`
		Name     string   `json:"name"`
		PoolName []string `json:"poolName,omitempty"`
	}

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

type ListVolumeTypeZonesResponse struct {
	VolumeTypeZones []VolumeTypeZone `json:"volumeTypeZones"`
}

type ListVolumeTypeResponse struct {
	VolumeTypes []VolumeType `json:"volumeTypes"`
}

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

func (s VolumeTypeZone) toEntityVolumeTypeZone() *lsentity.VolumeTypeZone {
	return &lsentity.VolumeTypeZone{
		Id:       s.Id,
		Name:     s.Name,
		PoolName: s.PoolName,
	}
}

func (s *ListVolumeTypeZonesResponse) ToEntityListVolumeTypeZones() *lsentity.ListVolumeTypeZones {
	sl := new(lsentity.ListVolumeTypeZones)

	for _, item := range s.VolumeTypeZones {
		sl.VolumeTypeZones = append(sl.VolumeTypeZones, item.toEntityVolumeTypeZone())
	}

	return sl
}

func (s *ListVolumeTypeResponse) ToEntityListVolumeType() *lsentity.ListVolumeType {
	sl := new(lsentity.ListVolumeType)

	for _, item := range s.VolumeTypes {
		sl.VolumeTypes = append(sl.VolumeTypes, item.toEntityVolumeType())
	}

	return sl
}
