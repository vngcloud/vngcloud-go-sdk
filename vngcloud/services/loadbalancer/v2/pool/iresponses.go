package pool

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

type ICreateResponse interface {
	ToPoolObject() *objects.Pool
}

type IListPoolsBasedLoadBalancerResponse interface {
	ToListPoolObjects() []*objects.Pool
}
