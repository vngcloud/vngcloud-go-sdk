package subnet

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

type IGetResponse interface {
	ToSubnetObject() *objects.Subnet
}
