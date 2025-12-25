package network

import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lsnetworkSvcV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network/v1"
	lsnetworkSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network/v2"
)

type INetworkServiceV1 interface {
	GetEndpointById(pops lsnetworkSvcV1.IGetEndpointByIdRequest) (*lsentity.Endpoint, lserr.IError)
	CreateEndpoint(popts lsnetworkSvcV1.ICreateEndpointRequest) (*lsentity.Endpoint, lserr.IError)
	DeleteEndpointById(popts lsnetworkSvcV1.IDeleteEndpointByIdRequest) lserr.IError
	ListEndpoints(popts lsnetworkSvcV1.IListEndpointsRequest) (*lsentity.ListEndpoints, lserr.IError)
}

type INetworkServiceInternalV1 interface {
	ListTagsByEndpointId(popts lsnetworkSvcV1.IListTagsByEndpointIdRequest) (*lsentity.ListTags, lserr.IError)
	CreateTagsWithEndpointId(popts lsnetworkSvcV1.ICreateTagsWithEndpointIdRequest) lserr.IError
	DeleteTagOfEndpoint(popts lsnetworkSvcV1.IDeleteTagOfEndpointRequest) lserr.IError
	UpdateTagValueOfEndpoint(popts lsnetworkSvcV1.IUpdateTagValueOfEndpointRequest) lserr.IError
	CreateEndpoint(popts lsnetworkSvcV1.ICreateEndpointRequest) (*lsentity.Endpoint, lserr.IError)
}

type INetworkServiceV2 interface {
	// The group of Network APIs
	GetNetworkById(popts lsnetworkSvcV2.IGetNetworkByIdRequest) (*lsentity.Network, lserr.IError)

	// The group of Secgroup APIs

	GetSecgroupById(popts lsnetworkSvcV2.IGetSecgroupByIdRequest) (*lsentity.Secgroup, lserr.IError)
	CreateSecgroup(popts lsnetworkSvcV2.ICreateSecgroupRequest) (*lsentity.Secgroup, lserr.IError)
	ListSecgroup(popts lsnetworkSvcV2.IListSecgroupRequest) (*lsentity.ListSecgroups, lserr.IError)
	DeleteSecgroupById(popts lsnetworkSvcV2.IDeleteSecgroupByIdRequest) lserr.IError

	// The group of SecgroupRule APIs

	CreateSecgroupRule(popts lsnetworkSvcV2.ICreateSecgroupRuleRequest) (*lsentity.SecgroupRule, lserr.IError)
	DeleteSecgroupRuleById(popts lsnetworkSvcV2.IDeleteSecgroupRuleByIdRequest) lserr.IError
	ListSecgroupRulesBySecgroupId(popts lsnetworkSvcV2.IListSecgroupRulesBySecgroupIdRequest) (*lsentity.ListSecgroupRules, lserr.IError)

	// Subnet
	GetSubnetById(popts lsnetworkSvcV2.IGetSubnetByIdRequest) (*lsentity.Subnet, lserr.IError)
	UpdateSubnetById(popts lsnetworkSvcV2.IUpdateSubnetByIdRequest) (*lsentity.Subnet, lserr.IError)

	// Address Pair
	GetAllAddressPairByVirtualSubnetId(popts lsnetworkSvcV2.IGetAllAddressPairByVirtualSubnetIdRequest) ([]*lsentity.AddressPair, lserr.IError)
	SetAddressPairInVirtualSubnet(popts lsnetworkSvcV2.ISetAddressPairInVirtualSubnetRequest) (*lsentity.AddressPair, lserr.IError)
	DeleteAddressPair(popts lsnetworkSvcV2.IDeleteAddressPairRequest) lserr.IError
	CreateAddressPair(popts lsnetworkSvcV2.ICreateAddressPairRequest) (*lsentity.AddressPair, lserr.IError)

	// Servers
	ListAllServersBySecgroupId(popts lsnetworkSvcV2.IListAllServersBySecgroupIdRequest) (*lsentity.ListServers, lserr.IError)

	// Virtual Address API group
	CreateVirtualAddressCrossProject(popts lsnetworkSvcV2.ICreateVirtualAddressCrossProjectRequest) (*lsentity.VirtualAddress, lserr.IError)
	DeleteVirtualAddressById(popts lsnetworkSvcV2.IDeleteVirtualAddressByIdRequest) lserr.IError
	GetVirtualAddressById(popts lsnetworkSvcV2.IGetVirtualAddressByIdRequest) (*lsentity.VirtualAddress, lserr.IError)
	ListAddressPairsByVirtualAddressId(popts lsnetworkSvcV2.IListAddressPairsByVirtualAddressIdRequest) (*lsentity.ListAddressPairs, lserr.IError)
}
