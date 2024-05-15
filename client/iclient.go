package client

import (
	lctx "context"
	ltime "time"

	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsgateway "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/gateway"
)

type IClient interface {
	// List of builder methods
	WithHttpClient(phttpClient lsclient.IHttpClient) IClient
	WithContext(pctx lctx.Context) IClient
	WithAuthOption(pauthOption authOpts, pauthConfig ISdkConfigure) IClient
	WithRetryCount(pretry int) IClient
	WithDelay(pdelay ltime.Duration) IClient
	WithSleep(psleep ltime.Duration) IClient

	// List of functional methods
	Configure(psdkCfg ISdkConfigure) IClient

	// List of gateways
	IamGateway() lsgateway.IIamGateway
	VServerGateway() lsgateway.IVServerGateway
	VLBGateway() lsgateway.IVLBGateway
	VBackUpGateway() lsgateway.VBackUpGateway
}
