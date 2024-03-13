package server

import "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"

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
		SecGroups          []ServerSecGroup   `json:"secGroups"`
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

	ServerSecGroup struct {
		Name string `json:"name"`
		Uuid string `json:"uuid"`
	}
)

type GetResponse struct {
	Data Server `json:"data"`
}

func (s *GetResponse) ToServerObject() *objects.Server {
	if s == nil {
		return nil
	}

	server := new(objects.Server)
	server.BootVolumeId = s.Data.BootVolumeId
	server.CreatedAt = s.Data.CreatedAt
	server.EncryptionVolume = s.Data.EncryptionVolume
	server.Licence = s.Data.Licence
	server.Location = s.Data.Location
	server.Metadata = s.Data.Metadata
	server.MigrateState = s.Data.MigrateState
	server.Name = s.Data.Name
	server.Product = s.Data.Product
	server.ServerGroupId = s.Data.ServerGroupId
	server.ServerGroupName = s.Data.ServerGroupName
	server.SshKeyName = s.Data.SshKeyName
	server.Status = s.Data.Status
	server.StopBeforeMigrate = s.Data.StopBeforeMigrate
	server.User = s.Data.User
	server.Uuid = s.Data.Uuid
	server.Image = objects.Image{
		FlavorZoneIds: s.Data.Image.FlavorZoneIds,
		Id:            s.Data.Image.Id,
		ImageType:     s.Data.Image.ImageType,
		ImageVersion:  s.Data.Image.ImageVersion,
		Licence:       s.Data.Image.Licence,
		PackageLimit: objects.PackageLimit{
			Cpu:      s.Data.Image.PackageLimit.Cpu,
			DiskSize: s.Data.Image.PackageLimit.DiskSize,
			Memory:   s.Data.Image.PackageLimit.Memory,
		},
	}
	server.Flavor = objects.Flavor{
		Bandwidth:              s.Data.Flavor.Bandwidth,
		BandwidthUnit:          s.Data.Flavor.BandwidthUnit,
		Cpu:                    s.Data.Flavor.Cpu,
		CpuPlatformDescription: s.Data.Flavor.CpuPlatformDescription,
		FlavorId:               s.Data.Flavor.FlavorId,
		Gpu:                    s.Data.Flavor.Gpu,
		Group:                  s.Data.Flavor.Group,
		Memory:                 s.Data.Flavor.Memory,
		MetaData:               s.Data.Flavor.MetaData,
		Name:                   s.Data.Flavor.Name,
		RemainingVms:           s.Data.Flavor.RemainingVms,
		ZoneId:                 s.Data.Flavor.ZoneId,
	}

	for _, secGroup := range s.Data.SecGroups {
		server.SecGroups = append(server.SecGroups, objects.ServerSecGroup{
			Name: secGroup.Name,
			Uuid: secGroup.Uuid,
		})
	}

	for _, externalInterface := range s.Data.ExternalInterfaces {
		server.ExternalInterfaces = append(server.ExternalInterfaces, objects.NetworkInterface{
			CreatedAt:     externalInterface.CreatedAt,
			FixedIp:       externalInterface.FixedIp,
			FloatingIp:    externalInterface.FloatingIp,
			FloatingIpId:  externalInterface.FloatingIpId,
			InterfaceType: externalInterface.InterfaceType,
			Mac:           externalInterface.Mac,
			NetworkUuid:   externalInterface.NetworkUuid,
			PortUuid:      externalInterface.PortUuid,
			Product:       externalInterface.Product,
			ServerUuid:    externalInterface.ServerUuid,
			Status:        externalInterface.Status,
			SubnetUuid:    externalInterface.SubnetUuid,
			Type:          externalInterface.Type,
			UpdatedAt:     externalInterface.UpdatedAt,
			Uuid:          externalInterface.Uuid,
		})
	}

	for _, internalInterface := range s.Data.InternalInterfaces {
		server.InternalInterfaces = append(server.InternalInterfaces, objects.NetworkInterface{
			CreatedAt:     internalInterface.CreatedAt,
			FixedIp:       internalInterface.FixedIp,
			FloatingIp:    internalInterface.FloatingIp,
			FloatingIpId:  internalInterface.FloatingIpId,
			InterfaceType: internalInterface.InterfaceType,
			Mac:           internalInterface.Mac,
			NetworkUuid:   internalInterface.NetworkUuid,
			PortUuid:      internalInterface.PortUuid,
			Product:       internalInterface.Product,
			ServerUuid:    internalInterface.ServerUuid,
			Status:        internalInterface.Status,
			SubnetUuid:    internalInterface.SubnetUuid,
			Type:          internalInterface.Type,
			UpdatedAt:     internalInterface.UpdatedAt,
			Uuid:          internalInterface.Uuid,
		})
	}

	return server
}

