package entity

import "encoding/json"

type Pool struct {
	UUID              string
	Name              string
	Protocol          string
	Description       string
	LoadBalanceMethod string
	Status            string
	Stickiness        bool
	TLSEncryption     bool
	Members           *ListMembers
	HealthMonitor     *HealthMonitor
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

type HealthMonitor struct {
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

func (s *HealthMonitor) String() string {
	// parse to string and return
	out, err := json.Marshal(s)
	if err != nil {
		return "Error parsing to string"
	}
	return string(out)
}

type ListPools struct {
	Items []*Pool
}

type ListMembers struct {
	Items []*Member
}

func (s *Pool) GetId() string {
	return s.UUID
}

func (s *ListPools) Add(pools ...*Pool) {
	s.Items = append(s.Items, pools...)
}

func (s *ListMembers) Add(members ...*Member) {
	s.Items = append(s.Items, members...)
}

func (s *ListPools) Len() int {
	return len(s.Items)
}

func (s *ListPools) Empty() bool {
	return s.Len() < 1
}

func (s *ListPools) At(index int) *Pool {
	if index < 0 || index >= s.Len() {
		return nil
	}

	return s.Items[index]
}
