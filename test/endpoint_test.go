package test

import (
	lsnwv1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network/v1"
	ltesting "testing"
)

func TestGetEndpointSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsnwv1.NewGetEndpointByIdRequest("enp-8e54a777-1342-4926-b21b-d405eff16649")

	lb, sdkerr := vngcloud.VNetworkGateway().V1().NetworkService().GetEndpointById(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr.GetMessage())
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}
