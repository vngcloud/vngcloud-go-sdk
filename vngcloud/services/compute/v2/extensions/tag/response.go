package tag

import "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"

type CreateResponse struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	CreatedAt string `json:"createdAt"`
	SystemTag bool   `json:"systemTag"`
}

func (s *CreateResponse) ToServerTagObject() *objects.ServerTag {
	if s == nil {
		return nil
	}

	return &objects.ServerTag{
		Key:       s.Key,
		Value:     s.Value,
		CreatedAt: s.CreatedAt,
		SystemTag: s.SystemTag,
	}
}

type GetResponse struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	CreatedAt string `json:"createdAt"`
	SystemTag bool   `json:"systemTag"`
}

func (s *GetResponse) ToServerTagObject() *objects.ServerTag {
	if s == nil {
		return nil
	}

	return &objects.ServerTag{
		Key:       s.Key,
		Value:     s.Value,
		CreatedAt: s.CreatedAt,
		SystemTag: s.SystemTag,
	}
}
