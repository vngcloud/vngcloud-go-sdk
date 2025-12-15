package v1

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
)

type GetHostedZoneByIdResponse struct {
	Data *lsentity.HostedZone `json:"data"`
}

func (r *GetHostedZoneByIdResponse) ToEntityHostedZone() *lsentity.HostedZone {
	return r.Data
}