package policy

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
	lbCm "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/loadbalancer/v2"
)

type ListOptsBuilder struct {
	common.CommonOpts
	lbCm.LoadBalancerV2Common
	lbCm.ListenerV2Common
}

type (
	CreateOptsActionOpt      string
	CreateOptsCompareTypeOpt string
	CreateOptsRuleTypeOpt    string
)

const (
	CreateOptsActionOptREJECT         CreateOptsActionOpt = "REJECT"
	CreateOptsActionOptREDIRECTTOURL  CreateOptsActionOpt = "REDIRECT_TO_URL"
	CreateOptsActionOptREDIRECTTOPOOL CreateOptsActionOpt = "REDIRECT_TO_POOL"

	CreateOptsCompareTypeOptCONTAINS CreateOptsCompareTypeOpt = "CONTAINS"
	CreateOptsCompareTypeOptEQUALS   CreateOptsCompareTypeOpt = "EQUAL_TO"

	CreateOptsRuleTypeOptHOSTNAME CreateOptsRuleTypeOpt = "HOST_NAME"
	CreateOptsRuleTypeOptPATH     CreateOptsRuleTypeOpt = "PATH"
)

type CreateOptsBuilder struct {
	Name   string              `json:"name"`
	Action CreateOptsActionOpt `json:"action"`
	Rules  []Rule              `json:"rules"`

	RedirectPoolID string `json:"redirectPoolId"`

	RedirectURL      string `json:"redirectUrl"`
	RedirectHTTPCode int    `json:"redirectHttpCode"`
	KeepQueryString  bool   `json:"keepQueryString"`

	common.CommonOpts
	lbCm.LoadBalancerV2Common
	lbCm.ListenerV2Common
}
type Rule struct {
	CompareType CreateOptsCompareTypeOpt `json:"compareType"`
	RuleType    CreateOptsRuleTypeOpt    `json:"ruleType"`
	RuleValue   string                   `json:"ruleValue"`
}

func (s *CreateOptsBuilder) ToRequestBody() interface{} {
	switch s.Action {
	case CreateOptsActionOptREJECT:
		return struct {
			Name   string `json:"name"`
			Action string `json:"action"`
			Rules  []Rule `json:"rules"`
		}{
			Name:   s.Name,
			Action: string(s.Action),
			Rules:  s.Rules,
		}
	case CreateOptsActionOptREDIRECTTOURL:
		return struct {
			Name             string `json:"name"`
			Action           string `json:"action"`
			Rules            []Rule `json:"rules"`
			RedirectURL      string `json:"redirectUrl"`
			RedirectHTTPCode int    `json:"redirectHttpCode"`
			KeepQueryString  bool   `json:"keepQueryString"`
		}{
			Name:             s.Name,
			Action:           string(s.Action),
			Rules:            s.Rules,
			RedirectURL:      s.RedirectURL,
			RedirectHTTPCode: s.RedirectHTTPCode,
			KeepQueryString:  s.KeepQueryString,
		}
	case CreateOptsActionOptREDIRECTTOPOOL:
		return struct {
			Name           string `json:"name"`
			Action         string `json:"action"`
			Rules          []Rule `json:"rules"`
			RedirectPoolID string `json:"redirectPoolId"`
		}{
			Name:           s.Name,
			Action:         string(s.Action),
			Rules:          s.Rules,
			RedirectPoolID: s.RedirectPoolID,
		}
	}
	return nil
}

type DeleteOptsBuilder struct {
	common.CommonOpts
	lbCm.LoadBalancerV2Common
	lbCm.ListenerV2Common
	lbCm.PolicyV2Common
}
