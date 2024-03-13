package tag

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
	v2 "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/compute/v2"
)

// ******************************************** Create Attach Volume Request *******************************************

type CreateOpts struct {
	ResourceID     string                     `json:"resourceId"`
	ResourceType   string                     `json:"resourceType"`
	TagRequestList []CreateOptsTagRequestItem `json:"tagRequestList"`

	v2.ServerV2Common
	common.CommonOpts
}

type CreateOptsTagRequestItem struct {
	IsEdited bool   `json:"isEdited"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}

func (s *CreateOpts) ToRequestBody() interface{} {
	return s
}

// ***************************************** Delete Volume Attachment Requests *****************************************

type GetOpts struct {
	v2.ServerV2Common
	common.CommonOpts
}
