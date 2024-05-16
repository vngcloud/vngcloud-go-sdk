package client

import (
	lctx "context"
	v1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/portal/v1"
	"testing"
	ltime "time"

	v2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/identity/v2"
)

// ca984cf19c0c48a0838ba9f06d259ffa

func TestNewClient(t *testing.T) {
	sdkConfig := NewSdkConfigure().
		WithClientId("___").
		WithClientSecret("___").
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway")

	vngcloud := NewClient(lctx.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
	opt := v2.NewGetAccessTokenRequest("___", "___")
	token, err := vngcloud.IamGateway().V2().IdentityService().GetAccessToken(opt)
	t.Log(token, err)
}

func TestUnixNano(t *testing.T) {
	now := ltime.Now()

	t.Log(now.UnixNano())

	// convert unix nano to time
	me := ltime.Unix(0, now.UnixNano())
	ltime.Sleep(5 * ltime.Second)
	t.Log(me.UnixNano())

	now = ltime.Now()
	t.Log(now.UnixNano())

	t.Log(now.Sub(me) >= 5*ltime.Second)
}

func TestGetPortalInfo(t *testing.T) {
	sdkConfig := NewSdkConfigure().
		WithClientId("___").
		WithClientSecret("___").
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway")

	vngcloud := NewClient(lctx.TODO()).Configure(sdkConfig)
	opt := v1.NewGetPortalInfoRequest("ca984cf19c0c48a0838ba9f06d259ffa")
	portal, err := vngcloud.VServerGateway().V1().PortalService().GetPortalInfo(opt)
	t.Log(portal, err)

	ltime.Sleep(10 * ltime.Second)
}
