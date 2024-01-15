package loadbalancer

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

// *********************************************** Response of List API ************************************************

type IListResponse interface {
	ToListLoadBalancerObjects() []*objects.LoadBalancer
}

// ************************************ Common methods that often used in responses ************************************

type ICommonResponse interface {
	ToLoadBalancerObject() *objects.LoadBalancer
}
