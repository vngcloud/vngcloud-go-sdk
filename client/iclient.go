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
	WithAuthOption(pauthOption lsclient.AuthOpts, pauthConfig ISdkConfigure) IClient
	WithKvDefaultHeaders(pargs ...string) IClient
	WithRetryCount(pretry int) IClient
	WithSleep(psleep ltime.Duration) IClient
	WithProjectId(pprojectId string) IClient

	// List of functional methods
	Configure(psdkCfg ISdkConfigure) IClient
	GetUserAgent() string

	// List of gateways
	IamGateway() lsgateway.IIamGateway
	VServerGateway() lsgateway.IVServerGateway
	VLBGateway() lsgateway.IVLBGateway
	VBackUpGateway() lsgateway.IVBackUpGateway
	VNetworkGateway() lsgateway.IVNetworkGateway
	GLBGateway() lsgateway.IGLBGateway
}
