package test

import (
	ltesting "testing"

	lsnetworkSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network/v2"
)

func TestGetAllAddressPairByVirtualSubnetId(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsnetworkSvcV2.NewGetAllAddressPairByVirtualSubnetIdRequest(getValueOfEnv("VIRTUAL_SUBNET_ID"))
	network, err := vngcloud.VServerGateway().V2().NetworkService().GetAllAddressPairByVirtualSubnetId(opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	if network == nil {
		t.Fatalf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", network)
	for _, addressPair := range network {
		t.Logf("AddressPair: %+v", addressPair)
	}
	t.Log("PASS")
}

func TestSetAddressPairInVirtualSubnet(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsnetworkSvcV2.NewSetAddressPairInVirtualSubnetRequest(getValueOfEnv("VIRTUAL_SUBNET_ID"), getValueOfEnv("NETWORK_INTERFACE_ID"), "10.30.1.28/30")
	network, err := vngcloud.VServerGateway().V2().NetworkService().SetAddressPairInVirtualSubnet(opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	if network == nil {
		t.Fatalf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", network)
}
