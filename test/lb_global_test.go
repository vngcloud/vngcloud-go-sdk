package test

import (
	ltesting "testing"

	"github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/loadbalancer/global"
)

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
