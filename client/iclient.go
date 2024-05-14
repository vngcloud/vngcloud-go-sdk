package client

import (
	lctx "context"
	ltime "time"

	lsgateway "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/gateway"
	lshttp "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/http"
)

type IClient interface {
	WithHttpClient(phttpClient lshttp.IHttpClient) IClient
	WithContext(pctx lctx.Context) IClient
	WithAuthOption(pauthOption authOpts, pauthConfig ISdkConfigure) IClient
	WithRetryCount(pretry int) IClient
	WithDelay(pdelay ltime.Duration) IClient
	WithSleep(psleep ltime.Duration) IClient

	IamGateway() lsgateway.IIamGateway
	VServerGateway() lsgateway.IVServerGateway
	VLBGateway() lsgateway.IVLBGateway
}
