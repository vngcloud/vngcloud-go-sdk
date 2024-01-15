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
	Address string
	Backup  bool
	Status  string
	Name    string
	UUID    string
}
