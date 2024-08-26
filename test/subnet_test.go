package test

import (
	lsnetworkSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network/v2"
	ltesting "testing"
)

func TestGetSubnetByIdSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsnetworkSvcV2.NewGetSubnetByIdRequest(getValueOfEnv("NETWORK_ID"), getValueOfEnv("SUBNET_ID"))
	network, err := vngcloud.VServerGateway().V2().NetworkService().GetSubnetById(opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	if network == nil {
		t.Fatalf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", network)
	t.Log("PASS")
}

func TestUpdateSubnetById(t *ltesting.T) {
	vngcloud := validSdkConfig()
	updateBody := lsnetworkSvcV2.UpdateSubnetBody{
		Name: "subnet-1",
		CIDR: "10.30.0.0/24",
		SecondarySubnetRequests: []lsnetworkSvcV2.SecondarySubnetUpdateBody{
			{Name: "subnet3", CIDR: "10.30.6.0/24"},
			{Name: "subnet2", CIDR: "10.30.7.0/24"},
		},
	}

	opt := lsnetworkSvcV2.NewUpdateSubnetByIdRequest(getValueOfEnv("NETWORK_ID"), getValueOfEnv("SUBNET_ID"), &updateBody)
	network, err := vngcloud.VServerGateway().V2().NetworkService().UpdateSubnetById(opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	if network == nil {
		t.Fatalf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", network)
	t.Log("PASS")
}
