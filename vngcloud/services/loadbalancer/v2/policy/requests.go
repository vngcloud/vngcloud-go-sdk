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
	PolicyOptsActionOpt      string
	PolicyOptsCompareTypeOpt string
	PolicyOptsRuleTypeOpt    string
)

const (
	PolicyOptsActionOptREJECT         PolicyOptsActionOpt = "REJECT"
	PolicyOptsActionOptREDIRECTTOURL  PolicyOptsActionOpt = "REDIRECT_TO_URL"
	PolicyOptsActionOptREDIRECTTOPOOL PolicyOptsActionOpt = "REDIRECT_TO_POOL"

	PolicyOptsCompareTypeOptCONTAINS PolicyOptsCompareTypeOpt = "CONTAINS"
	PolicyOptsCompareTypeOptEQUALS   PolicyOptsCompareTypeOpt = "EQUAL_TO"

	PolicyOptsRuleTypeOptHOSTNAME PolicyOptsRuleTypeOpt = "HOST_NAME"
	PolicyOptsRuleTypeOptPATH     PolicyOptsRuleTypeOpt = "PATH"
)

type CreateOptsBuilder struct {
	Name   string              `json:"name"`
	Action PolicyOptsActionOpt `json:"action"`
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
	CompareType PolicyOptsCompareTypeOpt `json:"compareType"`
	RuleType    PolicyOptsRuleTypeOpt    `json:"ruleType"`
	RuleValue   string                   `json:"ruleValue"`
}

func (s *CreateOptsBuilder) ToRequestBody() interface{} {
	switch s.Action {
	case PolicyOptsActionOptREJECT:
		return struct {
			Name   string `json:"name"`
			Action string `json:"action"`
			Rules  []Rule `json:"rules"`
		}{
			Name:   s.Name,
			Action: string(s.Action),
			Rules:  s.Rules,
		}
	case PolicyOptsActionOptREDIRECTTOURL:
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
	case PolicyOptsActionOptREDIRECTTOPOOL:
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

type UpdateOptsBuilder struct {
	Action PolicyOptsActionOpt `json:"action"`
	Rules  []Rule              `json:"rules"`

	RedirectPoolID string `json:"redirectPoolId"`

	RedirectURL      string `json:"redirectUrl"`
	RedirectHTTPCode int    `json:"redirectHttpCode"`
	KeepQueryString  bool   `json:"keepQueryString"`

	common.CommonOpts
	lbCm.LoadBalancerV2Common
	lbCm.ListenerV2Common
	lbCm.PolicyV2Common
}

func (s *UpdateOptsBuilder) ToRequestBody() interface{} {
	switch s.Action {
	case PolicyOptsActionOptREJECT:
		return struct {
			Action string `json:"action"`
			Rules  []Rule `json:"rules"`
		}{
			Action: string(s.Action),
			Rules:  s.Rules,
		}
	case PolicyOptsActionOptREDIRECTTOURL:
		return struct {
			Action           string `json:"action"`
			Rules            []Rule `json:"rules"`
			RedirectURL      string `json:"redirectUrl"`
			RedirectHTTPCode int    `json:"redirectHttpCode"`
			KeepQueryString  bool   `json:"keepQueryString"`
		}{
			Action:           string(s.Action),
			Rules:            s.Rules,
			RedirectURL:      s.RedirectURL,
			RedirectHTTPCode: s.RedirectHTTPCode,
			KeepQueryString:  s.KeepQueryString,
		}
	case PolicyOptsActionOptREDIRECTTOPOOL:
		return struct {
			Action         string `json:"action"`
			Rules          []Rule `json:"rules"`
			RedirectPoolID string `json:"redirectPoolId"`
		}{
			Action:         string(s.Action),
			Rules:          s.Rules,
			RedirectPoolID: s.RedirectPoolID,
		}
	}
	return nil
}

type GetOptsBuilder struct {
	common.CommonOpts
	lbCm.LoadBalancerV2Common
	lbCm.ListenerV2Common
	lbCm.PolicyV2Common
}