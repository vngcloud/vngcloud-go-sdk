package flavor

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
)

type GetOpts struct {
	FlavorID string
	common.CommonOpts
}

func (s *GetOpts) GetFlavorID() string {
	return s.FlavorID
}
