package secgroup_rule

import "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"

type ICreateResponse interface {
	ToSecgroupRuleObject() *objects.SecgroupRule
}
