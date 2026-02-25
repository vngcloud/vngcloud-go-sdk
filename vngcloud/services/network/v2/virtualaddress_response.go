package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type VirtualAddressDataResponse struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	NetworkName    string   `json:"networkName"`
	NetworkCIDR    string   `json:"networkCIDR"`
	SubnetCIDR     string   `json:"subnetCIDR"`
	SubnetName     string   `json:"subnetName"`
	UUID           string   `json:"uuid"`
	NetworkID      string   `json:"networkId"`
	SubnetID       string   `json:"subnetId"`
	IPAddress      string   `json:"ipAddress"`
	CreatedAt      string   `json:"createdAt"`
	AddressPairIps []string `json:"addressPairIps"`
}

func (s *VirtualAddressDataResponse) toEntityVirtualAddress() *lsentity.VirtualAddress {
	return &lsentity.VirtualAddress{
		Id:              s.UUID,
		Name:            s.Name,
		EndpointAddress: s.IPAddress,
		VpcId:           s.NetworkID,
		SubnetId:        s.SubnetID,
		Description:     s.Description,
		SubnetCidr:      s.SubnetCIDR,
		VpcCidr:         s.NetworkCIDR,
		AddressPairIps:  s.AddressPairIps,
	}
}

// Response struct for API create virtual address cross project

type CreateVirtualAddressCrossProjectResponse struct {
	Data VirtualAddressDataResponse `json:"data"`
}

func (s *CreateVirtualAddressCrossProjectResponse) ToEntityVirtualAddress() *lsentity.VirtualAddress {
	return s.Data.toEntityVirtualAddress()
}

// Response struct for API get virtual address by ID
type GetVirtualAddressByIdResponse struct {
	Data VirtualAddressDataResponse `json:"data"`
}

func (s *GetVirtualAddressByIdResponse) ToEntityVirtualAddress() *lsentity.VirtualAddress {
	return s.Data.toEntityVirtualAddress()
}

// Response struct for API list address pair by virtual address ID
type ListAddressPairsByVirtualAddressIdResponse struct {
	Data []AddressPairResponse `json:"data"`
}

func (s *ListAddressPairsByVirtualAddressIdResponse) ToEntityListAddressPairs() *lsentity.ListAddressPairs {
	addressPairs := make([]*lsentity.AddressPair, 0, len(s.Data))
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
	return &lsentity.ListAddressPairs{Items: addressPairs}
}
