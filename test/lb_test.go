package test

import (
	ltesting "testing"

	lsinter "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/loadbalancer/inter"
	lslbv2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/loadbalancer/v2"
)

func TestCreateInterLoadBalancerSuccess(t *ltesting.T) {
	vngcloud := validSuperSdkConfig()
	opt := lsinter.NewCreateLoadBalancerRequest(
		getValueOfEnv("VNGCLOUD_PORTAL_USER_ID"),
		"cuongdm3-test-intervpc",
		"lbp-96b6b072-aadb-4b58-9d5f-c16ad69d36aa",
		"sub-27a0562d-07f9-4e87-81fd-e0ba9658f156",
		"sub-888d8fd2-3fed-4aaa-a62d-8554c0aff651",
	)
	lb, sdkerr := vngcloud.VLBGateway().Internal().LoadBalancerService().CreateLoadBalancer(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestCreateLoadBalancerSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewCreateLoadBalancerRequest(
		"cuongdm3-testlb-tags",
		"lbp-96b6b072-aadb-4b58-9d5f-c16ad69d36aa",
		"sub-27a0562d-07f9-4e87-81fd-e0ba9658f156").
		WithTags("cuongdm3", "cuongdm33333", "vinhnt8", "vinhnt8888888").
		WithListener(lslbv2.NewCreateListenerRequest("cuongdm3-test-listener", lslbv2.CreateOptsListenerProtocolOptTCP, 80)).
		WithPool(lslbv2.NewCreatePoolRequest("cuongdm3-test-pool", lslbv2.PoolProtocolTCP).
			WithMembers(lslbv2.NewMember("cuongdm3-member-1", "10.84.0.22", 80, 80)).
			WithHealthMonitor(lslbv2.NewHealthMonitor(lslbv2.HealthCheckProtocolTCP)))

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

func TestCreateLoadBalancerEmptyMemberSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewCreateLoadBalancerRequest(
		"cuongdm3-testlb-empty-members",
		"lbp-96b6b072-aadb-4b58-9d5f-c16ad69d36aa",
		"sub-27a0562d-07f9-4e87-81fd-e0ba9658f156").
		WithTags("cuongdm3", "cuongdm33333", "vinhnt8", "vinhnt8888888").
		WithListener(lslbv2.NewCreateListenerRequest("cuongdm3-test-listener", lslbv2.CreateOptsListenerProtocolOptTCP, 80)).
		WithPool(lslbv2.NewCreatePoolRequest("cuongdm3-test-pool", lslbv2.PoolProtocolTCP).
			WithHealthMonitor(lslbv2.NewHealthMonitor(lslbv2.HealthCheckProtocolTCP)))

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

func TestCreateInterVPCLoadBalancerWithPoolAndListenerSuccess(t *ltesting.T) {
	vngcloud := validSuperSdkConfig()
	opt := lsinter.NewCreateLoadBalancerRequest(
		getValueOfEnv("VNGCLOUD_PORTAL_USER_ID"),
		"cuongdm3-test-intervpc",
		"lbp-96b6b072-aadb-4b58-9d5f-c16ad69d36aa",
		"sub-27a0562d-07f9-4e87-81fd-e0ba9658f156",
		"sub-888d8fd2-3fed-4aaa-a62d-8554c0aff651").
		WithTags("cuongdm3", "cuongdm33333", "vinhnt8", "vinhnt8888888").
		WithListener(lsinter.NewCreateListenerRequest("cuongdm3-test-listener", lsinter.CreateOptsListenerProtocolOptTCP, 80)).
		WithPool(lsinter.NewCreatePoolRequest("cuongdm3-test-pool", lsinter.PoolProtocolTCP).
			WithMembers(lsinter.NewMember("cuongdm3-member-1", "10.84.0.22", 80, 80)).
			WithHealthMonitor(lsinter.NewHealthMonitor(lsinter.HealthCheckProtocolTCP)))

	lb, sdkerr := vngcloud.VLBGateway().Internal().LoadBalancerService().CreateLoadBalancer(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}
