package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type CreateVirtualAddressCrossProjectResponse struct {
	Data struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		NetworkName string `json:"networkName"`
		NetworkCIDR string `json:"networkCIDR"`
		SubnetCIDR  string `json:"subnetCIDR"`
		SubnetName  string `json:"subnetName"`
		UUID        string `json:"uuid"`
		NetworkID   string `json:"networkId"`
		SubnetID    string `json:"subnetId"`
		IPAddress   string `json:"ipAddress"`
		CreatedAt   string `json:"createdAt"`
	} `json:"data"`
}

func (s *CreateVirtualAddressCrossProjectResponse) ToEntityVirtualAddress() *lsentity.VirtualAddress {
	return &lsentity.VirtualAddress{
		Id:              s.Data.UUID,
		Name:            s.Data.Name,
		EndpointAddress: s.Data.IPAddress,
		VpcId:           s.Data.NetworkID,
		SubnetId:        s.Data.SubnetID,
		Description:     s.Data.Description,
		SubnetCidr:      s.Data.SubnetCIDR,
		VpcCidr:         s.Data.NetworkCIDR,
	}
}
