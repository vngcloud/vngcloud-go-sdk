package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type (
	Server struct {
		BootVolumeId       string             `json:"bootVolumeId"`
		CreatedAt          string             `json:"createdAt"`
		EncryptionVolume   bool               `json:"encryptionVolume"`
		Licence            bool               `json:"licence"`
		Location           string             `json:"location"`
		Metadata           string             `json:"metadata"`
		MigrateState       string             `json:"migrateState"`
		Name               string             `json:"name"`
		Product            string             `json:"product"`
		ServerGroupId      string             `json:"serverGroupId"`
		ServerGroupName    string             `json:"serverGroupName"`
		SshKeyName         string             `json:"sshKeyName"`
		Status             string             `json:"status"`
		StopBeforeMigrate  bool               `json:"stopBeforeMigrate"`
		User               string             `json:"user"`
		Uuid               string             `json:"uuid"`
		Image              Image              `json:"image"`
		Flavor             Flavor             `json:"flavor"`
		SecGroups          []ServerSecgroup   `json:"secGroups"`
		ExternalInterfaces []NetworkInterface `json:"externalInterfaces"`
		InternalInterfaces []NetworkInterface `json:"internalInterfaces"`
	}

	NetworkInterface struct {
		CreatedAt     string `json:"createdAt"`
		FixedIp       string `json:"fixedIp"`
		FloatingIp    string `json:"floatingIp"`
		FloatingIpId  string `json:"floatingIpId"`
		InterfaceType string `json:"interfaceType"`
		Mac           string `json:"mac"`
		NetworkUuid   string `json:"networkUuid"`
		PortUuid      string `json:"portUuid"`
		Product       string `json:"product"`
		ServerUuid    string `json:"serverUuid"`
		Status        string `json:"status"`
		SubnetUuid    string `json:"subnetUuid"`
		Type          string `json:"type"`
		UpdatedAt     string `json:"updatedAt"`
		Uuid          string `json:"uuid"`
	}

	Flavor struct {
		Bandwidth              int64  `json:"bandwidth"`
		BandwidthUnit          string `json:"bandwidthUnit"`
		Cpu                    int64  `json:"cpu"`
		CpuPlatformDescription string `json:"cpuPlatformDescription"`
		FlavorId               string `json:"flavorId"`
		Gpu                    int64  `json:"gpu"`
		Group                  string `json:"group"`
		Memory                 int64  `json:"memory"`
		MetaData               string `json:"metaData"`
		Name                   string `json:"name"`
		RemainingVms           int64  `json:"remainingVms"`
		ZoneId                 string `json:"zoneId"`
	}

	Image struct {
		FlavorZoneIds []string     `json:"flavorZoneIds"`
		Id            string       `json:"id"`
		ImageType     string       `json:"imageType"`
		ImageVersion  string       `json:"imageVersion"`
		Licence       bool         `json:"licence"`
		PackageLimit  PackageLimit `json:"packageLimit"`
	}

	PackageLimit struct {
		Cpu      int64 `json:"cpu"`
		DiskSize int64 `json:"diskSize"`
		Memory   int64 `json:"memory"`
	}

	ServerSecgroup struct {
		Name string `json:"name"`
		Uuid string `json:"uuid"`
	}
)

func (s Image) toEntityImage() lsentity.Image {
	return lsentity.Image{
		FlavorZoneIds: s.FlavorZoneIds,
		Id:            s.Id,
		ImageType:     s.ImageType,
		ImageVersion:  s.ImageVersion,
		Licence:       s.Licence,
		PackageLimit: lsentity.PackageLimit{
			Cpu:      s.PackageLimit.Cpu,
			DiskSize: s.PackageLimit.DiskSize,
			Memory:   s.PackageLimit.Memory,
		},
	}
}

func (s NetworkInterface) toEntityNetworkInterface() lsentity.NetworkInterface {
	return lsentity.NetworkInterface{
		CreatedAt:     s.CreatedAt,
		FixedIp:       s.FixedIp,
		FloatingIp:    s.FloatingIp,
		FloatingIpId:  s.FloatingIpId,
		InterfaceType: s.InterfaceType,
		Mac:           s.Mac,
		NetworkUuid:   s.NetworkUuid,
		PortUuid:      s.PortUuid,
		Product:       s.Product,
		ServerUuid:    s.ServerUuid,
		Status:        s.Status,
		SubnetUuid:    s.SubnetUuid,
		Type:          s.Type,
		UpdatedAt:     s.UpdatedAt,
		Uuid:          s.Uuid,
	}
}

func (s ServerSecgroup) toEntityServerSecgroup() lsentity.ServerSecgroup {
	return lsentity.ServerSecgroup{
		Name: s.Name,
		Uuid: s.Uuid,
	}
}

func (s Flavor) toEntityFlavor() lsentity.Flavor {
	return lsentity.Flavor{
		Bandwidth:              s.Bandwidth,
		BandwidthUnit:          s.BandwidthUnit,
		Cpu:                    s.Cpu,
		CpuPlatformDescription: s.CpuPlatformDescription,
		FlavorId:               s.FlavorId,
		Gpu:                    s.Gpu,
		Group:                  s.Group,
		Memory:                 s.Memory,
		MetaData:               s.MetaData,
		Name:                   s.Name,
		RemainingVms:           s.RemainingVms,
		ZoneId:                 s.ZoneId,
	}
}

func (s Server) toEntityServer() *lsentity.Server {
	server := new(lsentity.Server)
	server.BootVolumeId = s.BootVolumeId
	server.CreatedAt = s.CreatedAt
	server.EncryptionVolume = s.EncryptionVolume
	server.Licence = s.Licence
	server.Location = s.Location
	server.Metadata = s.Metadata
	server.MigrateState = s.MigrateState
	server.Name = s.Name
	server.Product = s.Product
	server.ServerGroupId = s.ServerGroupId
	server.ServerGroupName = s.ServerGroupName
	server.SshKeyName = s.SshKeyName
	server.Status = s.Status
	server.StopBeforeMigrate = s.StopBeforeMigrate
	server.User = s.User
	server.Uuid = s.Uuid
	server.Image = s.Image.toEntityImage()
	server.Flavor = s.Flavor.toEntityFlavor()

	for _, secGroup := range s.SecGroups {
		server.SecGroups = append(server.SecGroups, secGroup.toEntityServerSecgroup())
	}

	for _, externalInterface := range s.ExternalInterfaces {
		server.ExternalInterfaces = append(server.ExternalInterfaces, externalInterface.toEntityNetworkInterface())
	}

	for _, internalInterface := range s.InternalInterfaces {
		server.InternalInterfaces = append(server.InternalInterfaces, internalInterface.toEntityNetworkInterface())
	}

	return server
}

type CreateServerResponse struct {
	Data Server `json:"data"`
}

func (s *CreateServerResponse) ToEntityServer() *lsentity.Server {
	return s.Data.toEntityServer()
}
