package client

import (
	lctx "context"
	ltime "time"

	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsgateway "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/gateway"
)

type IClient interface {
	WithHttpClient(phttpClient lsclient.IHttpClient) IClient
	WithContext(pctx lctx.Context) IClient
	WithAuthOption(pauthOption authOpts, pauthConfig ISdkConfigure) IClient
	WithRetryCount(pretry int) IClient
	WithDelay(pdelay ltime.Duration) IClient
	WithSleep(psleep ltime.Duration) IClient

	Configure(psdkCfg ISdkConfigure) IClient

	IamGateway() lsgateway.IIamGateway
	VServerGateway() lsgateway.IVServerGateway
	VLBGateway() lsgateway.IVLBGateway
}
