package v2

// Secgroup

type IGetSecgroupByIdRequest interface {
	AddUserAgent(pagent ...string) IGetSecgroupByIdRequest
	ParseUserAgent() string
	GetSecgroupId() string
}

type ICreateSecgroupRequest interface {
	ToRequestBody() interface{}
	GetSecgroupName() string
	AddUserAgent(pagent ...string) ICreateSecgroupRequest
	ParseUserAgent() string
}

type IDeleteSecgroupByIdRequest interface {
	GetSecgroupId() string
	AddUserAgent(pagent ...string) IDeleteSecgroupByIdRequest
	ParseUserAgent() string
}

type IListSecgroupRequest interface {
	AddUserAgent(pagent ...string) IListSecgroupRequest
	ParseUserAgent() string
}

// Secgroup Rule
type ICreateSecgroupRuleRequest interface {
	GetSecgroupId() string
	ToRequestBody() interface{}
	AddUserAgent(pagent ...string) ICreateSecgroupRuleRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IDeleteSecgroupRuleByIdRequest interface {
	GetSecgroupId() string
	GetSecgroupRuleId() string
	AddUserAgent(pagent ...string) IDeleteSecgroupRuleByIdRequest
	ParseUserAgent() string
}

type IListSecgroupRulesBySecgroupIdRequest interface {
	GetSecgroupId() string
	AddUserAgent(pagent ...string) IListSecgroupRulesBySecgroupIdRequest
	ParseUserAgent() string
}

// Network

type IGetNetworkByIdRequest interface {
	GetNetworkId() string
	AddUserAgent(pagent ...string) IGetNetworkByIdRequest
	ParseUserAgent() string
}

// Subnet

type IGetSubnetByIdRequest interface {
	AddUserAgent(pagent ...string) IGetSubnetByIdRequest
	ParseUserAgent() string
	GetNetworkId() string
	GetSubnetId() string
}

type IUpdateSubnetByIdRequest interface {
	AddUserAgent(pagent ...string) IUpdateSubnetByIdRequest
	ParseUserAgent() string
	GetNetworkId() string
	GetSubnetId() string
	ToRequestBody() interface{}
}

/**
 * The interface request group of Address Pair API
 */

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
	AddUserAgent(pagent ...string) IDeleteAddressPairRequest
}

type IListAllServersBySecgroupIdRequest interface {
	GetSecgroupId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IListAllServersBySecgroupIdRequest
}

type ICreateAddressPairRequest interface {
	GetVirtualAddressId() string
	ToRequestBody() interface{}
	ParseUserAgent() string
	ToMap() map[string]interface{}
	AddUserAgent(pagent ...string) ICreateAddressPairRequest
	WithMode(pmode AddressPairMode) ICreateAddressPairRequest
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

// Request interface for getting virtual address by ID
type IGetVirtualAddressByIdRequest interface {
	GetVirtualAddressId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IGetVirtualAddressByIdRequest
	ToMap() map[string]interface{}
}

// Request interface for listing address pairs of virtual address by ID

type IListAddressPairsByVirtualAddressIdRequest interface {
	GetVirtualAddressId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IListAddressPairsByVirtualAddressIdRequest
	ToMap() map[string]interface{}
}
