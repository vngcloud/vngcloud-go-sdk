package tag

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
	v2 "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/compute/v2"
)

type TagRequestItem struct {
	IsEdited bool   `json:"isEdited"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}

// ******************************************** Create Attach Volume Request *******************************************

type CreateOpts struct {
	ResourceType   string           `json:"resourceType"`
	ResourceID     string           `json:"resourceId"`
	TagRequestList []TagRequestItem `json:"tagRequestList"`

	v2.ResourceV2Common
	common.CommonOpts
}

func (s *CreateOpts) ToRequestBody() interface{} {
	return s
}

// ***************************************** Delete Volume Attachment Requests *****************************************

type GetOpts struct {
	v2.ResourceV2Common
	common.CommonOpts
}

// ***************************************** Update Tags *****************************************

type UpdateOpts struct {
	ResourceType   string           `json:"resourceType"`
	ResourceID     string           `json:"resourceId"`
	TagRequestList []TagRequestItem `json:"tagRequestList"`

	v2.ResourceV2Common
	common.CommonOpts
}

func (s *UpdateOpts) ToRequestBody() interface{} {
	return s
}
