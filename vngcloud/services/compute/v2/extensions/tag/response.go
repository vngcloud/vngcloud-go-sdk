package tag

import "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"

type ItemTagResponse struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	CreatedAt string `json:"createdAt"`
	SystemTag bool   `json:"systemTag"`
}

func (s *ItemTagResponse) ToResourceTagObject() *objects.ResourceTag {
	if s == nil {
		return nil
	}

	return &objects.ResourceTag{
		Key:       s.Key,
		Value:     s.Value,
		CreatedAt: s.CreatedAt,
		SystemTag: s.SystemTag,
	}
}
