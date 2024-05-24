package portal

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lsdkErr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lsportalV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/portal/v1"
)

type IPortalServiceV1 interface {
	GetPortalInfo(popts lsportalV1.IGetPortalInfoRequest) (*lsentity.Portal, lsdkErr.ISdkError)
}

type IPortalServiceV2 interface {
}
