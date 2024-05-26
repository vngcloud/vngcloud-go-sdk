package test

import (
	v2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/volume/v2"
	ltesting "testing"
)

func TestCreateVolumeFailure(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewCreateBlockVolumeRequest(
		"testsadsadsada",
		"vtype-7a7a8610-34f5-11ee-be56-0242ac120002",
		10,
	)
	volume, sdkerr := vngcloud.VServerGateway().V2().VolumeService().CreateBlockVolume(opt)
	if sdkerr == nil {
		t.Fatalf("Expect error but got nil")
	}

	if volume != nil {
		t.Fatalf("Expect nil but got %v", volume)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestCreateVolumeSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewCreateBlockVolumeRequest(
		"cuongdm3-test-tags",
		"vtype-7a7a8610-34f5-11ee-be56-0242ac120002",
		20,
	).WithTags("cuongdm3", "deptrai", "wife", "unknown")
	volume, sdkerr := vngcloud.VServerGateway().V2().VolumeService().CreateBlockVolume(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	if volume == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", volume)
	t.Log("PASS")
}
