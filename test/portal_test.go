package test

import (
	ltesting "testing"

	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
	lsportalV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/portal/v1"
	lsportalV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/portal/v2"
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

func TestGetPortalInfoSuccess2(t *ltesting.T) {
	backendProjectId := getValueOfEnv("USER_11412_OS_PROJECT_ID")
	vngcloud := validUser11412SdkConfig()
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

func TestListAllQuotaSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	quotas, err := vngcloud.VServerGateway().V2().PortalService().ListAllQuotaUsed()

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if quotas == nil {
		t.Errorf("Expect quotas not to be nil but got nil")
	}

	t.Log("RESULT:", quotas)
	t.Log("PASS")
}

func TestGetQuotaByNameFailure(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsportalV2.NewGetQuotaByNameRequest("fake-quota-name")
	quota, err := vngcloud.VServerGateway().V2().PortalService().GetQuotaByName(opt)

	if err == nil {
		t.Errorf("Expect error but got nil")
	}

	if quota != nil {
		t.Errorf("Expect quota to be nil but got %+v", quota)
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestGetQuotaByNamePass(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsportalV2.NewGetQuotaByNameRequest(lsportalV2.QtVolumeAttachLimit)
	quota, err := vngcloud.VServerGateway().V2().PortalService().GetQuotaByName(opt)

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if quota == nil {
		t.Errorf("Expect quota not to be nil but got nil")
	}

	t.Log("RESULT:", quota)
	t.Log("PASS")
}

func TestListProjects(t *ltesting.T) {
	vngcloud := validSdkConfig()
	projects, err := vngcloud.VServerGateway().V1().PortalService().ListProjects(lsportalV1.NewListProjectsRequest())
	if err != nil {
		t.Log("Error: ", err)
	}

	t.Log("Result: ", projects.At(0))
}

func TestListPortalUser11412(t *ltesting.T) {
	vngcloud := validUser11412()
	projects, err := vngcloud.VServerGateway().V1().PortalService().ListProjects(lsportalV1.NewListProjectsRequest())
	if err != nil {
		t.Log("Error: ", err)
	}

	t.Log("Result: ", projects.At(0))
}

func TestListZones(t *ltesting.T) {
	vngcloud := validSdkConfig()
	zones, err := vngcloud.VServerGateway().V1().PortalService().ListZones()
	if err != nil {
		t.Log("Error: ", err)
	}

	t.Log("Result: ", zones.Items)
}
