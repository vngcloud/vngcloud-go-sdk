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

type GlobalPool struct {
	CreatedAt            string  `json:"createdAt"`
	UpdatedAt            string  `json:"updatedAt"`
	DeletedAt            *string `json:"deletedAt"`
	ID                   string  `json:"id"`
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	GlobalLoadBalancerID string  `json:"globalLoadBalancerId"`
	Algorithm            string  `json:"algorithm"`
	StickySession        *string `json:"stickySession"`
	TLSEnabled           *string `json:"tlsEnabled"`
	Protocol             string  `json:"protocol"`
	Status               string  `json:"status"`
}

type GlobalPoolHealthMonitor struct {
	CreatedAt            string  `json:"createdAt"`
	UpdatedAt            string  `json:"updatedAt"`
	DeletedAt            *string `json:"deletedAt"`
	ID                   string  `json:"id"`
	GlobalPoolID         string  `json:"globalPoolId"`
	GlobalLoadBalancerID string  `json:"globalLoadBalancerId"`
	Protocol             string  `json:"protocol"`
	Path                 *string `json:"path"`
	Timeout              int     `json:"timeout"`
	IntervalTime         int     `json:"intervalTime"`
	HealthyThreshold     int     `json:"healthyThreshold"`
	UnhealthyThreshold   int     `json:"unhealthyThreshold"`
	HTTPVersion          *string `json:"httpVersion"`
	HTTPMethod           *string `json:"httpMethod"`
	DomainName           *string `json:"domainName"`
	SuccessCode          *string `json:"successCode"`
	Status               string  `json:"status"`
}

type GlobalPoolMember struct {
	CreatedAt            string             `json:"createdAt"`
	UpdatedAt            string             `json:"updatedAt"`
	DeletedAt            *string            `json:"deletedAt"`
	ID                   string             `json:"id"`
	Name                 string             `json:"name"`
	Description          string             `json:"description"`
	Region               string             `json:"region"`
	GlobalPoolID         string             `json:"globalPoolId"`
	GlobalLoadBalancerID string             `json:"globalLoadBalancerId"`
	TrafficDial          int                `json:"trafficDial"`
	VpcID                string             `json:"vpcId"`
	Status               string             `json:"status"`
	Members              *ListGlobalMembers `json:"members"`
}

type ListGlobalMembers struct {
	Items []*GlobalPoolMemberDetail
}

type GlobalPoolMemberDetail struct {
	CreatedAt            string  `json:"createdAt"`
	UpdatedAt            string  `json:"updatedAt"`
	DeletedAt            *string `json:"deletedAt"`
	ID                   string  `json:"id"`
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	GlobalPoolMemberID   string  `json:"globalPoolMemberId"`
	GlobalLoadBalancerID string  `json:"globalLoadBalancerId"`
	SubnetID             string  `json:"subnetId"`
	Address              string  `json:"address"`
	Weight               int     `json:"weight"`
	Port                 int     `json:"port"`
	MonitorPort          int     `json:"monitorPort"`
	BackupRole           bool    `json:"backupRole"`
	Status               string  `json:"status"`
}

type ListGlobalPools struct {
	Items []*GlobalPool
}

type ListGlobalPoolMembers struct {
	Items []*GlobalPoolMember
}

// --------------------------------------------------------

type GlobalListener struct {
	CreatedAt            string  `json:"createdAt"`
	UpdatedAt            string  `json:"updatedAt"`
	DeletedAt            *string `json:"deletedAt"`
	ID                   string  `json:"id"`
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Protocol             string  `json:"protocol"`
	Port                 int     `json:"port"`
	GlobalLoadBalancerID string  `json:"globalLoadBalancerId"`
	GlobalPoolID         string  `json:"globalPoolId"`
	TimeoutClient        int     `json:"timeoutClient"`
	TimeoutMember        int     `json:"timeoutMember"`
	TimeoutConnection    int     `json:"timeoutConnection"`
	AllowedCidrs         string  `json:"allowedCidrs"`
	Headers              *string `json:"headers"`
	Status               string  `json:"status"`
}

type ListGlobalListeners struct {
	Items []*GlobalListener
}
