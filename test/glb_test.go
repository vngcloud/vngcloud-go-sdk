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

func TestGetGlobalPoolMemberSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewGetGlobalPoolMemberRequest("glb-a9799830-f7ef-40a8-ad05-ba7f81a8bb8d", "gpool-e5de4670-27e6-45cf-bc68-ec3803ed6849", "gpool-mem-4b3a819d-a83f-4964-8336-da6cb8edf529")
	poolMember, sdkerr := vngcloud.GLBGateway().V1().GLBService().GetGlobalPoolMember(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if poolMember == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Pool Member: %+v", poolMember)
	if poolMember.Members != nil && len(poolMember.Members.Items) > 0 {
		t.Logf("Members count: %d", len(poolMember.Members.Items))
		for i, member := range poolMember.Members.Items {
			t.Logf("Member[%d]: %+v", i, member)
		}
	} else {
		t.Log("No members found")
	}
	t.Log("PASS")
}

func TestDeleteGlobalPoolMemberSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewDeleteGlobalPoolMemberRequest(
		"glb-a9799830-f7ef-40a8-ad05-ba7f81a8bb8d",
		"gpool-e5de4670-27e6-45cf-bc68-ec3803ed6849",
		"gpool-mem-4b3a819d-a83f-4964-8336-da6cb8edf529",
	)

	sdkerr := vngcloud.GLBGateway().V1().GLBService().DeleteGlobalPoolMember(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("PASS")
}
