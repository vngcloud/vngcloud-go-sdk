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
	Members           []Member
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
