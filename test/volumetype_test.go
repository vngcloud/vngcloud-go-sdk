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
	opt := lsvolume.NewGetVolumeTypeByIdRequest("vtype-61c3fc5b-f4e9-45b4-8957-8aa7b6029018")
	volume, sdkerr := vngcloud.VServerGateway().V1().VolumeService().GetVolumeTypeById(opt)

	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	if volume == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", volume)
	t.Log("PASS")
}

func TestGetDefaultVolumeType(t *ltesting.T) {
	vngcloud := validSdkConfig()
	volType, sdkerr := vngcloud.VServerGateway().V1().VolumeService().GetDefaultVolumeType()

	t.Log("Result: ", volType)
	t.Log("Error: ", sdkerr)
	t.Log("PASS")
}
