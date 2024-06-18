package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type CreatePoolResponse struct {
	UUID string `json:"uuid"`
}

func (s *CreatePoolResponse) ToEntityPool() *lsentity.Pool {
	return &lsentity.Pool{
		UUID: s.UUID,
	}
}
