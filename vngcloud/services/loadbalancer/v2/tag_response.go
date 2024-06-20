package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type ListTagResponse struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	CreatedAt string `json:"createdAt"`
}

type ListTagsResponse []ListTagResponse

func (s ListTagResponse) ToEntityTag() *lsentity.Tag {
	return &lsentity.Tag{
		Key:   s.Key,
		Value: s.Value,
	}
}

func (s *ListTagsResponse) ToEntityListTags() *lsentity.ListTags {
	result := new(lsentity.ListTags)
	if s == nil {
		return result
	}

	for _, item := range *s {
		result.Add(item.ToEntityTag())
	}

	return result
}
