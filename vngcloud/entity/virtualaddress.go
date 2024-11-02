package entity

type VirtualAddress struct {
	Id              string
	Name            string
	EndpointAddress string
	VpcId           string
	SubnetId        string
	Description     string
	SubnetCidr      string
	VpcCidr         string
	AddressPairIps  []string
}

func (s *VirtualAddress) GetId() string {
	return s.Id
}

func (s *VirtualAddress) GetName() string {
	return s.Name
}

func (s *VirtualAddress) GetEndpointAddress() string {
	return s.EndpointAddress
}

func (s *VirtualAddress) GetVpcId() string {
	return s.VpcId
}

func (s *VirtualAddress) GetSubnetId() string {
	return s.SubnetId
}

func (s *VirtualAddress) GetDescription() string {
	return s.Description
}

func (s *VirtualAddress) GetSubnetCidr() string {
	return s.SubnetCidr
}

func (s *VirtualAddress) GetVpcCidr() string {
	return s.VpcCidr
}

func (s *VirtualAddress) GetAddressPairIps() []string {
	return s.AddressPairIps
}
