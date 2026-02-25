package entity

import lstr "strings"

const (
	ServerStatusActive  = "ACTIVE"
	ServerStatusError   = "ERROR"
	ServerStatusStopped = "STOPPED"
)

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
		SecGroups          []ServerSecgroup
		ExternalInterfaces []NetworkInterface
		InternalInterfaces []NetworkInterface
		ZoneId             string
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

	ServerSecgroup struct {
		Name string
		Uuid string
	}

	SystemTag struct {
		Key       string
		Value     string
		CreatedAt string
		SystemTag bool
	}
)

type ListServers struct {
	Items []*Server
}

func (s *Server) CanDelete() bool {
	switch lstr.ToUpper(s.Status) {
	case ServerStatusActive, ServerStatusError, ServerStatusStopped:
		return true
	}

	return false
}

func (s *Server) IsRunning() bool {
	switch lstr.ToUpper(s.Status) {
	case ServerStatusActive:
		return true
	}

	return false
}

func (s *Server) GetInternalInterfaceWanInfo() (string, string, string, bool) {
	for _, i := range s.InternalInterfaces {
		if i.FloatingIp != "" {
			return i.Uuid, i.FloatingIpId, i.FloatingIp, true
		}
	}

	return "", "", "", false
}

func (s *Server) GetInternalNetworkInterfaceIds() []string {
	ids := make([]string, 0, len(s.InternalInterfaces))
	for _, i := range s.InternalInterfaces {
		ids = append(ids, i.Uuid)
	}

	return ids
}

func (s *Server) InternalNetworkInterfacePossible() bool {
	return len(s.InternalInterfaces) > 0
}

func (s *Server) CanAttachFloatingIp() bool {
	if !s.InternalNetworkInterfacePossible() {
		return false
	}

	for _, i := range s.InternalInterfaces {
		if i.FloatingIp != "" {
			return false
		}
	}

	return true
}
