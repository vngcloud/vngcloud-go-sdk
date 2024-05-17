package network

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lsnetworkSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network/v2"
)

type INetworkServiceV2 interface {
	GetSecgroupById(popts lsnetworkSvcV2.IGetSecgroupByIdRequest) (*lsentity.Secgroup, lserr.ISdkError)
	CreateSecgroup(popts lsnetworkSvcV2.ICreateSecgroupRequest) (*lsentity.Secgroup, lserr.ISdkError)
}
