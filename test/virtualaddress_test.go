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
		WithDescription("Private DNS endpoint address for VPC cuongdm3 created by vDNS. Please DO NOT delete this address.").
		AddUserAgent(vngcloud.GetUserAgent())
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

func TestDeleteVirtualAddressById(t *ltesting.T) {
	virtualAddressId := "vip-1a17ffb3-28e5-4a7a-a4e0-17af09de28aa"

	vngcloud := validSdkConfigHanRegion()
	opt := lsnetworkSvcV2.NewDeleteVirtualAddressByIdRequest(virtualAddressId).
		AddUserAgent(vngcloud.GetUserAgent())
	err := vngcloud.VServerGateway().V2().NetworkService().DeleteVirtualAddressById(opt)

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	t.Log("PASS")
}


func TestGetVirtualAddessById(t *ltesting.T) {
	virtualAddressId := "vip-0d2402cf-49e8-43bf-abbe-b707597320e9"

	vngcloud := validSdkConfigHanRegion()
	opt := lsnetworkSvcV2.NewGetVirtualAddressByIdRequest(virtualAddressId).
		AddUserAgent(vngcloud.GetUserAgent())
	vaddr, err := vngcloud.VServerGateway().V2().NetworkService().GetVirtualAddressById(opt)

	if err != nil {
		t.Errorf("Expect error to be nil but got %+v", err)
	}

	if vaddr == nil {
		t.Errorf("Expect portal not to be nil but got nil")
	}

	t.Log("RESULT:", vaddr)
	t.Log("PASS")
}