package v2

// Secgroup

type IGetSecgroupByIdRequest interface {
	GetSecgroupId() string
}

type ICreateSecgroupRequest interface {
	ToRequestBody() interface{}
	GetSecgroupName() string
}

type IDeleteSecgroupByIdRequest interface {
	GetSecgroupId() string
}

type IListSecgroupRequest interface {
}

// Secgroup Rule
type ICreateSecgroupRuleRequest interface {
	GetSecgroupId() string
	ToRequestBody() interface{}
}

type IDeleteSecgroupRuleByIdRequest interface {
	GetSecgroupId() string
	GetSecgroupRuleId() string
}

type IListSecgroupRulesBySecgroupIdRequest interface {
	GetSecgroupId() string
}

// Network

type IGetNetworkByIdRequest interface {
	GetNetworkId() string
}

// Subnet

type IGetSubnetByIdRequest interface {
	ParseUserAgent() string
	GetNetworkId() string
	GetSubnetId() string
}

type IUpdateSubnetByIdRequest interface {
	ParseUserAgent() string
	GetNetworkId() string
	GetSubnetId() string
	ToRequestBody() interface{}
}

// Address Pair

type IGetAllAddressPairByVirtualSubnetIdRequest interface {
	GetVirtualSubnetId() string
	ParseUserAgent() string
}

type ISetAddressPairInVirtualSubnetRequest interface {
	GetVirtualSubnetId() string
	ParseUserAgent() string
	ToRequestBody() interface{}
}

type IDeleteAddressPairRequest interface {
	ParseUserAgent() string
	GetAddressPairID() string
}

type IListAllServersBySecgroupIdRequest interface {
	GetSecgroupId() string
}

/**
 * The interface request group of Virtual Address API
 */

// Request interface for creating virtual address cross project
type ICreateVirtualAddressCrossProjectRequest interface {
	ToRequestBody() interface{}
	ParseUserAgent() string
	ToMap() map[string]interface{}
	AddUserAgent(pagent ...string) ICreateVirtualAddressCrossProjectRequest
	WithDescription(pdescription string) ICreateVirtualAddressCrossProjectRequest
}

// Request interface for deleting virtual address by ID
type IDeleteVirtualAddressByIdRequest interface {
	GetVirtualAddressId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IDeleteVirtualAddressByIdRequest
	ToMap() map[string]interface{}
}


type IGetVirtualAddressByIdRequest interface {
	GetVirtualAddressId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IGetVirtualAddressByIdRequest
	ToMap() map[string]interface{}
}
