package portal

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lsportalV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/portal/v1"
	lsportalV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/portal/v2"
)

type IPortalServiceV1 interface {
	ListZones() (*lsentity.ListZones, lserr.IError)
	GetPortalInfo(popts lsportalV1.IGetPortalInfoRequest) (*lsentity.Portal, lserr.IError)
	ListProjects(popts lsportalV1.IListProjectsRequest) (*lsentity.ListPortals, lserr.IError)
}

type IPortalServiceV2 interface {
	ListAllQuotaUsed() (*lsentity.ListQuotas, lserr.IError)
	GetQuotaByName(popts lsportalV2.IGetQuotaByNameRequest) (*lsentity.Quota, lserr.IError)
}
