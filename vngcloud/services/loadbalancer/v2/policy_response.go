package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type responseData struct {
	UUID             string   `json:"uuid"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	RedirectPoolID   string   `json:"redirectPoolId"`
	RedirectPoolName string   `json:"redirectPoolName"`
	Action           string   `json:"action"`
	RedirectURL      string   `json:"redirectUrl"`
	RedirectHTTPCode int      `json:"redirectHttpCode"`
	KeepQueryString  bool     `json:"keepQueryString"`
	Position         int      `json:"position"`
	L7Rules          []l7Rule `json:"l7Rules"`
	DisplayStatus    string   `json:"displayStatus"`
	CreatedAt        string   `json:"createdAt"`
	UpdatedAt        string   `json:"updatedAt"`
	ProgressStatus   string   `json:"progressStatus"`
}

type l7Rule struct {
	UUID               string `json:"uuid"`
	CompareType        string `json:"compareType"`
	RuleValue          string `json:"ruleValue"`
	RuleType           string `json:"ruleType"`
	CreatedAt          string `json:"createdAt"`
	UpdatedAt          string `json:"updatedAt"`
	ProvisioningStatus string `json:"provisioningStatus"`
	OperatingStatus    string `json:"operatingStatus"`
}

func (s *l7Rule) toL7Rule() *lsentity.L7Rule {
	return &lsentity.L7Rule{
		UUID:               s.UUID,
		CompareType:        s.CompareType,
		RuleValue:          s.RuleValue,
		RuleType:           s.RuleType,
		ProvisioningStatus: s.ProvisioningStatus,
		OperatingStatus:    s.OperatingStatus,
	}
}

// ListPoliciesResponse

type ListPoliciesResponse struct {
	ListData  []responseData `json:"data"`
	Page      int            `json:"page"`
	PageSize  int            `json:"pageSize"`
	TotalPage int            `json:"totalPage"`
	TotalItem int            `json:"totalItem"`
}

func (s *ListPoliciesResponse) ToEntityListPolicies() *lsentity.ListPolicies {
	if s == nil || s.ListData == nil {
		return nil
	}

	items := make([]*lsentity.Policy, 0, len(s.ListData))
	for _, item := range s.ListData {
		l7Rule := make([]*lsentity.L7Rule, 0)
		for _, rule := range item.L7Rules {
			l7Rule = append(l7Rule, rule.toL7Rule())
		}
		items = append(items, &lsentity.Policy{
			UUID:             item.UUID,
			Name:             item.Name,
			Description:      item.Description,
			RedirectPoolID:   item.RedirectPoolID,
			RedirectPoolName: item.RedirectPoolName,
			Action:           item.Action,
			RedirectURL:      item.RedirectURL,
			RedirectHTTPCode: item.RedirectHTTPCode,
			KeepQueryString:  item.KeepQueryString,
			Position:         item.Position,
			DisplayStatus:    item.DisplayStatus,
			CreatedAt:        item.CreatedAt,
			UpdatedAt:        item.UpdatedAt,
			ProgressStatus:   item.ProgressStatus,
			L7Rules:          l7Rule,
		})
	}
	return &lsentity.ListPolicies{
		Items: items,
	}
}

// CreatePolicyResponse

type CreatePolicyResponse struct {
	UUID string `json:"uuid"`
}

func (s *CreatePolicyResponse) ToEntityPolicy() *lsentity.Policy {
	if s == nil {
		return nil
	}
	return &lsentity.Policy{
		UUID: s.UUID,
	}
}

// GetPolicyResponse

type GetPolicyResponse struct {
	Data responseData `json:"data"`
}

func (s *GetPolicyResponse) ToEntityPolicy() *lsentity.Policy {
	if s == nil {
		return nil
	}
	l7Rule := make([]*lsentity.L7Rule, 0)
	for _, rule := range s.Data.L7Rules {
		l7Rule = append(l7Rule, rule.toL7Rule())
	}
	return &lsentity.Policy{
		UUID:             s.Data.UUID,
		Name:             s.Data.Name,
		Description:      s.Data.Description,
		RedirectPoolID:   s.Data.RedirectPoolID,
		RedirectPoolName: s.Data.RedirectPoolName,
		Action:           s.Data.Action,
		RedirectURL:      s.Data.RedirectURL,
		RedirectHTTPCode: s.Data.RedirectHTTPCode,
		KeepQueryString:  s.Data.KeepQueryString,
		Position:         s.Data.Position,
		DisplayStatus:    s.Data.DisplayStatus,
		CreatedAt:        s.Data.CreatedAt,
		UpdatedAt:        s.Data.UpdatedAt,
		ProgressStatus:   s.Data.ProgressStatus,
		L7Rules:          l7Rule,
	}
}

// UpdatePolicyResponse and DeletePolicyResponse not have response data
