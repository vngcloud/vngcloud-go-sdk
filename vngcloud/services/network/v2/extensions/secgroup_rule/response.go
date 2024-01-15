package secgroup_rule

import "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"

type CreateResponse struct {
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

func (s *CreateResponse) ToSecgroupRuleObject() *objects.SecgroupRule {
	if s == nil {
		return nil
	}

	return &objects.SecgroupRule{
		UUID:           s.Data.UUID,
		SecgroupUUID:   s.Data.SecgroupUuid,
		Direction:      s.Data.Direction,
		EtherType:      s.Data.EtherType,
		PortRangeMax:   s.Data.PortRangeMax,
		PortRangeMin:   s.Data.PortRangeMin,
		Protocol:       s.Data.Protocol,
		Description:    s.Data.Description,
		RemoteIPPrefix: s.Data.RemoteIpPrefix,
	}
}
