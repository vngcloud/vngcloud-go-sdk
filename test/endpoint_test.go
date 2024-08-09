package test

import (
	lsnwv1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network/v1"
	ltesting "testing"

	"net/url"
)

func TestGetEndpointSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsnwv1.NewGetEndpointByIdRequest("enp-7575cb25-0033-4c26-9145-53cd90d7778c")

	lb, sdkerr := vngcloud.VNetworkGateway().V1().NetworkService().GetEndpointById(opt)
	if sdkerr != nil {
		t.Errorf("Expect nil but got %+v", sdkerr.GetErrorCode())
	}

	if lb == nil {
		t.Errorf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestCreateEndpoint(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsnwv1.NewCreateEndpointRequest("cuongdm3-test", "b9ba2b16-389e-48b7-9e75-4c991239da27", "net-5ac170fc-834a-4621-b512-481e09b82fc8", "sub-0c508dd6-5af6-4f0e-a860-35346b530cf1").WithDescription("This is the inter-vpc-loadbalancer for vstorage service")

	lb, sdkerr := vngcloud.VNetworkGateway().V1().NetworkService().CreateEndpoint(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr.GetError())
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestDeleteEndpoint(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsnwv1.NewDeleteEndpointByIdRequest("enp-56d7359f-4b9a-4f01-a210-54523c6d0c88", "net-5ac170fc-834a-4621-b512-481e09b82fc8", "b9ba2b16-389e-48b7-9e75-4c991239da27")

	sdkerr := vngcloud.VNetworkGateway().V1().NetworkService().DeleteEndpointById(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr.GetErrorCode())
	}

	t.Log("PASS")
}

func TestListEndpoints(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsnwv1.NewListEndpointsRequest(1, 100).WithVpcId("net-9f390556-f3f1-4989-97b0-0aa3fffe93dd")

	lb, sdkerr := vngcloud.VNetworkGateway().V1().NetworkService().ListEndpoints(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr.GetErrorCode())
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb.Items[0].IPv4Address)
	t.Log("PASS")
}

func TestEndpoint(t *ltesting.T) {
	raw := `{"page":1,"size":10,"search":[{"field":"vpcId","value":"net-5ac170fc-834a-4621-b512-481e09b82fc8"}]}`
	encode := url.QueryEscape(raw)

	t.Log("Encode: ", encode)
}
