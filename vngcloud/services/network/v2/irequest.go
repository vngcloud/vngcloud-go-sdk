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