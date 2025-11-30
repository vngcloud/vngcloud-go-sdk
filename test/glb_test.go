package test

import (
	lfmt "fmt"
	ltesting "testing"
	"time"

	v1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/glb/v1"
)

func TestGetGlobalListenerSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewGetGlobalListenerRequest(
		"glb-a9799830-f7ef-40a8-ad05-ba7f81a8bb8d",
		"glis-6e0b0a21-0b52-40b7-90e5-f59a83e577c2",
	)
	listener, sdkerr := vngcloud.GLBGateway().V1().GLBService().GetGlobalListener(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listener == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", listener)
	t.Log("PASS")
}

func TestGetGlobalPoolMemberSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewGetGlobalPoolMemberRequest(
		"glb-a9799830-f7ef-40a8-ad05-ba7f81a8bb8d",
		"gpool-e5de4670-27e6-45cf-bc68-ec3803ed6849",
		"gpool-mem-4b3a819d-a83f-4964-8336-da6cb8edf529",
	)
	poolMember, sdkerr := vngcloud.GLBGateway().V1().GLBService().GetGlobalPoolMember(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if poolMember == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Pool Member: %+v", poolMember)
	if poolMember.Members != nil && len(poolMember.Members.Items) > 0 {
		t.Logf("Members count: %d", len(poolMember.Members.Items))
		for i, member := range poolMember.Members.Items {
			t.Logf("Member[%d]: %+v", i, member)
		}
	} else {
		t.Log("No members found")
	}
	t.Log("PASS")
}

func TestDeleteGlobalPoolMemberSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewDeleteGlobalPoolMemberRequest(
		"glb-a9799830-f7ef-40a8-ad05-ba7f81a8bb8d",
		"gpool-e5de4670-27e6-45cf-bc68-ec3803ed6849",
		"gpool-mem-4b3a819d-a83f-4964-8336-da6cb8edf529",
	)

	sdkerr := vngcloud.GLBGateway().V1().GLBService().DeleteGlobalPoolMember(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("PASS")
}

func TestUpdateGlobalPoolMemberSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewUpdateGlobalPoolMemberRequest(
		"glb-a9799830-f7ef-40a8-ad05-ba7f81a8bb8d",
		"gpool-e5de4670-27e6-45cf-bc68-ec3803ed6849",
		"gpool-mem-4b3a819d-a83f-4964-8336-da6cb8edf529",
		100,
	).WithMembers(
		v1.NewGlobalMemberRequest("updated-member", "10.0.0.9", "sub-e208484a-69cd-4a70-a7dd-f60bbfd4b04d", 80, 80, 1, false),
		v1.NewGlobalMemberRequest("updated-member", "10.0.0.10", "sub-e208484a-69cd-4a70-a7dd-f60bbfd4b04d", 80, 80, 1, false),
	)

	poolMember, sdkerr := vngcloud.GLBGateway().V1().GLBService().UpdateGlobalPoolMember(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if poolMember == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Updated Pool Member: %+v", poolMember)
	t.Log("PASS")
}

func TestListGlobalPackagesSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewListGlobalPackagesRequest()
	packages, sdkerr := vngcloud.GLBGateway().V1().GLBService().ListGlobalPackages(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if packages == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Packages: %+v", packages)
	if len(packages.Items) > 0 {
		t.Logf("Packages count: %d", len(packages.Items))
		for i, pkg := range packages.Items {
			t.Logf("Package[%d]: ID=%s, Name=%s, Description=%s", i, pkg.ID, pkg.Name, pkg.Description)
		}
	} else {
		t.Log("No packages found")
	}
	t.Log("PASS")
}

func TestListGlobalRegionsSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewListGlobalRegionsRequest()
	regions, sdkerr := vngcloud.GLBGateway().V1().GLBService().ListGlobalRegions(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if regions == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Regions: %+v", regions)
	if len(regions.Items) > 0 {
		t.Logf("Regions count: %d", len(regions.Items))
		for i, region := range regions.Items {
			t.Logf("Region[%d]: %+v", i, region)
		}
	} else {
		t.Log("No regions found")
	}
	t.Log("PASS")
}

func TestGetGlobalLoadBalancerUsageHistoriesSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	// Using timestamps from 1 day ago to current moment
	now := time.Now().Unix()
	oneDayAgo := now - 86400 // 86400 seconds = 24 hours
	from := lfmt.Sprintf("%d", oneDayAgo)
	to := lfmt.Sprintf("%d", now)
	usageType := "connection_rate"
	opt := v1.NewGetGlobalLoadBalancerUsageHistoriesRequest("glb-a9799830-f7ef-40a8-ad05-ba7f81a8bb8d", from, to, usageType)
	histories, sdkerr := vngcloud.GLBGateway().V1().GLBService().GetGlobalLoadBalancerUsageHistories(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if histories == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Usage Histories: %+v", histories)
	if len(histories.Items) > 0 {
		t.Logf("Usage histories count: %d", len(histories.Items))
		for i, history := range histories.Items {
			t.Logf("History[%d]: %+v", i, history)
		}
	} else {
		t.Log("No usage histories found")
	}
	t.Log("PASS")
}