type CreateResponse struct {
	Data Server `json:"data"`
}

func (s *CreateResponse) ToServerObject() *objects.Server {
	if s == nil {
		return nil
	}
	return s.Data.toServerObject()
}

func (s Server) toServerObject() *objects.Server {
	server := new(objects.Server)
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
	server.Image = objects.Image{
		FlavorZoneIds: s.Image.FlavorZoneIds,
		Id:            s.Image.Id,
		ImageType:     s.Image.ImageType,
		ImageVersion:  s.Image.ImageVersion,
		Licence:       s.Image.Licence,
		PackageLimit: objects.PackageLimit{
			Cpu:      s.Image.PackageLimit.Cpu,
			DiskSize: s.Image.PackageLimit.DiskSize,
			Memory:   s.Image.PackageLimit.Memory,
		},
	}
	server.Flavor = objects.Flavor{
		Bandwidth:              s.Flavor.Bandwidth,
		BandwidthUnit:          s.Flavor.BandwidthUnit,
		Cpu:                    s.Flavor.Cpu,
		CpuPlatformDescription: s.Flavor.CpuPlatformDescription,
		FlavorId:               s.Flavor.FlavorId,
		Gpu:                    s.Flavor.Gpu,
		Group:                  s.Flavor.Group,
		Memory:                 s.Flavor.Memory,
		MetaData:               s.Flavor.MetaData,
		Name:                   s.Flavor.Name,
		RemainingVms:           s.Flavor.RemainingVms,
		ZoneId:                 s.Flavor.ZoneId,
	}

	for _, secGroup := range s.SecGroups {
		server.SecGroups = append(server.SecGroups, objects.ServerSecGroup{
			Name: secGroup.Name,
			Uuid: secGroup.Uuid,
		})
	}

	for _, externalInterface := range s.ExternalInterfaces {
		server.ExternalInterfaces = append(server.ExternalInterfaces, objects.NetworkInterface{
			CreatedAt:     externalInterface.CreatedAt,
			FixedIp:       externalInterface.FixedIp,
			FloatingIp:    externalInterface.FloatingIp,
			FloatingIpId:  externalInterface.FloatingIpId,
			InterfaceType: externalInterface.InterfaceType,
			Mac:           externalInterface.Mac,
			NetworkUuid:   externalInterface.NetworkUuid,
			PortUuid:      externalInterface.PortUuid,
			Product:       externalInterface.Product,
			ServerUuid:    externalInterface.ServerUuid,
			Status:        externalInterface.Status,
			SubnetUuid:    externalInterface.SubnetUuid,
			Type:          externalInterface.Type,
			UpdatedAt:     externalInterface.UpdatedAt,
			Uuid:          externalInterface.Uuid,
		})
	}

	for _, internalInterface := range s.InternalInterfaces {
		server.InternalInterfaces = append(server.InternalInterfaces, objects.NetworkInterface{
			CreatedAt:     internalInterface.CreatedAt,
			FixedIp:       internalInterface.FixedIp,
			FloatingIp:    internalInterface.FloatingIp,
			FloatingIpId:  internalInterface.FloatingIpId,
			InterfaceType: internalInterface.InterfaceType,
			Mac:           internalInterface.Mac,
			NetworkUuid:   internalInterface.NetworkUuid,
			PortUuid:      internalInterface.PortUuid,
			Product:       internalInterface.Product,
			ServerUuid:    internalInterface.ServerUuid,
			Status:        internalInterface.Status,
			SubnetUuid:    internalInterface.SubnetUuid,
			Type:          internalInterface.Type,
			UpdatedAt:     internalInterface.UpdatedAt,
			Uuid:          internalInterface.Uuid,
		})
	}

	return server
}

type ListResponse struct {
	Data      []Server `json:"listData"`
	Page      int      `json:"page"`
	PageSize  int      `json:"pageSize"`
	TotalPage int      `json:"totalPage"`
	TotalItem int      `json:"totalItem"`
}

func (s *ListResponse) ToServerList() ([]*objects.Server, error) {
	var servers []*objects.Server

	if s == nil || s.Data == nil || len(s.Data) < 1 {
		return servers, nil
	}

	for _, server := range s.Data {
		servers = append(servers, server.toServerObject())
	}

	return servers, nil
}
