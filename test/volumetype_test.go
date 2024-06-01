package test

import (
	lsvolume "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/volume/v1"
	ltesting "testing"
)

func TestGetVolumeTypeFailure(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsvolume.NewGetVolumeTypeByIdRequest("fake-id")
	volume, sdkerr := vngcloud.VServerGateway().V1().VolumeService().GetVolumeTypeById(opt)
	if sdkerr == nil {
		t.Fatalf("Expect error but got nil")
	}

	if volume != nil {
		t.Fatalf("Expect nil but got %v", volume)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestGetVolumeTypeSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsvolume.NewGetVolumeTypeByIdRequest("vtype-7a7a8610-34f5-11ee-be56-0242ac120002")
	volume, sdkerr := vngcloud.VServerGateway().V1().VolumeService().GetVolumeTypeById(opt)

	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	if volume == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", volume)
	t.Log("PASS")
}
