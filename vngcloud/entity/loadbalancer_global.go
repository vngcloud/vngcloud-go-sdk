package entity

type GlobalLoadBalancer struct {
	CreatedAt   string                      `json:"createdAt"`
	UpdatedAt   string                      `json:"updatedAt"`
	DeletedAt   string                      `json:"deletedAt"`
	ID          string                      `json:"id"`
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	Status      string                      `json:"status"`
	Package     string                      `json:"package"`
	Type        string                      `json:"type"`
	UserId      int                         `json:"userId"`
	Vips        []*GlobalLoadBalancerVIP    `json:"vips"`
	Domains     []*GlobalLoadBalancerDomain `json:"domains"`
}

type GlobalLoadBalancerVIP struct {
	ID                   int    `json:"id"`
	CreatedAt            string `json:"createdAt"`
	UpdatedAt            string `json:"updatedAt"`
	DeletedAt            string `json:"deletedAt"`
	Address              string `json:"address"`
	Status               string `json:"status"`
	Region               string `json:"region"`
	GlobalLoadBalancerID string `json:"globalLoadBalancerId"`
}

type GlobalLoadBalancerDomain struct {
	CreatedAt            string `json:"createdAt"`
	UpdatedAt            string `json:"updatedAt"`
	DeletedAt            string `json:"deletedAt"`
	ID                   int    `json:"id"`
	Hostname             string `json:"hostname"`
	Status               string `json:"status"`
	GlobalLoadBalancerID string `json:"globalLoadBalancerId"`
	DNSHostedZoneID      string `json:"dnsHostedZoneId"`
	DNSServerID          string `json:"dnsServerId"`
}

type ListGlobalLoadBalancers struct {
	Items  []*GlobalLoadBalancer `json:"items"`
	Limit  int                   `json:"limit"`
	Total  int                   `json:"total"`
	Offset int                   `json:"offset"`
}

// type ListGlobalLoadBalancerPackages struct {
// 	Items []*LoadBalancerPackage
// }

// type GlobalLoadBalancerPackage struct {
// 	UUID             string
// 	Name             string
// 	Type             string
// 	ConnectionNumber int
// 	DataTransfer     int
// 	Mode             string
// 	LbType           string
// 	DisplayLbType    string
// }
