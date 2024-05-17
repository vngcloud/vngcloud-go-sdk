package test

import (
	lsnetworkSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network/v2"
	ltesting "testing"
)

func TestGetSecgroupByIdSuccess(t *ltesting.T) {
	secgroupId := "secg-90d617b4-b893-407b-a9a8-3bd80c177918"
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
