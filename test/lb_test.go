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

func TestGetLoadBalancerSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewGetLoadBalancerByIdRequest("lb-f7adf4ba-7734-45f3-8cb5-9b0c3850cd6f").
		AddUserAgent("vks-cluster-id/user-1234:cluster-1")
	lb, sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().GetLoadBalancerById(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestGetLoadBalancerFailure(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewGetLoadBalancerByIdRequest("lb-f7adf4ba-7734-45f3-8cb5-9b0c3850cc6f").
		AddUserAgent("vks-cluster-id/user-1234:cluster-1")
	lb, sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().GetLoadBalancerById(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestListLoadBalancer(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewListLoadBalancersRequest(1, 10).WithName("cuongdm3-testlb-empty-members")
	lbs, sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().ListLoadBalancers(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lbs == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lbs)
	t.Log("PASS")
}

func TestCreatePoolWithoutMembersSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewCreatePoolRequest("cuongdm3-test-pool-7", lslbv2.PoolProtocolTCP).
		WithLoadBalancerId("lb-f7adf4ba-7734-45f3-8cb5-9b0c3850cd6f").
		WithHealthMonitor(lslbv2.NewHealthMonitor(lslbv2.HealthCheckProtocolTCP))
	pool, sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().CreatePool(opt)

	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if pool == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", pool)
	t.Log("PASS")
}

func TestCreatePoolWithMembersSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewCreatePoolRequest("cuongdm3-test-pool-3", lslbv2.PoolProtocolTCP).
		WithLoadBalancerId("lb-f7adf4ba-7734-45f3-8cb5-9b0c3850cddf").
		WithMembers(lslbv2.NewMember("cuongdm3-member-1", "10.84.0.32", 80, 80)).
		WithHealthMonitor(lslbv2.NewHealthMonitor(lslbv2.HealthCheckProtocolTCP))

	pool, sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().CreatePool(opt)

	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if pool == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", pool)
	t.Log("PASS")
}

func TestCreateListenerSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewCreateListenerRequest("cuongdm3-test-listener-10", lslbv2.CreateOptsListenerProtocolOptTCP, 8080).
		WithLoadBalancerId("lb-f7adf4ba-7734-45f3-8cb5-9b0c3850cddf")

	listener, sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().CreateListener(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listener == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", listener)
	t.Log("PASS")
}

func TestCreateListenerWithPoolIdSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewCreateListenerRequest("cuongdm3-test-listener-14", lslbv2.CreateOptsListenerProtocolOptTCP, 8087).
		WithLoadBalancerId("lb-f7adf4ba-7734-45f3-8cb5-9b0c3850cd6f").WithDefaultPoolId("pool-82c3c670-6662-4087-bfc1-8098f25e84df")

	listener, sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().CreateListener(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listener == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", listener)
	t.Log("PASS")
}

func TestUpdateListenerSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewUpdateListenerRequest("lb-f7adf4ba-7734-45f3-8cb5-9b0c3850cd6f", "lis-23655c30-e458-49ac-ba55-49dfcd104db8").
		WithTimeoutClient(100).
		WithTimeoutConnection(100).
		WithTimeoutMember(100).
		WithCidrs("0.0.0.0/0").
		WithDefaultPoolId("pool-b28eb5ce-c714-439d-b140-03ffe6bebf8f")

	sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().UpdateListener(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestListListenersByLoadBalancerId(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewListListenersByLoadBalancerIdRequest("lb-f7adf4ba-7734-45f3-8cb5-9b0c3850cd66")
	listeners, sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().ListListenersByLoadBalancerId(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listeners == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", listeners)
	t.Log("PASS")
}
