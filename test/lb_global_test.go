package test

import (
	ltesting "testing"

	"k8s.io/utils/ptr"

	"github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/glb/v1"
)

func TestListGlobalPoolsSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewListGlobalPoolsRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7")
	pools, sdkerr := vngcloud.GLBGateway().V1().GLBService().ListGlobalPools(opt)
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
	member := v1.NewGlobalMemberRequest("p_name", "10.105.0.4", "sub-8aa727dd-9857-472f-8766-ece41282d437", 80, 80, 1, false)
	poolMember := v1.NewGlobalPoolMemberRequest("p_name", "hcm", "net-80a4eb74-c7d9-46b4-9705-ffed0e2bc3c2", 100, v1.GlobalPoolMemberTypePublic)
	poolMember.WithMembers(member)
	opt := v1.NewCreateGlobalPoolRequest("annd2-test-pool-4", v1.GlobalPoolProtocolTCP).
		WithLoadBalancerId("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7").
		WithHealthMonitor(v1.NewGlobalHealthMonitor(v1.GlobalPoolHealthCheckProtocolTCP)).
		WithMembers(poolMember)
	pool, sdkerr := vngcloud.GLBGateway().V1().GLBService().CreateGlobalPool(opt)

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
	member := v1.NewGlobalMemberRequest("p_name", "10.105.0.4", "sub-8aa727dd-9857-472f-8766-ece41282d437", 80, 80, 1, false)
	poolMember := v1.NewGlobalPoolMemberRequest("p_name", "hcm", "net-80a4eb74-c7d9-46b4-9705-ffed0e2bc3c2", 100, v1.GlobalPoolMemberTypePrivate)
	poolMember.WithMembers(member)
	opt := v1.NewCreateGlobalPoolRequest("annd2-test-pool-5", v1.GlobalPoolProtocolTCP).
		WithLoadBalancerId("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7").
		WithHealthMonitor(
			v1.NewGlobalHealthMonitor(v1.GlobalPoolHealthCheckProtocolHTTPs).
				WithHealthCheckMethod(ptr.To(v1.GlobalPoolHealthCheckMethodGET)).
				WithPath(ptr.To("/sfdsaf")).
				WithHttpVersion(ptr.To(v1.GlobalPoolHealthCheckHttpVersionHttp1Minor1)).
				WithSuccessCode(ptr.To("200")).
				WithDomainName(ptr.To("example.com")),
		).
		WithMembers(poolMember)
	pool, sdkerr := vngcloud.GLBGateway().V1().GLBService().CreateGlobalPool(opt)

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
	httpMonitor := v1.NewGlobalHealthMonitor(v1.GlobalPoolHealthCheckProtocolHTTPs).
		WithDomainName(ptr.To("exampleee.com")).
		WithHealthCheckMethod(ptr.To(v1.GlobalPoolHealthCheckMethodPOST)).
		WithPath(ptr.To("/hghjgj")).
		WithHttpVersion(ptr.To(v1.GlobalPoolHealthCheckHttpVersionHttp1Minor1))
	opt := v1.NewUpdateGlobalPoolRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7", "gpool-30c2a387-7912-4be7-8e3b-448ef16548ab").
		WithHealthMonitor(httpMonitor)

	pool, sdkerr := vngcloud.GLBGateway().V1().GLBService().UpdateGlobalPool(opt)

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
	opt := v1.NewDeleteGlobalPoolRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7", "gpool-1ffbe2f4-0bb5-4272-afe9-0dfa8a4365df")
	sdkerr := vngcloud.GLBGateway().V1().GLBService().DeleteGlobalPool(opt)

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
	opt := v1.NewListGlobalPoolMembersRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7", "gpool-0f4ba08b-e09d-4a1c-b953-523179cea006")
	members, sdkerr := vngcloud.GLBGateway().V1().GLBService().ListGlobalPoolMembers(opt)
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
	createAction := v1.NewPatchGlobalPoolCreateBulkActionRequest(
		v1.NewGlobalPoolMemberRequest("patch_name", "hcm", "net-86b7c84a-b3dd-4e6a-b66b-d28f36f3fc5f", 100, v1.GlobalPoolMemberTypePublic).
			WithMembers(
				v1.NewGlobalMemberRequest("patch_name_4", "10.105.0.4", "sub-7ceeed28-2cad-4bcd-9a4a-a0041c6d6304", 80, 80, 1, false),
				v1.NewGlobalMemberRequest("patch_name_3", "10.105.0.3", "sub-a7fceae7-5ab5-4768-993f-8e6465f75050", 80, 80, 1, false),
			),
	)
	updateAction := v1.NewPatchGlobalPoolUpdateBulkActionRequest("gpool-mem-4568b7da-e82b-4417-b991-ac040967c0c1",
		v1.NewUpdateGlobalPoolMemberRequest("", "", "", 100).WithMembers(
			v1.NewGlobalMemberRequest("patch_name_44", "10.105.0.44", "sub-7ceeed28-2cad-4bcd-9a4a-a0041c6d6304", 80, 80, 1, false),
			v1.NewGlobalMemberRequest("patch_name_33", "10.105.0.33", "sub-a7fceae7-5ab5-4768-993f-8e6465f75050", 80, 80, 1, false),
		),
	)
	// deleteAction := v1.NewPatchGlobalPoolDeleteBulkActionRequest("gpool-mem-e4a56d03-baf8-448b-98ab-404219fddede")
	opt := v1.NewPatchGlobalPoolMembersRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7", "gpool-0f4ba08b-e09d-4a1c-b953-523179cea006").
		WithBulkAction(
			createAction,
			// deleteAction,
			updateAction,
		)
	sdkerr := vngcloud.GLBGateway().V1().GLBService().PatchGlobalPoolMembers(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("PASS")
}

func TestListGlobalListenersSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewListGlobalListenersRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7")
	listeners, sdkerr := vngcloud.GLBGateway().V1().GLBService().ListGlobalListeners(opt)
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
	opt := v1.NewCreateGlobalListenerRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7", "annd2-test").
		WithDescription("hihi").
		WithPort(85).
		WithTimeoutClient(50).
		WithTimeoutConnection(5).
		WithTimeoutMember(50).
		WithGlobalPoolId("gpool-7000d491-b441-40a0-af01-8039baa8e346")
	listener, sdkerr := vngcloud.GLBGateway().V1().GLBService().CreateGlobalListener(opt)
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
	opt := v1.NewUpdateGlobalListenerRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7", "glis-7ffc4f19-7218-4d38-8016-e3ad2401e3bd").
		WithTimeoutClient(60).
		WithTimeoutConnection(6).
		WithTimeoutMember(60).
		WithGlobalPoolId("gpool-7000d491-b441-40a0-af01-8039baa8e346")
	listener, sdkerr := vngcloud.GLBGateway().V1().GLBService().UpdateGlobalListener(opt)
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
	opt := v1.NewDeleteGlobalListenerRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7", "glis-7ffc4f19-7218-4d38-8016-e3ad2401e3bd")
	sdkerr := vngcloud.GLBGateway().V1().GLBService().DeleteGlobalListener(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("PASS")
}

// --------------------------------------------------

func TestListGlobalLoadBalancerSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewListGlobalLoadBalancersRequest(0, 10)
	lbs, sdkerr := vngcloud.GLBGateway().V1().GLBService().ListGlobalLoadBalancers(opt)
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
	pool := v1.NewCreateGlobalPoolRequest("annd2-test-pool-5", v1.GlobalPoolProtocolTCP).
		WithLoadBalancerId("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7").
		WithHealthMonitor(
			v1.NewGlobalHealthMonitor(v1.GlobalPoolHealthCheckProtocolHTTPs).
				WithHealthCheckMethod(ptr.To(v1.GlobalPoolHealthCheckMethodGET)).
				WithPath(ptr.To("/sfdsaf")).
				WithHttpVersion(ptr.To(v1.GlobalPoolHealthCheckHttpVersionHttp1Minor1)).
				WithSuccessCode(ptr.To("200")).
				WithDomainName(ptr.To("example.com")),
		).
		WithMembers(
			v1.NewGlobalPoolMemberRequest("p_name", "hcm", "net-80a4eb74-c7d9-46b4-9705-ffed0e2bc3c2", 100, v1.GlobalPoolMemberTypePrivate).
				WithMembers(
					v1.NewGlobalMemberRequest("p_name", "10.105.0.4", "sub-8aa727dd-9857-472f-8766-ece41282d437", 80, 80, 1, false),
				),
		)
	listener := v1.NewCreateGlobalListenerRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7", "annd2-test").
		WithDescription("hihi").
		WithPort(85).
		WithTimeoutClient(50).
		WithTimeoutConnection(5).
		WithTimeoutMember(50).
		WithGlobalPoolId("gpool-7000d491-b441-40a0-af01-8039baa8e346")
	vngcloud := validSdkConfig()
	opt := v1.NewCreateGlobalLoadBalancerRequest("annd2-testtt").
		WithDescription("hihi").
		WithGlobalListener(listener).
		WithGlobalPool(pool).WithPackage("pkg-b02e62ab-a282-4faf-8732-a172ef497a7b")

	lb, sdkerr := vngcloud.GLBGateway().V1().GLBService().CreateGlobalLoadBalancer(opt)
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
	opt := v1.NewDeleteGlobalLoadBalancerRequest("glb-3fd57a7e-7bb3-4152-a329-adba6d779c4a")
	sdkerr := vngcloud.GLBGateway().V1().GLBService().DeleteGlobalLoadBalancer(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("PASS")
}

func TestGetGlobalLoadBalancerSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewGetGlobalLoadBalancerByIdRequest("glb-2e550a10-8a9e-4e0e-9086-80d8297ca3f7")
	lb, sdkerr := vngcloud.GLBGateway().V1().GLBService().GetGlobalLoadBalancerById(opt)
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
