package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type GetAllAddressPairByVirtualSubnetIdResponse struct {
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

func (s *AddressPairResponse) toAddressPair() *lsentity.AddressPair {
	return &lsentity.AddressPair{
		Id:                 s.UUID,
		VirtualIpAddressId: s.VirtualIpAddressId,
		VirtualSubnetId:    s.VirtualSubnetId,
		NetworkInterfaceIp: s.NetworkInterfaceIp,
		NetworkInterfaceId: s.NetworkInterfaceId,
		CIDR:               s.CIDR,
	}
}

func (s *GetAllAddressPairByVirtualSubnetIdResponse) ToListAddressPair() []*lsentity.AddressPair {
	addressPairs := make([]*lsentity.AddressPair, 0, len(s.Data))
	for _, addressPair := range s.Data {
		addressPairs = append(addressPairs, addressPair.toAddressPair())
	}
	return addressPairs
}

type SetAddressPairInVirtualSubnetResponse struct {
	Data *AddressPairResponse `json:"data"`
}

func (s *SetAddressPairInVirtualSubnetResponse) ToAddressPair() *lsentity.AddressPair {
	return s.Data.toAddressPair()
}

type CreateAddressPairResponse struct {
	Data AddressPairResponse `json:"data"`
}

func (s *CreateAddressPairResponse) ToAddressPair() *lsentity.AddressPair {
	return s.Data.toAddressPair()
}
