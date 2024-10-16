package test

import (
	lscomputeSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/compute/v2"
	ltesting "testing"
)

func TestCreateServerFailed(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lscomputeSvcV2.NewCreateServerRequest(
		"cuongdm3-test",
		"img-b5bf635e-0456-4765-b493-31d5fcfc05aa",
		"flav-3929c073-9da9-486f-a96f-9282dbb8d83f",
		"net-4f35f173-e0fe-4202-9c2b-5121b558bcd2",
		"sub-1f98ff1e-2e36-4a40-a0f4-4eadfeb1ea63",
		"vtype-61c3fc5b-f4e9-45b4-8957-8aa7b6029018",
		30,
	)
	server, sdkerr := vngcloud.VServerGateway().V2().ComputeService().CreateServer(opt)
	if sdkerr == nil {
		t.Fatalf("Expect error but got nil")
	}

	if server != nil {
		t.Fatalf("Expect nil but got %v", server)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestCreateServerSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lscomputeSvcV2.NewCreateServerRequest(
		"cuongdm3-test-tags",
		"img-b5bf635e-0456-4765-b493-31d5fcfc05aa",
		"flav-3929c073-9da9-486f-a96f-9282dbb8d83f",
		"net-4f35f173-e0fe-4202-9c2b-5121b558bcd3",
		"sub-1f98ff1e-2e36-4a40-a0f4-4eadfeb1ea63",
		"vtype-61c3fc5b-f4e9-45b4-8957-8aa7b6029018",
		30,
	).WithTags("cuongdm3", "deptrai", "wife", "unknown")
	server, sdkerr := vngcloud.VServerGateway().V2().ComputeService().CreateServer(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	if server == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", server)
	t.Log("PASS")
}

func TestGetServerByIdFailure(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lscomputeSvcV2.NewGetServerByIdRequest("server-1")
	server, sdkerr := vngcloud.VServerGateway().V2().ComputeService().GetServerById(opt)
	if sdkerr == nil {
		t.Fatalf("Expect error but got nil")
	}

	if server != nil {
		t.Fatalf("Expect nil but got %v", server)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestGetServerByIdSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lscomputeSvcV2.NewGetServerByIdRequest(getValueOfEnv("SERVER_ID"))
	server, sdkerr := vngcloud.VServerGateway().V2().ComputeService().GetServerById(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	if server == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", server)
	t.Log("PASS")
}

func TestDeleteServerByIdFailure(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lscomputeSvcV2.NewDeleteServerByIdRequest("this-is-fake-id")
	sdkerr := vngcloud.VServerGateway().V2().ComputeService().DeleteServerById(opt)

	if sdkerr == nil {
		t.Fatalf("Expect error but got nil")
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestDeleteServerByIdSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lscomputeSvcV2.NewDeleteServerByIdRequest(getValueOfEnv("DELETE_SERVER_ID"))
	sdkerr := vngcloud.VServerGateway().V2().ComputeService().DeleteServerById(opt)

	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	t.Log("PASS")
}

func TestDeleteServerByIdSuccess2(t *ltesting.T) {
	vngcloud := customerSdkConfig()
	opt := lscomputeSvcV2.NewDeleteServerByIdRequest("ins-b40e3911-c524-47e1-a58f-d12efcb0ad66")
	sdkerr := vngcloud.VServerGateway().V2().ComputeService().DeleteServerById(opt)

	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	t.Log("PASS")
}

func TestUpdateServerSecgroupsByServerIdFailure(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lscomputeSvcV2.NewUpdateServerSecgroupsRequest("this-is-fake-id")
	server, sdkerr := vngcloud.VServerGateway().V2().ComputeService().UpdateServerSecgroupsByServerId(opt)

	if sdkerr == nil {
		t.Fatalf("Expect error but got nil")
	}

	if server != nil {
		t.Fatalf("Expect nil but got %v", server)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestUpdateServerSecgroupsByServerIdFailure2(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lscomputeSvcV2.NewUpdateServerSecgroupsRequest(getValueOfEnv("SERVER_ID"),
		"secg-93c6c259-0ec7-4905-a232-fe5dccb6674c",
		"this-is-fake-secgroup-id")
	server, sdkerr := vngcloud.VServerGateway().V2().ComputeService().UpdateServerSecgroupsByServerId(opt)

	if sdkerr == nil {
		t.Fatalf("Expect error but got nil")
	}

	if server != nil {
		t.Fatalf("Expect nil but got %v", server)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestUpdateServerSecgroupsByServerIdSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lscomputeSvcV2.NewUpdateServerSecgroupsRequest(getValueOfEnv("SERVER_ID"),
		"secg-93c6c259-0ec7-4905-a232-fe5dccb6674c")
	server, sdkerr := vngcloud.VServerGateway().V2().ComputeService().UpdateServerSecgroupsByServerId(opt)

	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	if server == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", server)
	t.Log("PASS")
}

func TestCreateServerWithAutoRenew(t *ltesting.T) {
	vngcloud := validSdkConfigDevops()
	opt := lscomputeSvcV2.NewCreateServerRequest(
		"cuongdm3-dep-trai-vo-dich-sieu-cap-vu-tru-4",
		"img-108b3a77-ab58-4000-9b3e-190d0b4b07fc",
		"flav-3929c073-9da9-486f-a96f-9282dbb8d83f",
		"net-dae83c7a-f837-4227-bcfa-ec0755549724",
		"sub-f7770744-6aa4-4292-9ff9-b43b44716ede",
		"vtype-61c3fc5b-f4e9-45b4-8957-8aa7b6029018",
		30).
		WithTags("cuongdm3", "deptrai", "wife", "unknown").
		WithAutoRenew(false).
		WithType("VKS").WithProduct("VKS").
		WithPoc(false)

	server, sdkerr := vngcloud.VServerGateway().V2().ComputeService().CreateServer(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	if server == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", server)
	t.Log("PASS")
}

func TestAttachVolumeFailure(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lscomputeSvcV2.NewAttachBlockVolumeRequest("this-is-fake-server-id", "vol-a9484a51-243b-4217-81d4-9f55a7ad426d")
	sdkerr := vngcloud.VServerGateway().V2().ComputeService().AttachBlockVolume(opt)

	if sdkerr == nil {
		t.Fatalf("Expect error but got nil")
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestAttachVolumeSuccess(t *ltesting.T) {
	// VOLUME vol-3aced398-5f8e-4040-84aa-7309a5c5b365 is IN-PROCESS" volumeID="vol-3aced398-5f8e-4040-84aa-7309a5c5b365" nodeID="ins-869ad034-60c1-4f39-bb41-fcdf6b3d4bd4"
	vngcloud := validSdkConfig()
	opt := lscomputeSvcV2.NewAttachBlockVolumeRequest("ins-869ad034-60c2-4f39-bb41-fcdf6b3d4bd4", "vol-3aced398-5f8e-4040-84aa-7309a5c5b365")
	sdkerr := vngcloud.VServerGateway().V2().ComputeService().AttachBlockVolume(opt)

	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	t.Log("PASS")
}

func TestDetachVolumeSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lscomputeSvcV2.NewDetachBlockVolumeRequest("ins-c7acd1d3-376c-xxxx-a4bd-27dd87b72a83", "vol-17dc6df0-43d3-4ad2-be88-69ddaef2f146")
	sdkerr := vngcloud.VServerGateway().V2().ComputeService().DetachBlockVolume(opt)

	if sdkerr != nil {
		t.Errorf("Expect error but got nil")
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestDetachVolumeFailure(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lscomputeSvcV2.NewDetachBlockVolumeRequest("ins-6f6b7238-1d57-4c4b-a683-225454a2e168", "this-is-fake-volume-id")
	sdkerr := vngcloud.VServerGateway().V2().ComputeService().DetachBlockVolume(opt)

	if sdkerr == nil {
		t.Fatalf("Expect error but got nil")
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestAttachFloatingIpSuccess(t *ltesting.T) {
	// VOLUME vol-3aced398-5f8e-4040-84aa-7309a5c5b365 is IN-PROCESS" volumeID="vol-3aced398-5f8e-4040-84aa-7309a5c5b365" nodeID="ins-869ad034-60c1-4f39-bb41-fcdf6b3d4bd4"
	vngcloud := validSdkConfig()
	serverId := "ins-1ac74cf8-df72-4ce2-926f-1c5f89ffda9f"
	netId := "net-in-53c08d94-9ab2-414a-8ee9-169d5b36d52c"
	opt := lscomputeSvcV2.NewAttachFloatingIpRequest(serverId, netId)
	sdkerr := vngcloud.VServerGateway().V2().ComputeService().AttachFloatingIp(opt)

	if sdkerr != nil {
		t.Errorf("Expect nil but got %v", sdkerr)
	}

	t.Log("PASS")
}

func TestDetachFloatingIp(t *ltesting.T) {
	vngcloud := validSdkConfig()
	serverId := "ins-1ac74cf8-df72-4ce2-926f-1c5f89ffda9f"
	wanId := "wan-61dcdecc-85d2-4cea-892a-a4a11be52d02"
	netId := "net-in-53c08d95-9ab2-414a-8ee9-169d5b36d52c"
	opt := lscomputeSvcV2.NewDetachFloatingIpRequest(serverId, wanId, netId)
	sdkerr := vngcloud.VServerGateway().V2().ComputeService().DetachFloatingIp(opt)

	if sdkerr != nil {
		t.Errorf("Expect nil but got %v", sdkerr)
	}

	t.Log("PASS")
}
