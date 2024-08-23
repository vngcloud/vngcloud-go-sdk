package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type GetAllAddressPairsByVirtualSubnetIdResponse struct {
	Data []*AddressPairResponse `json:"data"`
}

type AddressPairResponse struct {
	UUID               string `json:"uuid"`
	VirtualIpAddressId string `json:"virtualIpAddressId"`
	VirtualSubnetId    string `json:"virtualSubnetId"`
	NetworkInterfaceIp string `json:"networkInterfaceIp"`
	NetworkInterfaceId string `json:"networkInterfaceId"`

	CIDR string `json:"cidr"`
	// ID              string `json:"id"`
	// BackendSubnetId int    `json:"backendSubnetId"`
	// ProjectId       string `json:"projectId"`
	// CreatedAt       string `json:"createdAt"`
	// DeletedAt       string `json:"deletedAt"`
}

func (s *GetAllAddressPairsByVirtualSubnetIdResponse) ToListAddressPair() []*lsentity.AddressPair {
	addressPairs := make([]*lsentity.AddressPair, 0)
	for _, addressPair := range s.Data {
		addressPairs = append(addressPairs, &lsentity.AddressPair{
			Id:                 addressPair.UUID,
			VirtualIpAddressId: addressPair.VirtualIpAddressId,
			VirtualSubnetId:    addressPair.VirtualSubnetId,
			NetworkInterfaceIp: addressPair.NetworkInterfaceIp,
			NetworkInterfaceId: addressPair.NetworkInterfaceId,
			CIDR:               addressPair.CIDR,
		})
	}
	return addressPairs
}

type SetAddressPairInVirtualSubnetResponse struct {
	Data *AddressPairResponse `json:"data"`
}

func (s *SetAddressPairInVirtualSubnetResponse) ToAddressPair() *lsentity.AddressPair {
	return &lsentity.AddressPair{
		Id:                 s.Data.UUID,
		VirtualIpAddressId: s.Data.VirtualIpAddressId,
		VirtualSubnetId:    s.Data.VirtualSubnetId,
		NetworkInterfaceIp: s.Data.NetworkInterfaceIp,
		NetworkInterfaceId: s.Data.NetworkInterfaceId,
		CIDR:               s.Data.CIDR,
	}
}
