package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type GetSubnetByIdResponse struct {
	UUID                   string `json:"uuid"`
	CreatedAt              string `json:"createdAt"`
	UpdatedAt              string `json:"updatedAt,omitempty"`
	Status                 string `json:"status"`
	Cidr                   string `json:"cidr"`
	NetworkUuid            string `json:"networkUuid"`
	RouteTableUuid         string `json:"routeTableUuid,omitempty"`
	Name                   string `json:"name"`
	InterfaceAclPolicyUuid string `json:"interfaceAclPolicyUuid,omitempty"`
	InterfaceAclPolicyName string `json:"interfaceAclPolicyName,omitempty"`
	SecondarySubnets       []struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
		Cidr string `json:"cidr"`
	} `json:"secondarySubnets,omitempty"`
	Zone struct {
		UUID string `json:"uuid"`
	} `json:"zone"`
}

func (s *GetSubnetByIdResponse) ToEntitySubnet() *lsentity.Subnet {
	secondaryRange := make([]lsentity.SubnetSecondaryRange, 0, len(s.SecondarySubnets))
	for _, sr := range s.SecondarySubnets {
		secondaryRange = append(secondaryRange, lsentity.SubnetSecondaryRange{
			Id:   sr.UUID,
			Name: sr.Name,
			Cidr: sr.Cidr,
		})
	}
	return &lsentity.Subnet{
		Id:                     s.UUID,
		NetworkId:              s.NetworkUuid,
		Name:                   s.Name,
		Status:                 s.Status,
		Cidr:                   s.Cidr,
		RouteTableId:           s.RouteTableUuid,
		InterfaceAclPolicyId:   s.InterfaceAclPolicyUuid,
		InterfaceAclPolicyName: s.InterfaceAclPolicyName,
		SecondarySubnets:       secondaryRange,
		ZoneID:                 s.Zone.UUID,
	}
}

// --------------------------------------------------------
type UpdateSubnetByIdResponse struct {
	Data GetSubnetByIdResponse `json:"data"`
}

func (s *UpdateSubnetByIdResponse) ToEntitySubnet() *lsentity.Subnet {
	return s.Data.ToEntitySubnet()
}
