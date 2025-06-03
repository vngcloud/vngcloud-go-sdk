package test

import (
	lctx "context"
	"fmt"
	ltesting "testing"

	lgodotenv "github.com/joho/godotenv"

	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/client"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lsidentityV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/identity/v2"
)

func getEnv() (string, string) {
	envFile, _ := lgodotenv.Read("./env.yaml")
	clientId := envFile["VNGCLOUD_CLIENT_ID"]
	clientSecret := envFile["VNGCLOUD_CLIENT_SECRET"]

	return clientId, clientSecret
}

func getEnvCuongDm4() (string, string) {
	envFile, _ := lgodotenv.Read("/Users/cuongdm8499/Me/git-cuongpiger/secret/work/vngcloud/iam/env")
	clientId := envFile["CUONGDM4_CLIENT_ID"]
	clientSecret := envFile["CUONGDM4_CLIENT_SECRET"]

	return clientId, clientSecret
}

func getEnvDevOps() (string, string) {
	envFile, _ := lgodotenv.Read("./env.yaml")
	clientId := envFile["CLIENT_ID_DEVOPS"]
	clientSecret := envFile["CLIENT_SECRET_DEVOPS"]

	return clientId, clientSecret
}

func getValueOfEnv(pkey string) string {
	envFile, _ := lgodotenv.Read("./env.yaml")
	value := envFile[pkey]
	return value
}

