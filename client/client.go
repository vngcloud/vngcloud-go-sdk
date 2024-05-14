package client

import (
	lctx "context"
	lsync "sync"

	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

var (
	_ IClient = new(client)
)

type (
	client struct {
		Context     lctx.Context
		AccessToken string
		httpClient  IHttpClient

		mut       *lsync.RWMutex
		reauthMut *reauthLock
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
