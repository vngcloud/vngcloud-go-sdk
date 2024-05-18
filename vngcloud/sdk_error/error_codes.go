package sdk_error

// Common errors

const (
	EcUnknownError = ErrorCode("UnknownError")

	EcInternalServerError = ErrorCode("VngCloudApiInternalServerError")
)

// Internal SDK error

const (
	EcReauthFuncNotSet = ErrorCode("SdkReauthFunctionNotSet")
)

// VngCloud IAM gateway errros

const (
	EcAuthenticationFailed = ErrorCode("VngCloudIamAuthenticationFailed")
	EcTooManyFailedLogins  = ErrorCode("VngCloudIamTooManyFailedLogins")
)

// VServer secgroup

const (
	EcVServerSecgroupNotFound          = ErrorCode("VngCloudVServerSecgroupNotFound")
	EcVServerSecgroupNameAlreadyExists = ErrorCode("VngCloudVServerSecgroupNameAlreadyExists")
	EcVServerSecgroupExceedQuota       = ErrorCode("VngCloudVServerSecgroupExceedQuota")
	EcVServerSecgroupInUse             = ErrorCode("VngCloudVServerSecgroupInUse")
)

// VServer secgroup rule

const (
	EcVServerSecgroupRuleNotFound = ErrorCode("VngCloudVServerSecgroupRuleNotFound")
)

// VServer network

const (
	EcVServerNetworkNotFound = ErrorCode("VngCloudVServerNetworkNotFound")
)
