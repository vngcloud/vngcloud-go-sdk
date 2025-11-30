package test

import (
	ltesting "testing"

	v2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/volume/v2"
)

func TestListSnapshotFailure(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewListSnapshotsByBlockVolumeIdRequest(1, 10, "fsffsfsdfdsfsdf")
	_, sdkerr := vngcloud.VServerGateway().V2().VolumeService().ListSnapshotsByBlockVolumeId(opt)

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestListSnapshotSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewListSnapshotsByBlockVolumeIdRequest(1, 10, "vol-d360fd83-948d-4efa-ab46-aab97328e275")
	snapshots, sdkerr := vngcloud.VServerGateway().V2().VolumeService().ListSnapshotsByBlockVolumeId(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	if snapshots == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", snapshots)
	t.Log("PASS")
}

func TestCreateSnapshotFailure(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewCreateSnapshotByBlockVolumeIdRequest(
		"teasdadasdadst",
		"vol-d360fd83-948d-4efa-ab46-aab97328e275").WithPermanently(true)
	_, sdkerr := vngcloud.VServerGateway().V2().VolumeService().CreateSnapshotByBlockVolumeId(opt)

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestDeleteSnapshot(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewDeleteSnapshotByIdRequest("snap-vol-pt-03e5891b-xxxx-4eb9-b2e6-be599f4e2a4b")
	sdkerr := vngcloud.VServerGateway().V2().VolumeService().DeleteSnapshotById(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	t.Log("PASS")
}
