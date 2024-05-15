package client

import (
	"testing"

	v2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/identity/v2"
)

func TestNewClient(t *testing.T) {
	sdkConfig := NewSdkConfigure().
		WithClientId("c3f65a29").
		WithClientSecret("8637f6bf82").
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api")

	vngcloud := NewClient().WithRetryCount(3).WithSleep(10).Configure(sdkConfig)
	opt := v2.NewGetAccessTokenRequest("c3f65a5fc29", "8637fea6bf82")
	token, err := vngcloud.IamGateway().V2().IdentityService().GetAccessToken(opt)
	t.Log(token, err)
}
