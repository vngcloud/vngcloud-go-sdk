package test

import (
	ltesting "testing"

	"github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/dns/v1"
)

func TestDnsServiceV1_GetHostedZoneById(t *ltesting.T) {

	vngcloud := validSdkConfig()
	opt := v1.NewGetHostedZoneByIdRequest(
		"hosted-zone-32a21aa3-99a3-4d03-9045-37aa701fa03a",
	)
	hostedZone, sdkerr := vngcloud.VDnsGateway().V1().DnsService().GetHostedZoneById(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if hostedZone == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", hostedZone)
	t.Log("PASS")
}

func TestDnsServiceV1_ListHostedZonesDefault(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewListHostedZonesRequest()
	listHostedZones, sdkerr := vngcloud.VDnsGateway().V1().DnsService().ListHostedZones(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listHostedZones == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	for _, hostedZone := range listHostedZones.ListData {
		t.Logf("HostedZone: %+v", hostedZone)
	}
	t.Log("PASS")
}

func TestDnsServiceV1_ListHostedZonesWithFilter(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewListHostedZonesRequest().WithName("annd2")
	listHostedZones, sdkerr := vngcloud.VDnsGateway().V1().DnsService().ListHostedZones(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listHostedZones == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	for _, hostedZone := range listHostedZones.ListData {
		t.Logf("HostedZone: %+v", hostedZone)
	}
	t.Log("PASS")
}
