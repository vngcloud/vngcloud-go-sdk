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
	opt := lsnetworkSvcV2.NewSetAddressPairInVirtualSubnetRequest(
		getValueOfEnv("VIRTUAL_SUBNET_ID"),
		getValueOfEnv("NETWORK_INTERFACE_ID"),
		"10.30.1.28/30",
	)
	network, err := vngcloud.VServerGateway().V2().NetworkService().SetAddressPairInVirtualSubnet(opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	if network == nil {
		t.Fatalf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", network)
}

func TestDeleteAddressPair(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsnetworkSvcV2.NewDeleteAddressPairRequest(getValueOfEnv("ADDRESS_PAIR_ID"))
	err := vngcloud.VServerGateway().V2().NetworkService().DeleteAddressPair(opt)

	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	t.Log("PASS")
}

func TestCreateAddressPair(t *ltesting.T) {
	vngcloud := validSdkConfigHanRegion()

	virtualAddressId := "vip-0d2402cf-49e8-43bf-abbe-b707597320e9"
	internalNicId := "net-in-3b076753-6561-4e3e-8a66-e10dc79cab2d"

	opt := lsnetworkSvcV2.NewCreateAddressPairRequest(virtualAddressId, internalNicId).
		WithMode(lsnetworkSvcV2.AddressPairModeActiveActive)

	ap, err := vngcloud.VServerGateway().V2().NetworkService().CreateAddressPair(opt)
	if err != nil {
		t.Fatalf("Expect error to be nil but got %+v", err)
	}

	if ap == nil {
		t.Fatalf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", ap)
}
