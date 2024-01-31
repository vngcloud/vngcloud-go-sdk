package server

import "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"

type IGetResponse interface {
	ToServerObject() *objects.Server
}
