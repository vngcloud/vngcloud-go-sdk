package entity

const (
	OpenSearchClusterStatusActive   = "ACTIVE"
	OpenSearchClusterStatusCreating = "CREATING"
	OpenSearchClusterStatusError    = "ERROR"
)

type (
	OpenSearchCluster struct {
		// Id is the full UUID used in API URLs (e.g. op-d0e2c62b-8130-46be-b65e-3375549af42a).
		// ClusterId is a truncated public-facing identifier — keep both for round-trip parity.
		Id          string
		ClusterId   string
		ProjectId   string
		Name        string
		Description string
		Status      string
		Version     string
		Region      string

		NumberOfNodes int32
		StorageSize   int32

		EnableTls     bool
		PublicAccess  bool
		PrivateAccess bool
		EncryptVolume bool
		EnableVpcDns  bool

		VpcId    string
		VpcCidr  string
		SubnetId string

		PackageId     string
		StorageTypeId string
		ConfigGroupId string
		PortalUserId  int32

		CreatedAt string
		UpdatedAt string
		DeletedAt string

		PackageDetail OpenSearchPackageDetail
		StorageType   OpenSearchStorageType
		ConfigGroup   OpenSearchConfigGroup
		BillingStatus OpenSearchBillingStatus
	}

	OpenSearchPackageDetail struct {
		Id                 string
		Name               string
		Type               string
		Description        string
		NetworkPerformance string
		IsDefault          bool
		Ram                int32
		Cpu                int32
		Platform           string
	}

	OpenSearchStorageType struct {
		Id   string
		Name string
		Min  int32
		Max  int32
	}

	OpenSearchConfigGroup struct {
		Name    string
		Version string
	}

	OpenSearchBillingStatus struct {
		Status      string
		RenewPeriod int32
	}

	OpenSearchEndpoint struct {
		Url     string
		Type    string // "dashboard" (:443 UI) or "log" (:9200 REST API)
		Cidrs   []string
		Private bool
	}
)

const (
	OpenSearchEndpointTypeDashboard = "dashboard"
	OpenSearchEndpointTypeLog       = "log"
)

func (s *OpenSearchEndpoint) IsLogIngest() bool {
	return s.Type == OpenSearchEndpointTypeLog
}

func (s *OpenSearchEndpoint) IsDashboard() bool {
	return s.Type == OpenSearchEndpointTypeDashboard
}

func (s *OpenSearchCluster) IsActive() bool {
	return s.Status == OpenSearchClusterStatusActive
}

func (s *OpenSearchCluster) IsPubliclyAccessible() bool {
	return s.PublicAccess
}
