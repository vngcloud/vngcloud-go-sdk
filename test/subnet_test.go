package test

import (
	lsnetworkSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network/v2"
	ltesting "testing"
)

func TestGetSubnetByIdSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsnetworkSvcV2.NewGetSubnetByIdRequest("net-4f35f173-e0fe-4202-9c2b-5121b558bcd3", "sub-27a0562d-07f9-4e87-81fd-e0ba9658f156")
	network, err := vngcloud.VServerGateway().V2().NetworkService().GetSubnetById(opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	if network == nil {
		t.Fatalf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", network)
	t.Log("PASS")
}
