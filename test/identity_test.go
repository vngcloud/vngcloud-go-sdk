package test

import (
	lctx "context"
	ltesting "testing"

	lgodotenv "github.com/joho/godotenv"

	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/client"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lsidentityV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/identity/v2"
)

func getEnv() (string, string) {
	envFile, _ := lgodotenv.Read("/mnt/vayne/git-vngcloud/vngcloud-go-sdk/secret/env")
	clientId := envFile["VNGCLOUD_CLIENT_ID"]
	clientSecret := envFile["VNGCLOUD_CLIENT_SECRET"]

	return clientId, clientSecret
}

func TestAuthenFailed(t *ltesting.T) {
	clientId := "cc136360-709c-4248-9358-e8e96c74480a"
	clientSecret := "___"

	sdkConfig := lsclient.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway")

	vngcloud := lsclient.NewClient(lctx.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
	opt := lsidentityV2.NewGetAccessTokenRequest(clientId, clientSecret)
	token, err := vngcloud.IamGateway().V2().IdentityService().GetAccessToken(opt)

	if err == nil {
		t.Error("Error MUST not be nil")
	}

	if token != nil {
		t.Error("Token MUST be nil")
	}

	if !err.IsError(lserr.EcAuthenticationFailed) {
		t.Error("Error MUST be VngCloudIamAuthenticationFailed")
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestAuthenPass(t *ltesting.T) {
	clientId, clientSecret := getEnv()

	sdkConfig := lsclient.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway")

	vngcloud := lsclient.NewClient(lctx.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
	opt := lsidentityV2.NewGetAccessTokenRequest(clientId, clientSecret)
	token, err := vngcloud.IamGateway().V2().IdentityService().GetAccessToken(opt)

	if err != nil || token == nil {
		t.Error("This testcase MUST pass")
	}

	t.Log("RESULT:", token)
	t.Log("PASS")
}
