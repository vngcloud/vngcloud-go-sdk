package test

import (
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lsportalV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/portal/v1"
	ltesting "testing"
)

func TestGetPortalInfoFailed(t *ltesting.T) {
	backendProjectId := getValueOfEnv("BACKEND_PROJECT_ID")
	vngcloud := invalidSdkConfig()
	opt := lsportalV1.NewGetPortalInfoRequest(backendProjectId)
	portal, err := vngcloud.VServerGateway().V1().PortalService().GetPortalInfo(opt)

	if err == nil {
		t.Errorf("Expect error but got nil")
	}

	if portal != nil {
		t.Errorf("Expect portal to be nil but got %+v", portal)
	}

	if !err.IsError(lserr.EcAuthenticationFailed) {
		t.Errorf("Expect error code to be %s but got %s", lserr.EcAuthenticationFailed, err.GetErrorCode())
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestGetPortalInfoSuccess(t *ltesting.T) {
	backendProjectId := getValueOfEnv("BACKEND_PROJECT_ID")
	vngcloud := validSdkConfig()
	opt := lsportalV1.NewGetPortalInfoRequest(backendProjectId)
	portal, err := vngcloud.VServerGateway().V1().PortalService().GetPortalInfo(opt)

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if portal == nil {
		t.Errorf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", portal)
	t.Log("PASS")
}

func TestGetPortalInfoFailure(t *ltesting.T) {
	backendProjectId := getValueOfEnv("FAKE_BACKEND_PROJECT_ID")
	vngcloud := validSdkConfig()
	opt := lsportalV1.NewGetPortalInfoRequest(backendProjectId)
	portal, err := vngcloud.VServerGateway().V1().PortalService().GetPortalInfo(opt)

	if err == nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if portal != nil {
		t.Errorf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}
