package client

import (
	lctx "context"
	lsync "sync"

	lsgateway "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/gateway"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

var (
	_ IClient = new(client)
)

type (
	client struct {
		context    lctx.Context
		httpClient IHttpClient

		mut       *lsync.RWMutex
		reauthMut *reauthLock

		identityGateway lsgateway.IIdentityGateway
		vserverGateway  lsgateway.IVServerGateway
		vlbGateway      lsgateway.IVLBGateway
	}

	reauthLock struct {
		ongoing *reauthFuture

		lsync.RWMutex
	}

	reauthFuture struct {
		done   chan struct{}
		sdkErr lserr.ISdkError
	}
)

func NewClient() IClient {
	c := new(client)
	return c
}

func (s *client) WithHttpClient(pclient IHttpClient) IClient {
	s.httpClient = pclient
	return s
}

func (s *client) WithContext(pctx lctx.Context) IClient {
	s.context = pctx
	return s
}
