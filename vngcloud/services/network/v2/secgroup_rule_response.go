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
