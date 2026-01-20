package sdk_error

// quota error category
const (
	ErrCatQuota ErrorCategory = "quota"
)

// login error category

const (
	ErrCatIam ErrorCategory = "iam"
)

// infra error category
const (
	ErrCatInfra ErrorCategory = "infra"
)

// purchase error category

const (
	ErrCatPurchase ErrorCategory = "purchase"
)

const (
	ErrCatAll ErrorCategory = "all"
)

const (
	ErrCatProductVlb      ErrorCategory = "vlb"
	ErrCatProductVNetwork ErrorCategory = "vnetwork"
	ErrCatProductVdns     ErrorCategory = "vdns"
)

const (
	ErrCatVServer ErrorCategory = "vserver"

	ErrCatVirtualAddress ErrorCategory = "virtualaddress"
)
