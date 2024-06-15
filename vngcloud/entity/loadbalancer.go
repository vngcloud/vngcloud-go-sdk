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
}
