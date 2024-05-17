package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type IGetSecgroupResponse interface {
	ToEntitySecgroup() *lsentity.Secgroup
}

type ICreateSecgroupResponse interface {
	ToEntitySecgroup() *lsentity.Secgroup
}
