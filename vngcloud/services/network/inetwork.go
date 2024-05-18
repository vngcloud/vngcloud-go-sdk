package network

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lsnetworkSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network/v2"
)

type INetworkServiceV2 interface {
	// The group of Network APIs

	GetNetworkById(popts lsnetworkSvcV2.IGetNetworkByIdRequest) (*lsentity.Network, lserr.ISdkError)

	// The group of Secgroup APIs

	GetSecgroupById(popts lsnetworkSvcV2.IGetSecgroupByIdRequest) (*lsentity.Secgroup, lserr.ISdkError)
	CreateSecgroup(popts lsnetworkSvcV2.ICreateSecgroupRequest) (*lsentity.Secgroup, lserr.ISdkError)
	DeleteSecgroupById(popts lsnetworkSvcV2.IDeleteSecgroupByIdRequest) lserr.ISdkError

	// The group of SecgroupRule APIs

	CreateSecgroupRule(popts lsnetworkSvcV2.ICreateSecgroupRuleRequest) (*lsentity.SecgroupRule, lserr.ISdkError)
	DeleteSecgroupRuleById(popts lsnetworkSvcV2.IDeleteSecgroupRuleByIdRequest) lserr.ISdkError
	ListSecgroupRulesBySecgroupId(popts lsnetworkSvcV2.IListSecgroupRulesBySecgroupIdRequest) (*lsentity.ListSecgroupRules, lserr.ISdkError)
}
