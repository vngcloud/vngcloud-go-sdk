package objects

type (
	Server struct {
		BootVolumeId       string
		CreatedAt          string
		EncryptionVolume   bool
		Licence            bool
		Location           string
		Metadata           string
		MigrateState       string
		Name               string
		Product            string
		ServerGroupId      string
		ServerGroupName    string
		SshKeyName         string
		Status             string
		StopBeforeMigrate  bool
		User               string
		Uuid               string
		Image              Image
		Flavor             Flavor
		SecGroups          []ServerSecGroup
		ExternalInterfaces []NetworkInterface
		InternalInterfaces []NetworkInterface
	}

	NetworkInterface struct {
		CreatedAt     string
		FixedIp       string
		FloatingIp    string
		FloatingIpId  string
		InterfaceType string
		Mac           string
		NetworkUuid   string
		PortUuid      string
		Product       string
		ServerUuid    string
		Status        string
		SubnetUuid    string
		Type          string
		UpdatedAt     string
		Uuid          string
	}

	Flavor struct {
		Bandwidth              int64
		BandwidthUnit          string
		Cpu                    int64
		CpuPlatformDescription string
		FlavorId               string
		Gpu                    int64
		Group                  string
		Memory                 int64
		MetaData               string
		Name                   string
		RemainingVms           int64
		ZoneId                 string
	}

	Image struct {
		FlavorZoneIds []string
		Id            string
		ImageType     string
		ImageVersion  string
		Licence       bool
		PackageLimit  PackageLimit
	}

	PackageLimit struct {
		Cpu      int64
		DiskSize int64
		Memory   int64
	}

	ServerSecGroup struct {
		Name string
		Uuid string
	}
)
