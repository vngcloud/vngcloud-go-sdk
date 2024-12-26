package test

import (
	ltesting "testing"

	"github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/loadbalancer/global"
)

func TestListGlobalPoolsSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := global.NewListGlobalPoolsRequest("glb-d81be06b-5109-46ad-97d9-e97fe1aa7933")
	pools, sdkerr := vngcloud.VLBGateway().Global().LoadBalancerService().ListGlobalPools(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if pools == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", pools)
	for _, pool := range pools.Items {
		t.Logf("Pool: %+v", pool)
		t.Logf("Health: %+v", pool.Health)
	}
	t.Log("PASS")
}

func TestCreateGlobalPoolSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	member := global.NewGlobalMemberRequest("p_name", "10.105.0.4", "sub-8aa727dd-9857-472f-8766-ece41282d437", 80, 80, 1, false)
	poolMember := global.NewGlobalPoolMemberRequest("p_name", "hcm", "net-80a4eb74-c7d9-46b4-9705-ffed0e2bc3c2", 100)
	poolMember.WithMembers(member)
	opt := global.NewCreateGlobalPoolRequest("annd2-test-pool-4", global.GlobalPoolProtocolTCP).
		WithLoadBalancerId("glb-d81be06b-5109-46ad-97d9-e97fe1aa7933").
		WithHealthMonitor(global.NewGlobalHealthMonitor(global.GlobalPoolHealthCheckProtocolTCP)).
		WithMembers(poolMember)
	pool, sdkerr := vngcloud.VLBGateway().Global().LoadBalancerService().CreateGlobalPool(opt)

	if sdkerr != nil {
		t.Log(sdkerr.GetError())
		t.Log(sdkerr.GetErrorCode())
		t.Log(sdkerr.GetMessage())
		t.Log(sdkerr.GetErrorCategories())
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if pool == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", pool)
	t.Log("PASS")
}

func TestCreateGlobalPoolHTTPSSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	member := global.NewGlobalMemberRequest("p_name", "10.105.0.4", "sub-8aa727dd-9857-472f-8766-ece41282d437", 80, 80, 1, false)
	poolMember := global.NewGlobalPoolMemberRequest("p_name", "hcm", "net-80a4eb74-c7d9-46b4-9705-ffed0e2bc3c2", 100)
	poolMember.WithMembers(member)
	opt := global.NewCreateGlobalPoolRequest("annd2-test-pool-5", global.GlobalPoolProtocolTCP).
		WithLoadBalancerId("glb-d81be06b-5109-46ad-97d9-e97fe1aa7933").
		WithHealthMonitor(
			global.NewGlobalHealthMonitor(global.GlobalPoolHealthCheckProtocolHTTPs).
				WithHealthCheckMethod("GET").
				WithPath("/sfdsaf").
				WithHttpVersion("1.1").
				WithSuccessCode("200").
				WithDomainName("example.com"),
		).
		WithMembers(poolMember)
	pool, sdkerr := vngcloud.VLBGateway().Global().LoadBalancerService().CreateGlobalPool(opt)

	if sdkerr != nil {
		t.Log(sdkerr.GetError())
		t.Log(sdkerr.GetErrorCode())
		t.Log(sdkerr.GetMessage())
		t.Log(sdkerr.GetErrorCategories())
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if pool == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", pool)
	t.Log("PASS")
}

func TestUpdateGlobalPoolHTTPSSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	httpMonitor := global.NewGlobalHealthMonitor(global.GlobalPoolHealthCheckProtocolHTTPs).
		WithDomainName("exampleee.com").
		WithHealthCheckMethod("POST").
		WithPath("/hghjgj").
		WithHttpVersion("1.1")
	opt := global.NewUpdateGlobalPoolRequest("glb-d81be06b-5109-46ad-97d9-e97fe1aa7933", "gpool-30c2a387-7912-4be7-8e3b-448ef16548ab").
		WithHealthMonitor(httpMonitor)

	pool, sdkerr := vngcloud.VLBGateway().Global().LoadBalancerService().UpdateGlobalPool(opt)

	if sdkerr != nil {
		t.Log(sdkerr.GetError())
		t.Log(sdkerr.GetErrorCode())
		t.Log(sdkerr.GetMessage())
		t.Log(sdkerr.GetErrorCategories())
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if pool == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", pool)
	t.Log("PASS")
}

func TestDeleteGlobalPoolSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := global.NewDeleteGlobalPoolRequest("glb-d81be06b-5109-46ad-97d9-e97fe1aa7933", "gpool-1ffbe2f4-0bb5-4272-afe9-0dfa8a4365df")
	sdkerr := vngcloud.VLBGateway().Global().LoadBalancerService().DeleteGlobalPool(opt)

	if sdkerr != nil {
		t.Log(sdkerr.GetError())
		t.Log(sdkerr.GetErrorCode())
		t.Log(sdkerr.GetMessage())
		t.Log(sdkerr.GetErrorCategories())
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("PASS")
}

func TestListGlobalPoolMembersSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := global.NewListGlobalPoolMembersRequest("glb-d81be06b-5109-46ad-97d9-e97fe1aa7933", "gpool-0e2f3513-8a17-4fc9-b509-8d951fd25b03")
	members, sdkerr := vngcloud.VLBGateway().Global().LoadBalancerService().ListGlobalPoolMembers(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if members == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", members)
	for _, member := range members.Items {
		t.Logf("Member: %+v", member)
		for _, m := range member.Members.Items {
			t.Logf("  - Member: %+v", m)
		}
	}
	t.Log("PASS")
}

func TestPatchGlobalPoolMemberSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	createAction := global.NewPatchGlobalPoolCreateBulkActionRequest(
		global.NewGlobalPoolMemberRequest("patch_name", "hcm", "net-86b7c84a-b3dd-4e6a-b66b-d28f36f3fc5f", 100).WithMembers(
			global.NewGlobalMemberRequest("patch_name_4", "10.105.0.4", "sub-7ceeed28-2cad-4bcd-9a4a-a0041c6d6304", 80, 80, 1, false),
			global.NewGlobalMemberRequest("patch_name_3", "10.105.0.3", "sub-a7fceae7-5ab5-4768-993f-8e6465f75050", 80, 80, 1, false),
		),
	)
	updateAction := global.NewPatchGlobalPoolUpdateBulkActionRequest("gpool-mem-1ea280e7-6739-4ede-9d6a-7d2e39036b2c",
		global.NewUpdateGlobalPoolMemberRequest(100).WithMembers(
			global.NewGlobalMemberRequest("patch_name_44", "10.105.0.44", "sub-7ceeed28-2cad-4bcd-9a4a-a0041c6d6304", 80, 80, 1, false),
			global.NewGlobalMemberRequest("patch_name_33", "10.105.0.33", "sub-a7fceae7-5ab5-4768-993f-8e6465f75050", 80, 80, 1, false),
		),
	)
	deleteAction := global.NewPatchGlobalPoolDeleteBulkActionRequest("gpool-mem-e4a56d03-baf8-448b-98ab-404219fddede")
	opt := global.NewPatchGlobalPoolMemberRequest("glb-d81be06b-5109-46ad-97d9-e97fe1aa7933", "gpool-b31482ab-e84b-4768-a0dd-dacc4178a451").
		WithBulkAction(
			createAction,
			updateAction,
			deleteAction)
	sdkerr := vngcloud.VLBGateway().Global().LoadBalancerService().PatchGlobalPoolMember(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("PASS")
}

func TestListGlobalListenersSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := global.NewListGlobalListenersRequest("glb-d81be06b-5109-46ad-97d9-e97fe1aa7933")
	listeners, sdkerr := vngcloud.VLBGateway().Global().LoadBalancerService().ListGlobalListeners(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listeners == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", listeners)
	for _, listener := range listeners.Items {
		t.Logf("Listener: %+v", listener)
	}
	t.Log("PASS")
}

func TestCreateGlobalListenerSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := global.NewCreateGlobalListenerRequest("glb-d81be06b-5109-46ad-97d9-e97fe1aa7933", "annd2-test").
		WithDescription("hihi").
		WithPort(85).
		WithTimeoutClient(50).
		WithTimeoutConnection(5).
		WithTimeoutMember(50).
		WithGlobalPoolId("gpool-7000d491-b441-40a0-af01-8039baa8e346")
	listener, sdkerr := vngcloud.VLBGateway().Global().LoadBalancerService().CreateGlobalListener(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listener == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", listener)
	t.Log("PASS")
}

func TestUpdateGlobalListenerSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := global.NewUpdateGlobalListenerRequest("glb-d81be06b-5109-46ad-97d9-e97fe1aa7933", "glis-7ffc4f19-7218-4d38-8016-e3ad2401e3bd").
		WithTimeoutClient(60).
		WithTimeoutConnection(6).
		WithTimeoutMember(60).
		WithGlobalPoolId("gpool-7000d491-b441-40a0-af01-8039baa8e346")
	listener, sdkerr := vngcloud.VLBGateway().Global().LoadBalancerService().UpdateGlobalListener(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listener == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", listener)
	t.Log("PASS")
}

func TestDeleteGlobalListenerSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := global.NewDeleteGlobalListenerRequest("glb-d81be06b-5109-46ad-97d9-e97fe1aa7933", "glis-7ffc4f19-7218-4d38-8016-e3ad2401e3bd")
	sdkerr := vngcloud.VLBGateway().Global().LoadBalancerService().DeleteGlobalListener(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("PASS")
}

// --------------------------------------------------

func TestListGlobalLoadBalancerSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := global.NewListGlobalLoadBalancersRequest(0, 10)
	lbs, sdkerr := vngcloud.VLBGateway().Global().LoadBalancerService().ListGlobalLoadBalancers(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lbs == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", lbs)
	for _, lb := range lbs.Items {
		t.Logf("LB: %+v", lb)
		for _, vip := range lb.Vips {
			t.Logf("  - VIP: %+v", vip)
		}
		for _, domain := range lb.Domains {
			t.Logf("  - Domain: %+v", domain)
		}
	}
	t.Log("PASS")
}

func TestCreateGlobalLoadBalancerSuccess(t *ltesting.T) {
	pool := global.NewCreateGlobalPoolRequest("annd2-test-pool-5", global.GlobalPoolProtocolTCP).
		WithLoadBalancerId("glb-d81be06b-5109-46ad-97d9-e97fe1aa7933").
		WithHealthMonitor(
			global.NewGlobalHealthMonitor(global.GlobalPoolHealthCheckProtocolHTTPs).
				WithHealthCheckMethod("GET").
				WithPath("/sfdsaf").
				WithHttpVersion("1.1").
				WithSuccessCode("200").
				WithDomainName("example.com"),
		).
		WithMembers(
			global.NewGlobalPoolMemberRequest("p_name", "hcm", "net-80a4eb74-c7d9-46b4-9705-ffed0e2bc3c2", 100).
				WithMembers(
					global.NewGlobalMemberRequest("p_name", "10.105.0.4", "sub-8aa727dd-9857-472f-8766-ece41282d437", 80, 80, 1, false),
				),
		)
	listener := global.NewCreateGlobalListenerRequest("glb-d81be06b-5109-46ad-97d9-e97fe1aa7933", "annd2-test").
		WithDescription("hihi").
		WithPort(85).
		WithTimeoutClient(50).
		WithTimeoutConnection(5).
		WithTimeoutMember(50).
		WithGlobalPoolId("gpool-7000d491-b441-40a0-af01-8039baa8e346")
	vngcloud := validSdkConfig()
	opt := global.NewCreateGlobalLoadBalancerRequest("annd2-testtt").
		WithDescription("hihi").
		WithGlobalListener(listener).
		WithGlobalPool(pool)

	lb, sdkerr := vngcloud.VLBGateway().Global().LoadBalancerService().CreateGlobalLoadBalancer(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", lb)
	t.Log("PASS")
}

func TestDeleteGlobalLoadBalancerSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := global.NewDeleteGlobalLoadBalancerRequest("glb-3fd57a7e-7bb3-4152-a329-adba6d779c4a")
	sdkerr := vngcloud.VLBGateway().Global().LoadBalancerService().DeleteGlobalLoadBalancer(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("PASS")
}

func TestGetGlobalLoadBalancerSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := global.NewGetGlobalLoadBalancerByIdRequest("glb-d81be06b-5109-46ad-97d9-e97fe1aa7933")
	lb, sdkerr := vngcloud.VLBGateway().Global().LoadBalancerService().GetGlobalLoadBalancerById(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", lb)
	for _, vip := range lb.Vips {
		t.Logf("  - VIP: %+v", vip)
	}
	for _, domain := range lb.Domains {
		t.Logf("  - Domain: %+v", domain)
	}
	t.Log("PASS")
}
