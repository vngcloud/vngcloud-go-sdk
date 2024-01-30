package secgroup

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

type ICreateResponse interface {
	ToSecgroupObject() *objects.Secgroup
}

type IGetResponse interface {
	ToSecgroupObject() *objects.Secgroup
}

type IListResponse interface {
	ToListSecgroupObjects() []*objects.Secgroup
}
