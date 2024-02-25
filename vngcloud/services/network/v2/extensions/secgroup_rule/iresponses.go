package secgroup_rule

import "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"

type ICreateResponse interface {
	ToSecgroupRuleObject() *objects.SecgroupRule
}

type IListRulesBySecgroupIDResponse interface {
	ToListSecgroupRules() []*objects.SecgroupRule
}
