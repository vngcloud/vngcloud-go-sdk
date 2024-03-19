package objects

type Pool struct {
	UUID              string
	Name              string
	Protocol          string
	Description       string
	LoadBalanceMethod string
	Status            string
	Stickiness        bool
	TLSEncryption     bool
	Members           []*Member
	HealthMonitor     *PoolHealthMonitor
}

type Member struct {
	UUID           string
	Address        string
	ProtocolPort   int
	Weight         int
	MonitorPort    int
	SubnetID       string
	Name           string
	PoolID         string
	TypeCreate     string
	Backup         bool
	DisplayStatus  string
	CreatedAt      string
	UpdatedAt      string
	CreatedBy      string
	ProgressStatus string
}

type PoolHealthMonitor struct {
	Timeout             int     `json:"timeout"`
	CreatedAt           string  `json:"createdAt"`
	UpdatedAt           string  `json:"updatedAt"`
	DomainName          *string `json:"domainName"`
	HttpVersion         *string `json:"httpVersion"`
	HealthCheckProtocol string  `json:"healthCheckProtocol"`
	Interval            int     `json:"interval"`
	HealthyThreshold    int     `json:"healthyThreshold"`
	UnhealthyThreshold  int     `json:"unhealthyThreshold"`
	HealthCheckMethod   *string `json:"healthCheckMethod"`
	HealthCheckPath     *string `json:"healthCheckPath"`
	SuccessCode         *string `json:"successCode"`
	ProgressStatus      string  `json:"progressStatus"`
	DisplayStatus       string  `json:"displayStatus"`
}