func validSdkConfig() lsclient.IClient {
	clientId, clientSecret := getEnv()
	sdkConfig := lsclient.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithUserId(getValueOfEnv("VNGCLOUD_USER_ID")).
		WithZoneId(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectId(getValueOfEnv("VNGCLOUD_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork").
		WithGLBEndpoint("https://glb.console.vngcloud.vn/glb-controller/")

	return lsclient.NewClient(lctx.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validUserSdkConfig() lsclient.IClient {
	clientId, clientSecret := getValueOfEnv("USER_CLIENT_ID"), getValueOfEnv("USER_CLIENT_SECRET")
	sdkConfig := lsclient.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithUserId(getValueOfEnv("VNGCLOUD_USER_ID")).
		WithZoneId(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectId(getValueOfEnv("USER_PROJECT")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork")

	return lsclient.NewClient(lctx.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validUserSdkConfigForCuongDm4() lsclient.IClient {
	clientId, clientSecret := getEnvCuongDm4()
	sdkConfig := lsclient.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithProjectId("pro-462803f3-6858-466f-bf05-df2b33faa360").
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork")

	return lsclient.NewClient(lctx.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSdkConfigHanRegion() lsclient.IClient {
	clientId, clientSecret := getEnv()
	sdkConfig := lsclient.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithUserId(getValueOfEnv("VNGCLOUD_USER_ID")).
		WithZoneId(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectId(getValueOfEnv("HAN01_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://han-1.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://han-1.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork").
		WithUserAgent("vngcloud-go-sdk")

	return lsclient.NewClient(lctx.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validHcm3bSdkConfig() lsclient.IClient {
	sdkConfig := lsclient.NewSdkConfigure().
		WithClientId(getValueOfEnv("HCM3B_CLIENT_ID")).
		WithClientSecret(getValueOfEnv("HCM3B_CLIENT_SECRET")).
		WithZoneId(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectId(getValueOfEnv("HCM3B_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork")

	return lsclient.NewClient(lctx.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validHcm3bSuperSdkConfig() lsclient.IClient {
	sdkConfig := lsclient.NewSdkConfigure().
		WithClientId(getValueOfEnv("HCM3B_SUPER_CLIENT_ID")).
		WithClientSecret(getValueOfEnv("HCM3B_SUPER_CLIENT_SECRET")).
		WithZoneId(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectId(getValueOfEnv("HCM3B_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork")

	return lsclient.NewClient(lctx.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validVinhNt8SdkConfig() lsclient.IClient {
	sdkConfig := lsclient.NewSdkConfigure().
		WithClientId(getValueOfEnv("VINHCLIENT_ID")).
		WithClientSecret(getValueOfEnv("VINHCLIENT_SECRET")).
		WithZoneId(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectId(getValueOfEnv("VINH_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork").
		WithVNetworkEndpoint("https://hcm-3.console.vngcloud.vn/vserver/vnetwork-gateway/vnetwork")

	return lsclient.NewClient(lctx.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func customerSdkConfig() lsclient.IClient {
	sdkConfig := lsclient.NewSdkConfigure().
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway")

	return lsclient.NewClient(lctx.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSdkHannibalConfig() lsclient.IClient {
	clientId, clientSecret := getValueOfEnv("HANNIBAL_CLIENT_ID"), getValueOfEnv("HANNIBAL_CLIENT_SECRET")
	sdkConfig := lsclient.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithProjectId(getValueOfEnv("HANNIBAL_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway")

	return lsclient.NewClient(lctx.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSuperSdkConfig() lsclient.IClient {
	clientId, clientSecret := getValueOfEnv("VNGCLOUD_SUPER_CLIENT_ID"), getValueOfEnv("VNGCLOUD_SUPER_CLIENT_SECRET")
	sdkConfig := lsclient.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithZoneId(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectId(getValueOfEnv("VNGCLOUD_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03-api.vngcloud.vn/vnetwork-gateway")

	return lsclient.NewClient(lctx.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSuperSdkHcm03bConfig() lsclient.IClient {
	clientId, clientSecret := getValueOfEnv("VNGCLOUD_PROD_HCM03B_CLIENT_ID"), getValueOfEnv("VNGCLOUD_PROD_HCM03B_CLIENT_SECRET")
	sdkConfig := lsclient.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithZoneId(getValueOfEnv("VNGCLOUD_ZONE_ID")).
		WithProjectId(getValueOfEnv("VNGCLOUD_PROD_HCM03B_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03-api.vngcloud.vn/vnetwork-gateway")

	return lsclient.NewClient(lctx.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validUser11412SdkConfig() lsclient.IClient {
	clientId, clientSecret := getValueOfEnv("VNGCLOUD_SUPER_CLIENT_ID"), getValueOfEnv("VNGCLOUD_SUPER_CLIENT_SECRET")
	sdkConfig := lsclient.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithProjectId(getValueOfEnv("USER_11412_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway")

	return lsclient.NewClient(lctx.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validUser11412() lsclient.IClient {
	clientId, clientSecret := getValueOfEnv("USER_11412_CLIENT_ID"), getValueOfEnv("USER_11412_CLIENT_SECRET")
	sdkConfig := lsclient.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithProjectId(getValueOfEnv("USER_11412_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
		WithVNetworkEndpoint("https://vnetwork-hcm03-api.vngcloud.vn/vnetwork-gateway")

	return lsclient.NewClient(lctx.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSuperSdkConfig2() lsclient.IClient {
	clientId, clientSecret := getValueOfEnv("VNGCLOUD_SUPER_CLIENT_ID"), getValueOfEnv("VNGCLOUD_SUPER_CLIENT_SECRET")
	sdkConfig := lsclient.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithProjectId(getValueOfEnv("VINH_PROJECT_ID")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
		WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway")

	return lsclient.NewClient(lctx.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func validSdkConfigDevops() lsclient.IClient {
	clientId, clientSecret := getEnvDevOps()
	sdkConfig := lsclient.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithProjectId(getValueOfEnv("PROJECT_ID_DEVOPS")).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway")

	return lsclient.NewClient(lctx.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
}

func invalidSdkConfig() lsclient.IClient {
	clientId := "___"
	clientSecret := "___"
	sdkConfig := lsclient.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
		WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway")

	return lsclient.NewClient(lctx.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
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
	vngcloud := validSdkConfig()
	opt := lsidentityV2.NewGetAccessTokenRequest(clientId, clientSecret)
	token, err := vngcloud.IamGateway().V2().IdentityService().GetAccessToken(opt)

	if err != nil || token == nil {
		t.Error("This testcase MUST pass")
	}

	fmt.Println("RESULT 2:", token.Token)
	t.Log("RESULT:", token)
	t.Log("PASS")
}

func TestASuperuthenPass(t *ltesting.T) {
	clientId, clientSecret := getValueOfEnv("VNGCLOUD_SUPER_CLIENT_ID"), getValueOfEnv("VNGCLOUD_SUPER_CLIENT_SECRET")
	vngcloud := validSdkConfig()
	opt := lsidentityV2.NewGetAccessTokenRequest(clientId, clientSecret)
	token, err := vngcloud.IamGateway().V2().IdentityService().GetAccessToken(opt)

	if err != nil || token == nil {
		t.Error("This testcase MUST pass")
	}

	fmt.Println("RESULT 2:", token.Token)
	t.Log("RESULT:", token.Token)
	t.Log("PASS")
}

func TestVinhAuthenPass(t *ltesting.T) {
	clientId, clientSecret := getValueOfEnv("VINHCLIENT_ID"), getValueOfEnv("VINHCLIENT_SECRET")
	vngcloud := validSuperSdkConfig2()
	opt := lsidentityV2.NewGetAccessTokenRequest(clientId, clientSecret)
	token, err := vngcloud.IamGateway().V2().IdentityService().GetAccessToken(opt)

	if err != nil || token == nil {
		t.Error("This testcase MUST pass")
	}

	t.Log("RESULT:", token)
	t.Log("PASS")
}
