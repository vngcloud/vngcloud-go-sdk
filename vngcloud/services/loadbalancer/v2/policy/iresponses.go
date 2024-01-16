package policy

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

type IListResponse interface {
	ToListPolicyObjects() []*objects.Policy
}

type ICreateResponse interface {
	ToPolicyObject() *objects.Policy
}
