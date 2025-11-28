package test

import (
	ltesting "testing"

	"github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/glb/v1"
)

func TestGetGlobalListenerSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewGetGlobalListenerRequest("glb-a9799830-f7ef-40a8-ad05-ba7f81a8bb8d", "glis-6e0b0a21-0b52-40b7-90e5-f59a83e577c2")
	listener, sdkerr := vngcloud.GLBGateway().V1().GLBService().GetGlobalListener(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listener == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", listener)
	t.Log("PASS")
}
