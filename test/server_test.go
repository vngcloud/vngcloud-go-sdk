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
