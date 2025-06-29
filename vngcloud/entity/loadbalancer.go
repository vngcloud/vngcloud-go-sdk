package entity

type LoadBalancer struct {
	UUID               string
	Name               string
	DisplayStatus      string
	Address            string
	PrivateSubnetID    string
	PrivateSubnetCidr  string
	Type               string
	DisplayType        string
	LoadBalancerSchema string
	PackageID          string
	Description        string
	Location           string
	CreatedAt          string
	UpdatedAt          string
	ProgressStatus     string
	Status             string
	SubnetID           string
	Internal           bool
	AutoScalable       bool
	ZoneID             string
}

type ListLoadBalancers struct {
	Items     []*LoadBalancer
	Page      int
	PageSize  int
	TotalPage int
	TotalItem int
}

func (s *LoadBalancer) GetId() string {
	return s.UUID
}

func (s *LoadBalancer) GetName() string {
	return s.Name
}

func (s *LoadBalancer) GetAddress() string {
	return s.Address
}

func (s *ListLoadBalancers) Len() int {
	return len(s.Items)
}

func (s *ListLoadBalancers) Empty() bool {
	return s.Len() < 1
}

func (s *ListLoadBalancers) Add(item *LoadBalancer) {
	s.Items = append(s.Items, item)
}

func (s *ListLoadBalancers) At(pidx int) *LoadBalancer {
	if pidx < 0 || pidx >= s.Len() {
		return nil
	}

	return s.Items[pidx]
}

type ListLoadBalancerPackages struct {
	Items []*LoadBalancerPackage
}

type LoadBalancerPackage struct {
	UUID             string
	Name             string
	Type             string
	ConnectionNumber int
	DataTransfer     int
	Mode             string
	LbType           string
	DisplayLbType    string
}
