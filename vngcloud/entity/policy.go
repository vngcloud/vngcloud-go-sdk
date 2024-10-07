package entity

type Policy struct {
	UUID             string
	Name             string
	Description      string
	RedirectPoolID   string
	RedirectPoolName string
	Action           string
	RedirectURL      string
	RedirectHTTPCode int
	KeepQueryString  bool
	Position         int
	L7Rules          []*L7Rule
	DisplayStatus    string
	CreatedAt        string
	UpdatedAt        string
	ProgressStatus   string
}

type L7Rule struct {
	UUID               string
	CompareType        string
	RuleValue          string
	RuleType           string
	ProvisioningStatus string
	OperatingStatus    string
}

type ListPolicies struct {
	Items []*Policy
}
