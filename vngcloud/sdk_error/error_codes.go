package sdk_error

// Common errors

const (
	EcUnknownError = ErrorCode("UnknownError")

	EcInternalServerError = ErrorCode("VngCloudApiInternalServerError")
	EcPagingInvalid       = ErrorCode("VngCloudApiPagingInvalid")
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
	EcVServerVolumeTypeNotFound              = ErrorCode("VngCloudVServerVolumeTypeNotFound")
	EcVServerVolumeNameNotValid              = ErrorCode("VngCloudVServerVolumeNameNotValid")
	EcVServerVolumeSizeOutOfRange            = ErrorCode("VngCloudVServerVolumeSizeOutOfRange")
	EcVServerVolumeNotFound                  = ErrorCode("VngCloudVServerVolumeNotFound")
	EcVServerVolumeAvailable                 = ErrorCode("VngCloudVServerVolumeAvailable")
	EcVServerVolumeAlreadyAttached           = ErrorCode("VngCloudVServerVolumeAlreadyAttached")
	EcVServerVolumeAlreadyAttachedThisServer = ErrorCode("VngCloudVServerVolumeAlreadyAttachedThisServer")
	EcVServerVolumeInProcess                 = ErrorCode("VngCloudVServerVolumeInProcess")
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
	EcVServerServerVolumeAttachQuotaExceeded      = ErrorCode("VngCloudVServerServerVolumeAttachQuotaExceeded")
	EcVServerCreateBillingPaymentMethodNotAllowed = ErrorCode("VngCloudVServerCreateBillingPaymentMethodNotAllowed")
)

// vServer quota

const (
	EcVServerQuotaNotFound = ErrorCode("VngCloudVServerQuotaNotFound")
)
