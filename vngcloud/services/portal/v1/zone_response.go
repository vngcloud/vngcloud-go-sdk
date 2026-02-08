package v1

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
)

type Zone struct {
	Uuid          string `json:"uuid"`
	Name          string `json:"name"`
	OpenstackZone string `json:"openstackZone"`
}

type ListZoneResponse struct {
	Data []Zone `json:"data"`
}

func (s *ListZoneResponse) ToEntityListZones() *lsentity.ListZones {
	listZones := &lsentity.ListZones{
		Items: make([]*lsentity.Zone, 0),
	}
	for _, q := range s.Data {
		listZones.Items = append(listZones.Items, q.ToEntityZone())
	}

	return listZones
}

func (s *Zone) ToEntityZone() *lsentity.Zone {

	return &lsentity.Zone{
		Uuid:          s.Uuid,
		Name:          s.Name,
		OpenstackZone: s.OpenstackZone,
	}
}
