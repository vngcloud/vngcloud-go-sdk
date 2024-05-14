package flavor

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

type (
	GetResponse struct {
		Flavors []GetResponseFlavor `json:"flavors"`
	}

	GetResponseFlavor struct {
		Name string `json:"name"`
		Cpu  int    `json:"cpu"`
		Mem  int    `json:"memory"`
		Gpu  int    `json:"gpu"`
		Id   string `json:"flavorId"`
	}
)

func (s *GetResponse) ToFlavorObject() *objects.Flavor {
	if len(s.Flavors) != 1 {
		return nil
	}

	return &objects.Flavor{
		Name:     s.Flavors[0].Name,
		Cpu:      int64(s.Flavors[0].Cpu),
		Memory:   int64(s.Flavors[0].Mem),
		Gpu:      int64(s.Flavors[0].Gpu),
		FlavorId: s.Flavors[0].Id,
	}
}
