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

func TestDnsServiceV1_ListRecordsDefault(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewListRecordsRequest("hosted-zone-32a21aa3-99a3-4d03-9045-37aa701fa03a")
	listRecords, sdkerr := vngcloud.VDnsGateway().V1().DnsService().ListRecords(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listRecords == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	for _, record := range listRecords.ListData {
		t.Logf("Record: %+v", record)
	}
	t.Log("PASS")
}

func TestDnsServiceV1_ListRecordsWithFilter(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewListRecordsRequest("hosted-zone-32a21aa3-99a3-4d03-9045-37aa701fa03a").WithName("k8s")
	listRecords, sdkerr := vngcloud.VDnsGateway().V1().DnsService().ListRecords(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if listRecords == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	for _, record := range listRecords.ListData {
		t.Logf("Record: %+v", record)
	}
	t.Log("PASS")
}

func TestDnsServiceV1_CreateHostedZone(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewCreateHostedZoneRequest(
		"test-sdk.example.com",
		[]string{"net-dc14bb60-d500-40b5-945f-218540990187"},
		v1.HostedZoneTypePrivate,
	).WithDescription("Test hosted zone created by SDK")

	hostedZone, sdkerr := vngcloud.VDnsGateway().V1().DnsService().CreateHostedZone(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if hostedZone == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Created HostedZone: %+v", hostedZone)
	t.Log("PASS")
}

func TestDnsServiceV1_DeleteHostedZone(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewDeleteHostedZoneRequest("hosted-zone-8d556e58-e84c-4dff-aeda-dc246b296f32")

	sdkerr := vngcloud.VDnsGateway().V1().DnsService().DeleteHostedZone(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("HostedZone deleted successfully")
	t.Log("PASS")
}

func TestDnsServiceV1_UpdateHostedZone(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v1.NewUpdateHostedZoneRequest("hosted-zone-32a21aa3-99a3-4d03-9045-37aa701fa03a").
		WithAssocVpcIds([]string{"net-dc14bb60-d500-40b5-945f-218540990187"}).
		WithDescription("Updated description for hosted zone.")

	sdkerr := vngcloud.VDnsGateway().V1().DnsService().UpdateHostedZone(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	t.Log("HostedZone updated successfully")
	t.Log("PASS")
}
