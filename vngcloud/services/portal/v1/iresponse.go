package v1

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type IGetPortalInfoResponse interface {
	ToEntityPortal() *lsentity.Portal
}
