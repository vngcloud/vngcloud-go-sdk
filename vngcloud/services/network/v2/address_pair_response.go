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
		})
	}
	return addressPairs
}
