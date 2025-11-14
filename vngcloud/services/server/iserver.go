package server

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lsnserverSvcV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/server/v1"
)

type IServerServiceInternalV1 interface {
	CreateSystemTags(popts lsnserverSvcV1.ICreateSystemTagRequest) (*[]lsentity.SystemTag, lserr.IError)
}
