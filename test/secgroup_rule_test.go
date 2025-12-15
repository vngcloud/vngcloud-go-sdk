package test

import (
	ltesting "testing"

	lsnetworkSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network/v2"
)

func TestCreateSecgroupRuleSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsnetworkSvcV2.NewCreateSecgroupRuleRequest(
		lsnetworkSvcV2.SecgroupRuleDirectionIngress,
		lsnetworkSvcV2.SecgroupRuleEtherTypeIPv4,
		lsnetworkSvcV2.SecgroupRuleProtocolTCP,
		6444, 6450,
		"0.0.0.0",
		"secg-f5a2de4d-f8a2-4269-bcba-7c50f09ee838",
		"test2",
	)
	secgroupRule, err := vngcloud.VServerGateway().V2().NetworkService().CreateSecgroupRule(opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	if secgroupRule == nil {
		t.Fatalf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", secgroupRule)
	t.Log("PASS")
}

func TestCreateSecgroupRuleFailure(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsnetworkSvcV2.NewCreateSecgroupRuleRequest(
		lsnetworkSvcV2.SecgroupRuleDirectionIngress,
		lsnetworkSvcV2.SecgroupRuleEtherTypeIPv4,
		lsnetworkSvcV2.SecgroupRuleProtocolTCP,
		6443, 6443,
		"0.0.0.0",
		"secg-f5a2de4d-f8a2-4269-bcba-7c50f09ee840",
		"test",
	)
	secgroupRule, err := vngcloud.VServerGateway().V2().NetworkService().CreateSecgroupRule(opt)

	if err == nil {
		t.Errorf("Expect error not to be nil but got nil")
	}

	if secgroupRule != nil {
		t.Errorf("Expect portal to be nil but got %+v", secgroupRule)
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestDeleteSecgroupRuleFailure(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsnetworkSvcV2.NewDeleteSecgroupRuleByIdRequest("secr-8c44dc7f-3916-4952-8a01-884ad9360f00")
	err := vngcloud.VServerGateway().V2().NetworkService().DeleteSecgroupRuleById(opt)

	if err == nil {
		t.Errorf("Expect error not to be nil but got nil")
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestDeleteSecgroupRuleSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsnetworkSvcV2.NewDeleteSecgroupRuleByIdRequest("secr-d8c83235-93ae-4de1-85c2-365a32494121")
	err := vngcloud.VServerGateway().V2().NetworkService().DeleteSecgroupRuleById(opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	t.Log("RESULT:", err)
	t.Log("PASS")
}

func TestListSecgroupRulesBySecgroupIdFailure(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsnetworkSvcV2.NewListSecgroupRulesBySecgroupIdRequest("secg-f5a2de4d-f8a2-4269-bcba-7c50f09ee840")
	rules, err := vngcloud.VServerGateway().V2().NetworkService().ListSecgroupRulesBySecgroupId(opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	if rules == nil {
		t.Fatalf("Expect rules not to be nil but got nil")
	}

	t.Log("RESULT:", rules)
	t.Log("PASS")
}

func TestListSecgroupRulesBySecgroupIdSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsnetworkSvcV2.NewListSecgroupRulesBySecgroupIdRequest("secg-f5a2de4d-f8a2-4269-bcba-7c50f09ee838")
	rules, err := vngcloud.VServerGateway().V2().NetworkService().ListSecgroupRulesBySecgroupId(opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	if rules == nil {
		t.Fatalf("Expect rules not to be nil but got nil")
	}

	t.Log("RESULT:", rules)
	t.Log("PASS")
}
