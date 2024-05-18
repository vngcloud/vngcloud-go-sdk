package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

// Secgroup

type IGetSecgroupResponse interface {
	ToEntitySecgroup() *lsentity.Secgroup
}

type ICreateSecgroupResponse interface {
	ToEntitySecgroup() *lsentity.Secgroup
}

// Secgroup Rule

type ICreateSecgroupRuleResponse interface {
	ToEntitySecgroupRule() *lsentity.SecgroupRule
}

type IListSecgroupRulesBySecgroupIdResponse interface {
	ToEntityListSecgroupRules() *lsentity.ListSecgroupRules
}

type IGetNetworkByIdResponse interface {
	ToEntityNetwork() *lsentity.Network
}
