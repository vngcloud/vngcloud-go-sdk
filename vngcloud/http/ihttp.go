package http

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/client"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	ltime "time"
)

type IHttpClient interface {
	WithRetryCount(pretryCount int) IHttpClient
	WithDelay(pdelay ltime.Duration) IHttpClient
	WithSleep(psleep ltime.Duration) IHttpClient
	WithReauthFunc(preauthFunc func() (lsclient.ISdkAuthentication, lserr.ISdkError)) IHttpClient
}
