package test

import (
	lsnetworkSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network/v2"
	ltesting "testing"
)

func TestCreateVirtualAddressCrossProject(t *ltesting.T) {
	virtualAddressName := "cuongdm3-test-virtual-address"
	projectId := "pro-5ce9da27-8ac9-40db-8743-d80f6cbf1491"
	subnetId := "sub-33ec4719-915f-4818-85b0-f17bdf7f899b"

	vngcloud := validSdkConfigHanRegion()
	opt := lsnetworkSvcV2.NewCreateVirtualAddressCrossProjectRequest(
		virtualAddressName, projectId, subnetId).
		WithDescription("Private DNS endpoint address for VPC cuongdm3 created by vDNS. Please DO NOT delete this address.")
	vaddr, err := vngcloud.VServerGateway().V2().NetworkService().CreateVirtualAddressCrossProject(opt)

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if vaddr == nil {
		t.Errorf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", vaddr)
	t.Log("PASS")
}
