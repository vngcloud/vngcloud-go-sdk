package test

import (
	lctx "context"
	ltesting "testing"
	ltime "time"

	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/client"
	lsportalV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/portal/v1"
)

// ca984cf19c0c48a0838ba9f06d259ffa

func TestGetPortalInfo(t *ltesting.T) {
	sdkConfig := lsclient.NewSdkConfigure().
		WithClientId("___").
		WithClientSecret("___").
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway")

	vngcloud := lsclient.NewClient(lctx.TODO()).Configure(sdkConfig)
	opt := lsportalV1.NewGetPortalInfoRequest("ca984cf19c0c48a0838ba9f06d259ffa")
	portal, err := vngcloud.VServerGateway().V1().PortalService().GetPortalInfo(opt)
	t.Log(portal, err)

	ltime.Sleep(10 * ltime.Second)
}
