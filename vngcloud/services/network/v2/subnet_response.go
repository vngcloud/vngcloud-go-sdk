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
}

func (s *GetSubnetByIdResponse) ToEntitySubnet() *lsentity.Subnet {
	return &lsentity.Subnet{
		Id:                     s.UUID,
		NetworkId:              s.NetworkUuid,
		Name:                   s.Name,
		Status:                 s.Status,
		Cidr:                   s.Cidr,
		RouteTableId:           s.RouteTableUuid,
		InterfaceAclPolicyId:   s.InterfaceAclPolicyUuid,
		InterfaceAclPolicyName: s.InterfaceAclPolicyName,
	}
}
