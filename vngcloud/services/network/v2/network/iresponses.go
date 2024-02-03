package network

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

type IGetResponse interface {
	ToNetworkObject() *objects.Network
}
