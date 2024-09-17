package test

import (
	lsnetworkSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network/v2"
	ltesting "testing"
)

func TestGetSecgroupByIdSuccess(t *ltesting.T) {
	secgroupId := "secg-d803abe8-2cf9-4b46-b3e7-94d8ab3c94ca"
	vngcloud := validSdkConfig()
	opt := lsnetworkSvcV2.NewGetSecgroupByIdRequest(secgroupId)
	secgroup, err := vngcloud.VServerGateway().V2().NetworkService().GetSecgroupById(opt)

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if secgroup == nil {
		t.Errorf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", secgroup)
	t.Log("PASS")
}

func TestGetSecgroupByIdFailure(t *ltesting.T) {
	secgroupId := "secg-90d617b4-b893-407b-a9a8-3bd80c177920"
	vngcloud := validSdkConfig()
	opt := lsnetworkSvcV2.NewGetSecgroupByIdRequest(secgroupId)
	secgroup, err := vngcloud.VServerGateway().V2().NetworkService().GetSecgroupById(opt)

	if err == nil {
		t.Errorf("Expect error not to be nil but got nil")
	}

	if secgroup != nil {
		t.Errorf("Expect portal to be nil but got %+v", secgroup)
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestCreateSecgroupSameNameFailure(t *ltesting.T) {
	secgroupName := "cuongdm3-temporal"
	vngcloud := validSdkConfig()
	opt := lsnetworkSvcV2.NewCreateSecgroupRequest(secgroupName, "this is a test")
	secgroup, err := vngcloud.VServerGateway().V2().NetworkService().CreateSecgroup(opt)

	if err == nil {
		t.Errorf("Expect error not to be nil but got nil")
	}

	if secgroup != nil {
		t.Errorf("Expect portal to be nil but got %+v", secgroup)
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestCreateSecgroupSuccess(t *ltesting.T) {
	secgroupName := "cuongdm3-temporal-1"
	vngcloud := validSdkConfig()
	opt := lsnetworkSvcV2.NewCreateSecgroupRequest(secgroupName, "this is a test")
	secgroup, err := vngcloud.VServerGateway().V2().NetworkService().CreateSecgroup(opt)

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if secgroup == nil {
		t.Errorf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", secgroup)
	t.Log("PASS")
}

func TestDeleteSecgroupByIdFailure(t *ltesting.T) {
	secgroupId := "secg-90d617b4-b893-407b-a9a8-3bd80c177920"
	vngcloud := validSdkConfig()
	opt := lsnetworkSvcV2.NewDeleteSecgroupByIdRequest(secgroupId)
	err := vngcloud.VServerGateway().V2().NetworkService().DeleteSecgroupById(opt)

	if err == nil {
		t.Errorf("Expect error not to be nil but got nil")
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestDeleteSecgroupByIdSuccess(t *ltesting.T) {
	secgroupId := "secg-3787f73d-d62b-49ca-96cd-226b7dc8ead4"
	vngcloud := validSdkConfig()
	opt := lsnetworkSvcV2.NewDeleteSecgroupByIdRequest(secgroupId)
	err := vngcloud.VServerGateway().V2().NetworkService().DeleteSecgroupById(opt)

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestListAllServerBySecgroupIdSuccess(t *ltesting.T) {
	secgroupId := "secg-1395e86c-9631-4c13-xxxx-e41be5bdaab3"
	vngcloud := validSdkConfig()
	opt := lsnetworkSvcV2.NewListAllServersBySecgroupIdRequest(secgroupId)
	serbvers, err := vngcloud.VServerGateway().V2().NetworkService().ListAllServersBySecgroupId(opt)

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if serbvers == nil {
		t.Errorf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", serbvers)
	t.Log("PASS")
}
