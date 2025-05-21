package v2

import lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"

type (
	PolicyAction      string
	PolicyCompareType string
	PolicyRuleType    string
)

const (
	PolicyActionREJECT         PolicyAction = "REJECT"
	PolicyActionREDIRECTTOURL  PolicyAction = "REDIRECT_TO_URL"
	PolicyActionREDIRECTTOPOOL PolicyAction = "REDIRECT_TO_POOL"

	PolicyCompareTypeCONTAINS   PolicyCompareType = "CONTAINS"
	PolicyCompareTypeEQUALS     PolicyCompareType = "EQUAL_TO"
	PolicyCompareTypeREGEX      PolicyCompareType = "REGEX"
	PolicyCompareTypeSTARTSWITH PolicyCompareType = "STARTS_WITH"
	PolicyCompareTypeENDSWITH   PolicyCompareType = "ENDS_WITH"

	PolicyRuleTypeHOSTNAME PolicyRuleType = "HOST_NAME"
	PolicyRuleTypePATH     PolicyRuleType = "PATH"
)

// create policy request
func NewCreatePolicyRequest(lbID, lisID string) ICreatePolicyRequest {
	return &CreatePolicyRequest{
		LoadBalancerCommon: lscommon.LoadBalancerCommon{LoadBalancerId: lbID},
		ListenerCommon:     lscommon.ListenerCommon{ListenerId: lisID},
	}
}

var _ ICreatePolicyRequest = &CreatePolicyRequest{}

type CreatePolicyRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
	lscommon.ListenerCommon

	Name             string          `json:"name"`
	Action           PolicyAction    `json:"action"`
	Rules            []L7RuleRequest `json:"rules"`
	RedirectPoolID   string          `json:"redirectPoolId"`
	RedirectURL      string          `json:"redirectUrl"`
	RedirectHTTPCode int             `json:"redirectHttpCode"`
	KeepQueryString  bool            `json:"keepQueryString"`
}

type L7RuleRequest struct {
	CompareType PolicyCompareType `json:"compareType"`
	RuleType    PolicyRuleType    `json:"ruleType"`
	RuleValue   string            `json:"ruleValue"`
}

func (s *CreatePolicyRequest) ToRequestBody() interface{} {
	switch s.Action {
	case PolicyActionREJECT:
		return struct {
			Name   string          `json:"name"`
			Action string          `json:"action"`
			Rules  []L7RuleRequest `json:"rules"`
		}{
			Name:   s.Name,
			Action: string(s.Action),
			Rules:  s.Rules,
		}
	case PolicyActionREDIRECTTOURL:
		return struct {
			Name             string          `json:"name"`
			Action           string          `json:"action"`
			Rules            []L7RuleRequest `json:"rules"`
			RedirectURL      string          `json:"redirectUrl"`
			RedirectHTTPCode int             `json:"redirectHttpCode"`
			KeepQueryString  bool            `json:"keepQueryString"`
		}{
			Name:             s.Name,
			Action:           string(s.Action),
			Rules:            s.Rules,
			RedirectURL:      s.RedirectURL,
			RedirectHTTPCode: s.RedirectHTTPCode,
			KeepQueryString:  s.KeepQueryString,
		}
	case PolicyActionREDIRECTTOPOOL:
		return struct {
			Name           string          `json:"name"`
			Action         string          `json:"action"`
			Rules          []L7RuleRequest `json:"rules"`
			RedirectPoolID string          `json:"redirectPoolId"`
		}{
			Name:           s.Name,
			Action:         string(s.Action),
			Rules:          s.Rules,
			RedirectPoolID: s.RedirectPoolID,
		}
	}
	return nil
}

func (s *CreatePolicyRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":             s.Name,
		"action":           string(s.Action),
		"rules":            s.Rules,
		"redirectPoolId":   s.RedirectPoolID,
		"redirectUrl":      s.RedirectURL,
		"redirectHttpCode": s.RedirectHTTPCode,
		"keepQueryString":  s.KeepQueryString,
	}
}

func (s *CreatePolicyRequest) WithName(pname string) ICreatePolicyRequest {
	s.Name = pname
	return s
}

