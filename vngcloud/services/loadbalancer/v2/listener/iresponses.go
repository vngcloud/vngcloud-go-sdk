package listener

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

type ICreateResponse interface {
	ToListenerObject() *objects.Listener
}

type IGetBasedLoadBalancerResponse interface {
	ToListListenerObject() []*objects.Listener
}
