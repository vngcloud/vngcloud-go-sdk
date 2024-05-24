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
	EcUnknownAuthFailure   = ErrorCode("VngCloudIamUnknownAuthFailure")
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
	EcVServerSecgroupRuleNotFound      = ErrorCode("VngCloudVServerSecgroupRuleNotFound")
	EcVServerSecgroupRuleAlreadyExists = ErrorCode("VngCloudVServerSecgroupRuleAlreadyExists")
	EcVServerSecgroupRuleExceedQuota   = ErrorCode("VngCloudVServerSecgroupRuleExceedQuota")
)

// VServer network

const (
	EcVServerNetworkNotFound = ErrorCode("VngCloudVServerNetworkNotFound")
	EcVServerSubnetNotFound  = ErrorCode("VngCloudVServerSubnetNotFound")
)

// vServer volume

const (
	EcVServerVolumeTypeNotFound = ErrorCode("VngCloudVServerVolumeTypeNotFound")
)

// Billing

const (
	EcBillingOutOfPoc = ErrorCode("BillingOutOfPoc")
)

// vServer server
const (
	EcVServerServerNotFound                       = ErrorCode("VngCloudVServerServerNotFound")
	EcVServerServerDeleteCreatingServer           = ErrorCode("VngCloudVServerServerDeleteCreatingServer")
	EcVServerServerExceedQuota                    = ErrorCode("VngCloudVServerServerExceedQuota")
	EcVServerServerDeleteDeletingServer           = ErrorCode("VngCloudVServerServerDeleteDeletingServer")
	EcVServerServerDeleteBillingServer            = ErrorCode("VngCloudVServerServerDeleteBillingServer")
	EcVServerCreateBillingPaymentMethodNotAllowed = ErrorCode("VngCloudVServerCreateBillingPaymentMethodNotAllowed")
)

// vServer quota

const (
	EcVServerQuotaNotFound = ErrorCode("VngCloudVServerQuotaNotFound")
)
