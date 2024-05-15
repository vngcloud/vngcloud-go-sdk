package client

import (
	lctx "context"
	"testing"
	ltime "time"

	v2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/identity/v2"
)

func TestNewClient(t *testing.T) {
	sdkConfig := NewSdkConfigure().
		WithClientId("c3f65a29").
		WithClientSecret("8637f6bf82").
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api")

	vngcloud := NewClient(lctx.TODO()).WithRetryCount(3).WithSleep(10).Configure(sdkConfig)
	opt := v2.NewGetAccessTokenRequest("c3f65a5fc29", "8637fea6bf82")
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
