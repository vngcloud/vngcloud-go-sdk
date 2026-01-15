package test

import (
	"net/url"
	ltesting "testing"

	lsnwv1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network/v1"
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
	opt := lsnwv1.NewCreateEndpointRequest(
		"cuongdm3-test",
		"f3d11a4c-f071-4009-88a6-4a21346c8708",
		"net-5ac170fc-834a-4621-b512-481e09b82fc8",
		"sub-0c508dd6-5af6-4f0e-a860-35346b530cf1",
	).WithDescription(
		"This is the service endpoint for vStorage APIs established by the VKS product ")

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

func TestCreateEndpointInternal(t *ltesting.T) {
	vngcloud := validSuperSdkConfig()
	opt := lsnwv1.NewCreateEndpointRequest(
		"tytv2-test",
		"f081902c-25bb-441e-9a53-fd23fc47e3f1",
		"net-dc14bb60-d500-40b5-945f-218540990187",
		"sub-3f7a1d9b-1d68-44d0-a14f-4cc6bf18a7c4",
	).WithDescription(
		"This is the service endpoint for vStorage APIs - established by the VKS product "+
			"Please refrain from DELETING it manually",
	).AddNetworking("HCM03-1B", "sub-3f7a1d9b-1d68-44d0-a14f-4cc6bf18a7c4").
		AddNetworking("HCM03-1A", "sub-85ba01f6-02ec-4dfc-8884-ee0036c68a5b").
		WithScaling(2, 5).
		WithPortalUserId("53461").
		WithEnableDnsName(false).WithBuyMorePoc(false).WithPoc(false).WithEnableAutoRenew(true).
		WithPackageUuid("e060ef9e-a3cf-420e-9cbc-f203b7f413c4")

	lb, sdkerr := vngcloud.VNetworkGateway().V2().NetworkService().CreateEndpoint(opt)
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
	opt := lsnwv1.NewDeleteEndpointByIdRequest(
		"enp-56d7359f-4b9a-4f01-a210-54523c6d0c88",
		"net-5ac170fc-834a-4621-b512-481e09b82fc8",
		"b9ba2b16-389e-48b7-9e75-4c991239da27",
	)

	sdkerr := vngcloud.VNetworkGateway().V1().NetworkService().DeleteEndpointById(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr.GetErrorCode())
	}

	t.Log("PASS")
}

func TestListEndpoints(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lsnwv1.NewListEndpointsRequest(1, 100).WithUuid("enp-9349271b-af44-4e39-8829-615d945fa6c2")

	lb, sdkerr := vngcloud.VNetworkGateway().V1().NetworkService().ListEndpoints(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr.GetErrorCode())
	}

	if lb == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb.At(0))
	t.Log("PASS")
}

func TestEndpoint(t *ltesting.T) {
	raw := `{"page":1,"size":10,"search":[{"field":"vpcId","value":"net-5ac170fc-834a-4621-b512-481e09b82fc8"}]}`
	encode := url.QueryEscape(raw)

	t.Log("Encode: ", encode)
}

func TestListEndpointTags(t *ltesting.T) {
	vngcloud := validSuperSdkHcm03bConfig()
	opt := lsnwv1.NewListTagsByEndpointIdRequest(
		"95174",
		"pro-b2fff1cf-6d72-4643-a8e7-5907bc9e439c",
		"enp-3fe5d1e9-679e-4eb8-ad35-d9ce53243259",
	)

	lb, sdkerr := vngcloud.VNetworkGateway().InternalV1().NetworkService().ListTagsByEndpointId(opt)
	if sdkerr != nil {
		t.Logf("Expect nil but got %+v", sdkerr.GetErrorCode())
	}

	if lb == nil {
		t.Logf("Expect not nil but got nil")
	}

	t.Log("Result: ", lb)
	t.Log("PASS")
}

func TestCreateEndpointTags(t *ltesting.T) {
	vngcloud := validSuperSdkConfig()
	opt := lsnwv1.NewCreateTagsWithEndpointIdRequest(
		"60108",
		"pro-88265bae-d2ef-424b-b8a7-9eeb08aec1f7",
		"enp-7e8e4476-feeb-414c-ac03-3501aae607d0",
	).AddTag("cuongdm3", "test")

	sdkerr := vngcloud.VNetworkGateway().InternalV1().NetworkService().CreateTagsWithEndpointId(opt)
	if sdkerr != nil {
		t.Logf("Expect nil but got %+v", sdkerr.GetErrorCode())
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestDeleteTagByEndpointId(t *ltesting.T) {
	vngcloud := validSuperSdkConfig()
	opt := lsnwv1.NewDeleteTagOfEndpointRequest(
		"60108",
		"pro-88265bae-d2ef-424b-b8a7-9eeb08aec1f7",
		"tag-6ceb41e1-47e9-43f0-94dd-521a1af870ee",
	)

	sdkerr := vngcloud.VNetworkGateway().InternalV1().NetworkService().DeleteTagOfEndpoint(opt)
	if sdkerr != nil {
		t.Logf("Expect nil but got %+v", sdkerr.GetErrorCode())
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}

func TestUpdateEndpointTag(t *ltesting.T) {
	vngcloud := validSuperSdkConfig()
	opt := lsnwv1.NewUpdateTagValueOfEndpointRequest(
		"60108",
		"pro-88265bae-d2ef-424b-b8a7-9eeb08aec1f7",
		"tag-c6d6e343-ed13-4bf1-bf2e-e63a1a5e0eab",
		"cuonghahahah",
	)

	sdkerr := vngcloud.VNetworkGateway().InternalV1().NetworkService().UpdateTagValueOfEndpoint(opt)
	if sdkerr != nil {
		t.Logf("Expect nil but got %+v", sdkerr.GetErrorCode())
	}

	t.Log("Result: ", sdkerr)
	t.Log("PASS")
}
