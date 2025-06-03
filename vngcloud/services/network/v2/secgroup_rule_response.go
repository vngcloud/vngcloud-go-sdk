package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type CreateSecgroupRuleResponse struct {
	Data struct {
		ID             int     `json:"id"`
		UUID           string  `json:"uuid"`
		CreatedAt      string  `json:"createdAt"`
		DeletedAt      *string `json:"deletedAt,omitempty"`
		UpdatedAt      *string `json:"updatedAt,omitempty"`
		Status         string  `json:"status"`
		SecgroupUuid   string  `json:"secgroupUuid"`
		Direction      string  `json:"direction"`
		EtherType      string  `json:"etherType"`
		PortRangeMax   int     `json:"portRangeMax"`
		PortRangeMin   int     `json:"portRangeMin"`
		Protocol       string  `json:"protocol"`
		RuleId         int     `json:"ruleId"`
		Description    string  `json:"description"`
		RemoteIpPrefix string  `json:"remoteIpPrefix"`
		IsSystem       bool    `json:"isSystem"`
	} `json:"data"`
}

func (s *CreateSecgroupRuleResponse) ToEntitySecgroupRule() *lsentity.SecgroupRule {
	return &lsentity.SecgroupRule{
		Id:             s.Data.UUID,
		SecgroupId:     s.Data.SecgroupUuid,
		Direction:      s.Data.Direction,
		EtherType:      s.Data.EtherType,
		PortRangeMax:   s.Data.PortRangeMax,
		PortRangeMin:   s.Data.PortRangeMin,
		Protocol:       s.Data.Protocol,
		Description:    s.Data.Description,
		RemoteIPPrefix: s.Data.RemoteIpPrefix,
	}
}

type ListSecgroupRulesBySecgroupIdResponse struct {
	Data []struct {
		ID             string `json:"id"`
		Direction      string `json:"direction"`
		EtherType      string `json:"etherType"`
		Protocol       string `json:"protocol"`
		PortRangeMin   int    `json:"portRangeMin"`
		PortRangeMax   int    `json:"portRangeMax"`
		RemoteIpPrefix string `json:"remoteIpPrefix"`
		RemoteGroupId  string `json:"remoteGroupId"`
		Status         string `json:"status"`
		Description    string `json:"description"`
		CreatedAt      string `json:"createdAt"`
	} `json:"data"`
}

func (s *ListSecgroupRulesBySecgroupIdResponse) ToEntityListSecgroupRules() *lsentity.ListSecgroupRules {
	lsr := &lsentity.ListSecgroupRules{
		Items: make([]*lsentity.SecgroupRule, 0),
	}

	for _, rule := range s.Data {
		lsr.Items = append(lsr.Items, &lsentity.SecgroupRule{
			Id:             rule.ID,
			Direction:      rule.Direction,
			EtherType:      rule.EtherType,
			Protocol:       rule.Protocol,
			PortRangeMin:   rule.PortRangeMin,
			PortRangeMax:   rule.PortRangeMax,
			RemoteIPPrefix: rule.RemoteIpPrefix,
			Description:    rule.Description,
		})
	}

	return lsr
}
