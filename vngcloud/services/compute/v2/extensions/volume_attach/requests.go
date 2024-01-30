package volume_attach

import "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"

// ******************************************** Create Attach Volume Request *******************************************

type CreateOpts struct {
	AttachCommonOpts
	common.CommonOpts
}

// ***************************************** Delete Volume Attachment Requests *****************************************

type DeleteOpts struct {
	AttachCommonOpts
	common.CommonOpts
}