func (s *CreatePolicyRequest) WithAction(paction PolicyAction) ICreatePolicyRequest {
	s.Action = paction
	return s
}

func (s *CreatePolicyRequest) WithRules(prules ...L7RuleRequest) ICreatePolicyRequest {
	s.Rules = prules
	return s
}

func (s *CreatePolicyRequest) WithRedirectPoolId(predirectPoolId string) ICreatePolicyRequest {
	s.RedirectPoolID = predirectPoolId
	return s
}

func (s *CreatePolicyRequest) WithRedirectURL(predirectURL string) ICreatePolicyRequest {
	s.RedirectURL = predirectURL
	return s
}

func (s *CreatePolicyRequest) WithRedirectHTTPCode(predirectHTTPCode int) ICreatePolicyRequest {
	s.RedirectHTTPCode = predirectHTTPCode
	return s
}

func (s *CreatePolicyRequest) WithKeepQueryString(pkeepQueryString bool) ICreatePolicyRequest {
	s.KeepQueryString = pkeepQueryString
	return s
}

func (s *CreatePolicyRequest) AddUserAgent(pagent ...string) ICreatePolicyRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

// update policy request
func NewUpdatePolicyRequest(lbID, lisID, policyID string) IUpdatePolicyRequest {
	return &UpdatePolicyRequest{
		LoadBalancerCommon: lscommon.LoadBalancerCommon{LoadBalancerId: lbID},
		ListenerCommon:     lscommon.ListenerCommon{ListenerId: lisID},
		PolicyCommon:       lscommon.PolicyCommon{PolicyId: policyID},
	}
}

var _ IUpdatePolicyRequest = &UpdatePolicyRequest{}

type UpdatePolicyRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
	lscommon.ListenerCommon
	lscommon.PolicyCommon

	Action           PolicyAction    `json:"action"`
	Rules            []L7RuleRequest `json:"rules"`
	RedirectPoolID   string          `json:"redirectPoolId"`
	RedirectURL      string          `json:"redirectUrl"`
	RedirectHTTPCode int             `json:"redirectHttpCode"`
	KeepQueryString  bool            `json:"keepQueryString"`
}

func (s *UpdatePolicyRequest) ToRequestBody() interface{} {
	switch s.Action {
	case PolicyActionREJECT:
		return struct {
			Action string          `json:"action"`
			Rules  []L7RuleRequest `json:"rules"`
		}{
			Action: string(s.Action),
			Rules:  s.Rules,
		}
	case PolicyActionREDIRECTTOURL:
		return struct {
			Action           string          `json:"action"`
			Rules            []L7RuleRequest `json:"rules"`
			RedirectURL      string          `json:"redirectUrl"`
			RedirectHTTPCode int             `json:"redirectHttpCode"`
			KeepQueryString  bool            `json:"keepQueryString"`
		}{
			Action:           string(s.Action),
			Rules:            s.Rules,
			RedirectURL:      s.RedirectURL,
			RedirectHTTPCode: s.RedirectHTTPCode,
			KeepQueryString:  s.KeepQueryString,
		}
	case PolicyActionREDIRECTTOPOOL:
		return struct {
			Action         string          `json:"action"`
			Rules          []L7RuleRequest `json:"rules"`
			RedirectPoolID string          `json:"redirectPoolId"`
		}{
			Action:         string(s.Action),
			Rules:          s.Rules,
			RedirectPoolID: s.RedirectPoolID,
		}
	}
	return nil
}

func (s *UpdatePolicyRequest) WithAction(paction PolicyAction) IUpdatePolicyRequest {
	s.Action = paction
	return s
}

func (s *UpdatePolicyRequest) WithRules(prules ...L7RuleRequest) IUpdatePolicyRequest {
	s.Rules = prules
	return s
}

func (s *UpdatePolicyRequest) WithRedirectPoolID(predirectPoolId string) IUpdatePolicyRequest {
	s.RedirectPoolID = predirectPoolId
	return s
}

func (s *UpdatePolicyRequest) WithRedirectURL(predirectURL string) IUpdatePolicyRequest {
	s.RedirectURL = predirectURL
	return s
}

