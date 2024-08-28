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

func TestCreateInterLoadBalancerV2(t *ltesting.T) {
	vngcloud := validSuperSdkConfig()
	opt := lsinter.NewCreateLoadBalancerRequest(
		getValueOfEnv("VNGCLOUD_PORTAL_USER_ID"),
		"phongnt10-test-intervpc",
		"lbp-96b6b072-aadb-4b58-9d5f-c16ad69d36aa",
		"sub-403b36d2-39fc-47c4-b40b-8df0ecb71045",
		"sub-f7770744-6aa4-4292-9ff9-b43b44716ede",
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

func TestCreateInterLoadBalancerTanTm3(t *ltesting.T) {
	vngcloud := validSuperSdkConfig()
	opt := lsinter.NewCreateLoadBalancerRequest(
		getValueOfEnv("VNGCLOUD_PORTAL_USER_ID"),
		"tantm3-ruachen-giatdo-naucom",
		"lbp-96b6b072-aadb-4b58-9d5f-c16ad69d36aa",
		"sub-a5e4f8e9-e99e-498c-b99a-6b4720cf5c6f",
		"sub-0f20f37a-602c-4b17-b5f8-f81d4c36aab1",
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

func TestCreateInterLoadBalancerSuccess2(t *ltesting.T) {
	vngcloud := validVinhNt8SdkConfig()
	opt := lsinter.NewCreateLoadBalancerRequest(
		getValueOfEnv("VINHCLIENT_USER_ID"),
		"lb-overlap-private-2",
		"lbp-96b6b072-aadb-4b58-9d5f-c16ad69d36aa",
		"sub-0f20f37a-602c-4b17-b5f8-f81d4c36aab1",
		"sub-d1c1e1bf-2364-4a6a-b300-7bc25785a634",
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

func TestCreateInterVpcLbHcm3b(t *ltesting.T) {
	vngcloud := validHcm3bSuperSdkConfig()
	opt := lsinter.NewCreateLoadBalancerRequest(
		getValueOfEnv("HCM3B_USER_ID"),
		"duynh7-hcm04-vstorage",
		"lbp-96b6b072-aadb-4b58-9d5f-c16ad69d36aa",
		"sub-0f20f37a-602c-4b17-b5f8-f81d4c36aab1",
		"sub-511ef030-c961-45b5-baac-9d2dadf7e44c",
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

func TestCreateInterLoadBalancerSuccess3(t *ltesting.T) {
	vngcloud := validVinhNt8SdkConfig().WithProjectId("pro-c8e87532-dc1a-421c-8c5e-4604d772829f")
	opt := lsinter.NewCreateLoadBalancerRequest(
		getValueOfEnv("VINHCLIENT_USER_ID"),
		"vinhnt8-15percent-2",
		"lbp-96b6b072-aadb-4b58-9d5f-c16ad69d36aa",
		"sub-0f20f37a-602c-4b17-b5f8-f81d4c36aab1",
		"sub-0725ef54-a32e-404c-96f2-34745239c28d",
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
		WithListener(lslbv2.NewCreateListenerRequest("cuongdm3-test-listener", lslbv2.ListenerProtocolTCP, 80)).
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
		WithListener(lslbv2.NewCreateListenerRequest("cuongdm3-test-listener", lslbv2.ListenerProtocolTCP, 80)).
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
		WithListener(lsinter.NewCreateListenerRequest("cuongdm3-test-listener", lsinter.ListenerProtocolTCP, 80).
			WithAllowedCidrs("0.0.0.0/0").
			WithTimeoutClient(50).
			WithTimeoutMember(50).
			WithTimeoutConnection(5)).
		WithPool(lsinter.NewCreatePoolRequest("cuongdm3-test-pool", lsinter.PoolProtocolTCP).
			WithMembers(lsinter.NewMember("cuongdm3-member-1", "10.84.0.22", 80, 80)).
			WithHealthMonitor(lsinter.NewHealthMonitor(lsinter.HealthCheckProtocolTCP).
				WithHealthCheckMethod(lsinter.HealthCheckMethodGET).
				WithHttpVersion(lsinter.HealthCheckHttpVersionHttp1).
				WithHealthyThreshold(3).
				WithUnhealthyThreshold(3).
				WithTimeout(5).
				WithInterval(30).
				WithHealthCheckPath("/health").
				WithDomainName("vngcloud.com").
				WithSuccessCode("200")))

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
	opt := lslbv2.NewGetLoadBalancerByIdRequest("lb-10689014-4c30-415c-96f7-2293b137854f").
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

func TestListLoadBalancerByTagsSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewListLoadBalancersRequest(1, 10).
		WithTags("vks-owned-cluster", "hcm03a_user-11412_k8s-a3c03d8e-344c-4a1e-98e0-6d9999ac8077")

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
	vngcloud := validSdkHannibalConfig()
	opt := lslbv2.NewCreatePoolRequest("cuongdm3-test-pool-7", lslbv2.PoolProtocolTCP).
		WithLoadBalancerId("lb-fb378dc6-71c5-417a-9466-677c03885d6f").
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
		WithLoadBalancerId("lb-8bd4ea07-ab40-483d-8387-124ed2f2cecb").
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
	opt := lslbv2.NewCreateListenerRequest("cuongdm3-test-listener-10", lslbv2.ListenerProtocolTCP, 8080).
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
	opt := lslbv2.NewCreateListenerRequest("cuongdm3-test-listener-14", lslbv2.ListenerProtocolTCP, 8087).
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

	opt := lslbv2.NewUpdateListenerRequest("lb-ab73cad3-1dd3-4f2c-9c4c-49702133c5c9", "lis-ed226fe5-65d2-4bb1-986e-7814deb3a55b").
		WithTimeoutClient(50).
		WithTimeoutConnection(5).
		WithTimeoutMember(50).
		WithCidrs("0.0.0.0/0").
		WithDefaultPoolId("pool-a9239c24-9289-4641-a16b-2d71883d168b")

	sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().UpdateListener(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestListListenersByLoadBalancerId(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewListListenersByLoadBalancerIdRequest("lb-8bd4ea07-ab40-483d-8387-124ed2f2cecb")
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

func TestListPoolsByLoadBalancerId(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewListPoolsByLoadBalancerIdRequest("lb-8bd4ea07-ab40-483d-8387-124ed2f2cecb")
	pools, sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().ListPoolsByLoadBalancerId(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if pools == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", pools)
	t.Log("PASS")
}

func TestUpdatePoolMembersSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewUpdatePoolMembersRequest("lb-f7adf4ba-7734-45f3-8cb5-9b0c3850cd6f", "pool-82c3c670-6662-4087-bfc1-8098f25e84df").
		WithMembers(lslbv2.NewMember("cuongdm3-member-50", "10.84.0.50", 80, 80))

	sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().UpdatePoolMembers(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestListPoolMembersSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewListPoolMembersRequest("lb-8bd4ea07-ab40-483d-8387-124ed2f2cecb", "pool-528261c5-9fb4-40bb-bd48-f47b79b272f3")
	members, sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().ListPoolMembers(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if members == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", members)
	t.Log("PASS")
}

func TestDeletePoolSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewDeletePoolByIdRequest("lb-f7adf4ba-7734-45f3-8cb5-9b0c3850cd6f", "pool-82c3c670-6662-4087-bfc1-8098f25e84df")
	sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().DeletePoolById(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestDeleteListenterSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewDeleteListenerByIdRequest("lb-f7adf4ba-7734-45f3-8cb5-9b0c3850cd6f", "lis-23655c30-e458-49ac-ba55-49dfcd104db8")
	sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().DeleteListenerById(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestDeleteLoadBalancer(t *ltesting.T) {
	vngcloud := validHcm3bSdkConfig()
	opt := lslbv2.NewDeleteLoadBalancerByIdRequest("lb-50b72305-02a5-4235-8422-42d6638c7845")
	sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().DeleteLoadBalancerById(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestListTagsSuccess(t *ltesting.T) {
	vngcloud := validSuperSdkConfig2()
	opt := lslbv2.NewListTagsRequest("lb-b1153f05-fd44-4861-8b66-d8b811597faf")
	tags, sdkErr := vngcloud.VLBGateway().V2().LoadBalancerService().ListTags(opt)
	if sdkErr != nil {
		t.Fatalf("Expect nil but got %+v", sdkErr)
	}

	if tags == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", tags)
	t.Log("PASS")
}

func TestCreateTagsSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewCreateTagsRequest("lb-0d9bc46b-66db-4c57-8270-cd380226839d").
		WithTags("cuongdm4", "cuongdm4", "vinhnt9", "vinhnt9")
	sdkErr := vngcloud.VLBGateway().V2().LoadBalancerService().CreateTags(opt)
	if sdkErr != nil {
		t.Fatalf("Expect nil but got %+v", sdkErr.GetMessage())
	}

	t.Log("Result: ", sdkErr)
	t.Log("PASS")
}

func TestUpdateTagsSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lslbv2.NewUpdateTagsRequest("lb-0d9bc46b-66db-4c57-8270-cd380226839d").
		WithTags("cuongdm4", "cuongdm5")
	sdkErr := vngcloud.VLBGateway().V2().LoadBalancerService().UpdateTags(opt)
	if sdkErr != nil {
		t.Fatalf("Expect nil but got %+v", sdkErr.GetMessage())
	}

	t.Log("Result: ", sdkErr)
	t.Log("PASS")
}
