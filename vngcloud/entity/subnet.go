package entity

type Subnet struct {
	Id                     string
	NetworkId              string
	Name                   string
	Status                 string
	Cidr                   string
	RouteTableId           string
	InterfaceAclPolicyId   string
	InterfaceAclPolicyName string
}
