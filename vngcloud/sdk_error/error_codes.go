package sdk_error

// Common errors

const (
	EcUnknownError = ErrorCode("UnknownError")

	EcInternalServerError   = ErrorCode("VngCloudApiInternalServerError")
	EcServiceMaintenance    = ErrorCode("VngCloudServiceMaintenance")
	EcPagingInvalid         = ErrorCode("VngCloudApiPagingInvalid")
	EcTagKeyInvalid         = ErrorCode("VngCloudApiTagKeyInvalid")
	EcPermissionDenied      = ErrorCode("VngCloudApiPermissionDenied")
	EcUnexpectedError       = ErrorCode("VngCloudApiUnexpectedError")
	EcPaymentMethodNotAllow = ErrorCode("VngCloudPaymentMethodNotAllow")
	EcCreditNotEnough       = ErrorCode("VngCloudCredtNotEnough")
	EcProjectConflict       = ErrorCode("VngCloudProjectConflict")
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

	EcVServerInternalNetworkInterfaceNotFound = ErrorCode("VngCloudVServerInternalNetworkInterfaceNotFound")
	EcVServerWanIpAvailable                   = ErrorCode("VngCloudVServerWanIpAvailable")
	EcVServerWanIdNotFound                    = ErrorCode("VngCloudVServerWanIdNotFound")

	EcVServerAddressPairExisted = ErrorCode("VngCloudVServerAddressPairExisted")
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
	EcVServerFlavorNotSupported              = ErrorCode("VngCloudVServerFlavorNotSupported")
)

// Billing

const (
	EcPurchaseIssue = ErrorCode("PurchaseIssue")
)

// vServer server
const (
	EcVServerImageNotFound                        = ErrorCode("VngCloudVServerImageNotFound")
	EcVServerServerNotFound                       = ErrorCode("VngCloudVServerServerNotFound")
	EcVServerServerDeleteCreatingServer           = ErrorCode("VngCloudVServerServerDeleteCreatingServer")
	EcVServerServerExpired                        = ErrorCode("VngCloudVServerServerExpired")
	EcVServerServerUpdatingSecgroups              = ErrorCode("VngCloudVServerServerUpdatingSecgroups")
	EcVServerServerExceedQuota                    = ErrorCode("VngCloudVServerServerExceedQuota")
	EcVServerServerDeleteDeletingServer           = ErrorCode("VngCloudVServerServerDeleteDeletingServer")
	EcVServerServerDeleteBillingServer            = ErrorCode("VngCloudVServerServerDeleteBillingServer")
	EcVServerServerVolumeAttachQuotaExceeded      = ErrorCode("VngCloudVServerServerVolumeAttachQuotaExceeded")
	EcVServerServerAttachEncryptedVolume          = ErrorCode("VngCloudVServerServerAttachEncryptedVolume")
	EcVServerServerFlavorSystemExceedQuota        = ErrorCode("VngCloudVServerServerFlavorSystemExceedQuota")
	EcVServerServerExceedCpuQuota                 = ErrorCode("VngCloudVServerServerExceedCpuQuota")
	EcVServerCreateBillingPaymentMethodNotAllowed = ErrorCode("VngCloudVServerCreateBillingPaymentMethodNotAllowed")
	EcVServerServerImageNotSupported              = ErrorCode("VngCloudVServerServerImageNotSupported")
	EcVServerServerCanNotAttachFloatingIp         = ErrorCode("VngCloudVServerServerCanNotAttachFloatingIp")
	EcVServerServerExceedFloatingIpQuota          = ErrorCode("VngCloudVServerServerExceedFloatingIpQuota")
	EcVServerServerGroupNotFound                  = ErrorCode("VngCloudVServerServerGroupNotFound")
	EcVServerServerGroupInUse                     = ErrorCode("VngCloudVServerServerGroupInUse")
	EcVServerServerGroupNameMustBeUnique          = ErrorCode("VngCloudVServerServerGroupNameMustBeUnique")
)

// vServer virtual address
const (
	EcVServerVirtualAddressNotFound    = ErrorCode("VngCloudVServerVirtualAddressNotFound")
	EcVServerVirtualAddressExceedQuota = ErrorCode("VngCloudVServerVirtualAddressExceedQuota")
	EcVServerVirtualAddressInUse       = ErrorCode("VngCloudVServerVirtualAddressInUse")
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
	EcVLBLoadBalancerIsDeleting          = ErrorCode("VngCloudVLBLoadBalancerIsDeleting")
	EcVLBLoadBalancerIsCreating          = ErrorCode("VngCloudVLBLoadBalancerIsCreating")
	EcVLBLoadBalancerIsUpdating          = ErrorCode("VngCloudVLBLoadBalancerIsUpdating")
	EcVLBLoadBalancerResizeSamePackage   = ErrorCode("VngCloudVLBLoadBalancerResizeSamePackage")
	EcVLBLoadBalancerPackageNotFound     = ErrorCode("VngCloudVLBLoadBalancerPackageNotFound")
)

// Endpoint

const (
	EcVNetworkEndpointStatusInvalid                     = ErrorCode("EndpointStatusInvalid")
	EcVNetworkEndpointOfVpcExists                       = ErrorCode("EndpointOfVpcIsExists")
	EcVNetworkEndpointNotFound                          = ErrorCode("VngCloudVNetworkEndpointNotFound")
	EcVNetworkEndpointPackageNotBelongToEndpointService = ErrorCode("EndpointPackageNotBelongToEndpointService")
	EcVNetworkContainInvalidCharacter                   = ErrorCode("ContainInvalidCharacter")
	EcVNetworkLockOnProcess                             = ErrorCode("LockOnProcess")
	EcVNetworkEndpointTagNotFound                       = ErrorCode("EndpointTagNotFound")
	EcVNetworkEndpointTagExisted                        = ErrorCode("EndpointTagExisted")
)

// Global Load Balancer

const (
	EcGlobalLoadBalancerNotFound = ErrorCode("VngCloudGlobalLoadBalancerNotFound")
)
