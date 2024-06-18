package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type CreateListenerResponse struct {
	UUID string `json:"uuid"`
}

func (s *CreateListenerResponse) ToEntityListener() *lsentity.Listener {
	return &lsentity.Listener{
		UUID: s.UUID,
	}
}
