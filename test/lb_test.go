package test

import (
	lsinter "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/loadbalancer/inter"
	lslbv2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/loadbalancer/v2"
	ltesting "testing"
)

func TestCreateInterLoadBalancerSuccess(t *ltesting.T) {
	vngcloud := validSuperSdkConfig()
	opt := lsinter.NewCreateLoadBalancerRequest(
		"test-adkjhdaskjdh",
		"lbp-96b6b072-aadb-4b58-9d5f-c16ad69d36aa",
		"sub-27a0562d-07f9-4e87-81fd-e0ba9658f156",
		"sub-888d8fd2-3fed-4aaa-a62d-8554c0aff651",
	)
	volume, sdkerr := vngcloud.VLBGateway().Internal().LoadBalancerService().CreateLoadBalancer(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if volume == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", volume)
	t.Log("PASS")
}

func TestCreateLoadBalancerSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewCreateLoadBalancerRequest("cuongdm3-testlb", "lbp-96b6b072-aadb-4b58-9d5f-c16ad69d36aa", "sub-27a0562d-07f9-4e87-81fd-e0ba9658f156")
	lb, sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().CreateLoadBalancer(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}
