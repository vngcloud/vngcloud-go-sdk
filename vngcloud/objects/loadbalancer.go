package objects

type LoadBalancer struct {
	UUID     string
	Status   string
	Address  string
	Name     string
	SubnetID string
	Internal bool
}