func (s *UpdatePolicyRequest) WithRedirectHTTPCode(predirectHTTPCode int) IUpdatePolicyRequest {
	s.RedirectHTTPCode = predirectHTTPCode
	return s
}

func (s *UpdatePolicyRequest) WithKeepQueryString(pkeepQueryString bool) IUpdatePolicyRequest {
	s.KeepQueryString = pkeepQueryString
	return s
}

func (s *UpdatePolicyRequest) AddUserAgent(pagent ...string) IUpdatePolicyRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

// get policy by id request
func NewGetPolicyByIdRequest(lbID, lisID, policyID string) IGetPolicyByIdRequest {
	return &GetPolicyByIdRequest{
		LoadBalancerCommon: lscommon.LoadBalancerCommon{LoadBalancerId: lbID},
		ListenerCommon:     lscommon.ListenerCommon{ListenerId: lisID},
		PolicyCommon:       lscommon.PolicyCommon{PolicyId: policyID},
	}
}

func (s *GetPolicyByIdRequest) AddUserAgent(pagent ...string) IGetPolicyByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

var _ IGetPolicyByIdRequest = &GetPolicyByIdRequest{}

type GetPolicyByIdRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
	lscommon.ListenerCommon
	lscommon.PolicyCommon
}

// delete policy by id request
func NewDeletePolicyByIdRequest(lbID, lisID, policyID string) IDeletePolicyByIdRequest {
	return &DeletePolicyByIdRequest{
		LoadBalancerCommon: lscommon.LoadBalancerCommon{LoadBalancerId: lbID},
		ListenerCommon:     lscommon.ListenerCommon{ListenerId: lisID},
		PolicyCommon:       lscommon.PolicyCommon{PolicyId: policyID},
	}
}

func (s *DeletePolicyByIdRequest) AddUserAgent(pagent ...string) IDeletePolicyByIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

var _ IDeletePolicyByIdRequest = &DeletePolicyByIdRequest{}

type DeletePolicyByIdRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
	lscommon.ListenerCommon
	lscommon.PolicyCommon
}

type policyPositionRequest struct {
	Position int    `json:"position"`
	PolicyId string `json:"policyId"`
}

func NewReorderPoliciesRequest(lbID, lisID string) IReorderPoliciesRequest {
	return &ReorderPoliciesRequest{
		LoadBalancerCommon: lscommon.LoadBalancerCommon{LoadBalancerId: lbID},
		ListenerCommon:     lscommon.ListenerCommon{ListenerId: lisID},

		policyPositions: make([]policyPositionRequest, 0),
	}
}

var _ IReorderPoliciesRequest = &ReorderPoliciesRequest{}

type ReorderPoliciesRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
	lscommon.ListenerCommon

	policyPositions []policyPositionRequest
}

func (s *ReorderPoliciesRequest) AddUserAgent(pagent ...string) IReorderPoliciesRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *ReorderPoliciesRequest) WithPoliciesOrder(policies []string) IReorderPoliciesRequest {
	s.policyPositions = make([]policyPositionRequest, len(policies))
	for i, policy := range policies {
		s.policyPositions[i] = policyPositionRequest{
			Position: i + 1,
			PolicyId: policy,
		}
	}
	return s
}

func (s *ReorderPoliciesRequest) ToRequestBody() interface{} {
	return map[string]interface{}{
		"policies": s.policyPositions,
	}
}

// --------------------------------------------------------

// list policies request
func NewListPoliciesRequest(lbID, lisID string) IListPoliciesRequest {
	return &ListPoliciesRequest{
		LoadBalancerCommon: lscommon.LoadBalancerCommon{LoadBalancerId: lbID},
		ListenerCommon:     lscommon.ListenerCommon{ListenerId: lisID},
	}
}

func (s *ListPoliciesRequest) AddUserAgent(pagent ...string) IListPoliciesRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

var _ IListPoliciesRequest = &ListPoliciesRequest{}

type ListPoliciesRequest struct {
	lscommon.UserAgent
	lscommon.LoadBalancerCommon
	lscommon.ListenerCommon
}
