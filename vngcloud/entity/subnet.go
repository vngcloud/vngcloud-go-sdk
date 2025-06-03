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
	SecondarySubnets       []SubnetSecondaryRange
	ZoneID                 string
}

type SubnetSecondaryRange struct {
	Id   string
	Name string
	Cidr string
}
