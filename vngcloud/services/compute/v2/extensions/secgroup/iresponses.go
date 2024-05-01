package secgroup

import "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"

type IUpdateSecgroupResponse interface {
	ToServerObject() *objects.Server
}
