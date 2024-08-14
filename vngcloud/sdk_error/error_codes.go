package sdk_error

// Common errors

const (
	EcUnknownError = ErrorCode("UnknownError")

	EcInternalServerError = ErrorCode("VngCloudApiInternalServerError")
	EcPagingInvalid       = ErrorCode("VngCloudApiPagingInvalid")
	EcPermissionDenied    = ErrorCode("VngCloudApiPermissionDenied")
	EcUnexpectedError     = ErrorCode("VngCloudApiUnexpectedError")
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
	EcVServerNetworkNotFound        = ErrorCode("VngCloudVServerNetworkNotFound")
	EcVServerSubnetNotFound         = ErrorCode("VngCloudVServerSubnetNotFound")
	EcVServerSubnetNotBelongNetwork = ErrorCode("VngCloudVServerSubnetNotBelongNetwork")
)

// vServer volume

const (
	EcVServerVolumeTypeNotFound              = ErrorCode("VngCloudVServerVolumeTypeNotFound")
	EcVServerVolumeNameNotValid              = ErrorCode("VngCloudVServerVolumeNameNotValid")
	EcVServerVolumeSizeOutOfRange            = ErrorCode("VngCloudVServerVolumeSizeOutOfRange")
	EcVServerVolumeSizeExceedGlobalQuota     = ErrorCode("VngCloudVServerVolumeSizeExceedGlobalQuota")
	EcVServerVolumeExceedQuota               = ErrorCode("VngCloudVServerVolumeExceedQuota")
	EcVServerVolumeNotFound                  = ErrorCode("VngCloudVServerVolumeNotFound")
	EcVServerVolumeAvailable                 = ErrorCode("VngCloudVServerVolumeAvailable")
	EcVServerVolumeAlreadyAttached           = ErrorCode("VngCloudVServerVolumeAlreadyAttached")
	EcVServerVolumeAlreadyAttachedThisServer = ErrorCode("VngCloudVServerVolumeAlreadyAttachedThisServer")
	EcVServerVolumeInProcess                 = ErrorCode("VngCloudVServerVolumeInProcess")
	EcVServerVolumeUnchanged                 = ErrorCode("VngCloudVServerVolumeUnchanged")
	EcVServerVolumeMustSameZone              = ErrorCode("VngCloudVServerVolumeMustSameZone")
	EcVServerVolumeMigrateMissingInit        = ErrorCode("VngCloudVServerVolumeMigrateMissingInit")
	EcVServerVolumeMigrateNeedProcess        = ErrorCode("VngCloudVServerVolumeMigrateNeedProcess")
	EcVServerVolumeMigrateNeedConfirm        = ErrorCode("VngCloudVServerVolumeMigrateNeedConfirm")
	EcVServerVolumeMigrateBeingProcess       = ErrorCode("VngCloudVServerVolumeMigrateBeingProcess")
	EcVServerVolumeMigrateBeingFinish        = ErrorCode("VngCloudVServerVolumeMigrateBeingFinish")
	EcVServerVolumeMigrateProcessingConfirm  = ErrorCode("VngCloudVServerVolumeMigrateProcessingConfirm")
	EcVServerVolumeMigrateBeingMigrating     = ErrorCode("VngCloudVServerVolumeMigrateBeingMigrating")
	EcVServerVolumeMigrateInSameZone         = ErrorCode("VngCloudVServerVolumeMigrateInSameZone")
	EcVServerVolumeIsMigrating               = ErrorCode("VngCloudVServerVolumeIsMigrating")
)

// Billing

const (
	EcBillingOutOfPoc = ErrorCode("BillingOutOfPoc")
)

// vServer server
const (
	EcVServerServerNotFound                       = ErrorCode("VngCloudVServerServerNotFound")
	EcVServerServerDeleteCreatingServer           = ErrorCode("VngCloudVServerServerDeleteCreatingServer")
	EcVServerServerExpired                        = ErrorCode("VngCloudVServerServerExpired")
	EcVServerServerUpdatingSecgroups              = ErrorCode("VngCloudVServerServerUpdatingSecgroups")
	EcVServerServerExceedQuota                    = ErrorCode("VngCloudVServerServerExceedQuota")
	EcVServerServerDeleteDeletingServer           = ErrorCode("VngCloudVServerServerDeleteDeletingServer")
	EcVServerServerDeleteBillingServer            = ErrorCode("VngCloudVServerServerDeleteBillingServer")
	EcVServerServerVolumeAttachQuotaExceeded      = ErrorCode("VngCloudVServerServerVolumeAttachQuotaExceeded")
	EcVServerServerAttachEncryptedVolume          = ErrorCode("VngCloudVServerServerAttachEncryptedVolume")
	EcVServerCreateBillingPaymentMethodNotAllowed = ErrorCode("VngCloudVServerCreateBillingPaymentMethodNotAllowed")
	EcVServerServerFlavorSystemExceedQuota        = ErrorCode("VngCloudVServerServerFlavorSystemExceedQuota")
)

// vServer quota

const (
	EcVServerQuotaNotFound = ErrorCode("VngCloudVServerQuotaNotFound")
)

// vServer snapshot
const (
	EcVServerSnapshotNameNotValid = ErrorCode("VngCloudVServerSnapshotNameNotValid")
	EcVServerSnapshotNotFound     = ErrorCode("VngCloudVServerSnapshotNotFound")
)

// LoadBalancer
const (
	EcVLBLoadBalancerNotFound            = ErrorCode("VngCloudVLBLoadBalancerNotFound")
	EcVLBLoadBalancerDuplicatePoolName   = ErrorCode("VngCloudVLBLoadBalancerDuplicatePoolName")
	EcVLBListenerDuplicateProtocolOrPort = ErrorCode("VngCloudVLBListenerDuplicateProtocolOrPort")
	EcVLBListenerDuplicateName           = ErrorCode("VngCloudVLBListenerDuplicateName")
	EcVLBPoolNotFound                    = ErrorCode("VngCloudVLBPoolNotFound")
	EcVLBPoolInUse                       = ErrorCode("VngCloudVLBPoolInUse")
	EcVLBLoadBalancerNotReady            = ErrorCode("VngCloudVLBLoadBalancerNotReady")
	EcVLBListenerNotFound                = ErrorCode("VngCloudVLBListenerNotFound")
	EcVLBMemberMustIdentical             = ErrorCode("VngCloudVLBMemberMustIdentical")
	EcVLBLoadBalancerExceedQuota         = ErrorCode("VngCloudVLBLoadBalancerExceedQuota")
)

// Endpoint

const (
	EcVNetworkEndpointStatusInvalid                     = ErrorCode("EndpointStatusInvalid")
	EcVNetworkEndpointOfVpcExists                       = ErrorCode("EndpointOfVpcIsExists")
	EcVNetworkEndpointNotFound                          = ErrorCode("VngCloudVNetworkEndpointNotFound")
	EcVNetworkEndpointPackageNotBelongToEndpointService = ErrorCode("EndpointPackageNotBelongToEndpointService")
	EcVNetworkContainInvalidCharacter                   = ErrorCode("ContainInvalidCharacter")
	EcVNetworkLockOnProcess                             = ErrorCode("LockOnProcess")
)
