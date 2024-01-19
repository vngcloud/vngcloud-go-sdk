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

type IGetResponse interface {
	ToPoolObject() *objects.Pool
}

type IGetMemberResponse interface {
	ToListMemberObject() []*objects.Member
}
