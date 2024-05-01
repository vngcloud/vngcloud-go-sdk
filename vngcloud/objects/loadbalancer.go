package objects

type LoadBalancer struct {
	UUID               string `json:"uuid"`
	Name               string `json:"name"`
	DisplayStatus      string `json:"displayStatus"`
	Address            string `json:"address"`
	PrivateSubnetID    string `json:"privateSubnetId"`
	PrivateSubnetCidr  string `json:"privateSubnetCidr"`
	Type               string `json:"type"`
	DisplayType        string `json:"displayType"`
	LoadBalancerSchema string `json:"loadBalancerSchema"`
	PackageID          string `json:"packageId"`
	Description        string `json:"description"`
	Location           string `json:"location"`
	CreatedAt          string `json:"createdAt"`
	UpdatedAt          string `json:"updatedAt"`
	ProgressStatus     string `json:"progressStatus"`
	// PackageInfo        struct {
	// 	PackageID        string `json:"packageId"`
	// 	ConnectionNumber int    `json:"connectionNumber"`
	// 	DataTransfer     int    `json:"dataTransfer"`
	// 	Name             string `json:"name"`
	// } `json:"packageInfo"`

	Status   string
	SubnetID string
	Internal bool
}
