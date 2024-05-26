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

func TestDeleteVolumeByIdFailure(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewDeleteBlockVolumeByIdRequest("this-is-fake")
	sdkerr := vngcloud.VServerGateway().V2().VolumeService().DeleteBlockVolumeById(opt)
	if sdkerr == nil {
		t.Fatalf("Expect error but got nil")
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestDeleteVolumeByIdSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewDeleteBlockVolumeByIdRequest("vol-51f71146-9c20-4615-a73e-a43a39bf03ea")
	sdkerr := vngcloud.VServerGateway().V2().VolumeService().DeleteBlockVolumeById(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestListBlockVolumeSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewListBlockVolumesRequest(1, 10)
	volumes, sdkerr := vngcloud.VServerGateway().V2().VolumeService().ListBlockVolumes(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	if volumes == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", volumes)
	t.Log("PASS")
}

func TestListBlockVolumeWithNameSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewListBlockVolumesRequest(1, 10).WithName("pvc-24182151-aa4a-4a55-9572-f551c3d003aa")
	volumes, sdkerr := vngcloud.VServerGateway().V2().VolumeService().ListBlockVolumes(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	if volumes == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	if volumes.Len() != 1 {
		t.Fatalf("Expect 1 but got %d", volumes.Len())
	}

	t.Log("Result: ", volumes)
	t.Log("PASS")
}

func TestListBlockVolumeWithFailure(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewListBlockVolumesRequest(1, -10)
	volumes, sdkerr := vngcloud.VServerGateway().V2().VolumeService().ListBlockVolumes(opt)
	if sdkerr == nil {
		t.Fatalf("Expect error but got nil")
	}

	if volumes != nil {
		t.Fatalf("Expect nil but got %v", volumes)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestGetBlockVolumeByIdFailure(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewGetBlockVolumeByIdRequest("this-is-fake")
	volume, sdkerr := vngcloud.VServerGateway().V2().VolumeService().GetBlockVolumeById(opt)
	if sdkerr == nil {
		t.Fatalf("Expect error but got nil")
	}

	if volume != nil {
		t.Fatalf("Expect nil but got %v", volume)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestGetBlockVolumeByIdSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewGetBlockVolumeByIdRequest("vol-0597c473-5df1-4588-aef6-03a1c2814927")
	volume, sdkerr := vngcloud.VServerGateway().V2().VolumeService().GetBlockVolumeById(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	if volume == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", volume)
	t.Log("PASS")
}
