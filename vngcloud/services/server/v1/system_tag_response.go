package v1

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type SystemTagResponse struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	CreatedAt string `json:"createdAt"`
	SystemTag bool   `json:"systemTag"`
}

func (s *SystemTagResponse) toSystemTag() lsentity.SystemTag {
	return lsentity.SystemTag{
		Key:       s.Key,
		Value:     s.Value,
		CreatedAt: s.CreatedAt,
		SystemTag: s.SystemTag,
	}
}
