package test

import (
	ltesting "testing"

	lscomputeSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/compute/v2"
)

func TestCreateServerFailed(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lscomputeSvcV2.NewCreateServerRequest(
		"cuongdm3-test",
		"img-b5bf635e-0456-4765-b493-31d5fcfc05aa",
		"flav-3929c073-9da9-486f-a96f-9282dbb8d83f",
		"vtype-61c3fc5b-f4e9-45b4-8957-8aa7b6029018",
		30,
	).WithNetwork("net-4f35f173-e0fe-4202-9c2b-5121b558bcd2",
		"sub-1f98ff1e-2e36-4a40-a0f4-4eadfeb1ea63")
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
		"vtype-61c3fc5b-f4e9-45b4-8957-8aa7b6029018",
		30,
	).WithTags("cuongdm3", "deptrai", "wife", "unknown").
		WithNetwork("net-4f35f173-e0fe-4202-9c2b-5121b558bcd3",
			"sub-1f98ff1e-2e36-4a40-a0f4-4eadfeb1ea63")
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
		"vtype-61c3fc5b-f4e9-45b4-8957-8aa7b6029018",
		30).
		WithNetwork("net-dae83c7a-f837-4227-bcfa-ec0755549724",
			"sub-f7770744-6aa4-4292-9ff9-b43b44716ede").
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

func TestCreateDnsServer(t *ltesting.T) {
	vngcloud := validSdkConfigHanRegion()
	opt := lscomputeSvcV2.NewCreateServerRequest(
		"cuongdm3-test",
		"img-b5bf635e-0456-4765-b493-31d5fcfc05aa",
		"flav-8066e9ff-5d80-4e8f-aeae-9e8a934bfc44",
		"vtype-7a7a8610-34f5-11ee-be56-0242ac120002",
		30,
	).WithNetwork("", "").WithServerNetworkInterface(
		"pro-5ce9da27-8ac9-40db-8743-d80f6cbf1491",
		"net-6ad5cc2d-5dfe-4632-a578-6446e6503dd0",
		"sub-33ec4719-915f-4818-85b0-f17bdf7f899b",
		true,
	).WithServerNetworkInterface(
		"pro-e5af9dda-cccb-4f49-bb15-de890cb015c7",
		"net-0dc4cf1e-d961-4483-b848-62ed86fa69f1",
		"sub-1a7e3339-5a73-4cd4-a998-77df31e39303",
		false,
	)

	server, sdkerr := vngcloud.VServerGateway().V2().ComputeService().CreateServer(opt)

	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	t.Log("Result: ", server)
	t.Log("PASS")
}

func TestCreateServerVks(t *ltesting.T) {
	vngcloud := validUserSdkConfig()

	opt := lscomputeSvcV2.NewCreateServerRequest(
		"vsadad04f42",
		"img-e9e4240f-6534-47c3-b922-9489fdd30aa3",
		"flav-8066e9ff-5d80-4e8f-aeae-9e8a934bfc44",
		"vtype-7a7a8610-34f5-11ee-be56-0242ac120002",
		140).
		WithAttachFloating(true).
		WithNetwork("net-7c072f8cdb", "sub-07d3bf64684").
		WithPoc(true).
		WithAutoRenew(false).
		WithType("VKS").
		WithProduct("VKS")

	server, sdkerr := vngcloud.VServerGateway().V2().ComputeService().CreateServer(opt)
	if sdkerr != nil {
		t.Logf("Expect nil but got %v", sdkerr)
	}

	t.Logf("Result: %v", server)
	t.Logf("PASS")
}

func TestListServerGroupPolicies(t *ltesting.T) {
	vngcloud := validUserSdkConfigForCuongDm4()
	opt := lscomputeSvcV2.NewListServerGroupPoliciesRequest()
	policies, sdkerr := vngcloud.VServerGateway().V2().ComputeService().ListServerGroupPolicies(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr.GetErrorCode())
	}

	t.Logf("Result: %v", policies.At(0))
}

func TestDeleteServerGroupById(t *ltesting.T) {
	vngcloud := validUserSdkConfigForCuongDm4()
	opt := lscomputeSvcV2.NewDeleteServerGroupByIdRequest("server-group-ac0dc0ca-3ac3-4f52-a942-3269573ec1de")
	sdkerr := vngcloud.VServerGateway().V2().ComputeService().DeleteServerGroupById(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	t.Logf("PASS")
}

func TestListServerGroups(t *ltesting.T) {
	vngcloud := validUserSdkConfigForCuongDm4()
	opt := lscomputeSvcV2.NewListServerGroupsRequest(1, 10)
	groups, sdkerr := vngcloud.VServerGateway().V2().ComputeService().ListServerGroups(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	t.Logf("Result: %v", groups)
}

func TestCreateServerGroup(t *ltesting.T) {
	vngcloud := validUserSdkConfigForCuongDm4()
	opt := lscomputeSvcV2.NewCreateServerGroupRequest("do-not-want-to-talk-more", "you are a idiot guy", "a2162216-cff2-11eb-b8bc-0242ac130003")
	groups, sdkerr := vngcloud.VServerGateway().V2().ComputeService().CreateServerGroup(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %v", sdkerr)
	}

	t.Logf("Result: %v", groups)
}
