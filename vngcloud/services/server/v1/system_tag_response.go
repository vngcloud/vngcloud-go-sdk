package v1

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type SystemTagResponse struct {
}

func (s *SystemTagResponse) toSystemTag() *lsentity.SystemTag {
	return &lsentity.SystemTag{}
}

type CreateSystemTagResponse struct {
	Data SystemTagResponse `json:"data"`
}

func (s *CreateSystemTagResponse) ToAddressPair() *lsentity.SystemTag {
	return s.Data.toSystemTag()
}
