package policy

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

// *********************************************** Response of List API ************************************************

type IListResponse interface {
	ToListPolicyObjects() []*objects.Policy
}

// ************************************ Common methods that often used in responses ************************************

type ICommonResponse interface {
	ToPolicyObject() *objects.Policy
}
