package compute

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lscomputeSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/compute/v2"
)

type IComputeServiceV2 interface {
	CreateServer(popts lscomputeSvcV2.ICreateServerRequest) (*lsentity.Server, lserr.ISdkError)
}
