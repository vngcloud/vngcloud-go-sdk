package test

import (
	ltesting "testing"

	lsdnsSvcPortal "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/dns/internal_system/v1"
)

func TestDnsServiceInternal_ListHostedZones(t *ltesting.T) {
	vngcloud := validSdkConfig()
	portalUserId := "53461"

	req := lsdnsSvcPortal.NewListHostedZonesRequest().WithName("")

	resp, sdkErr := vngcloud.VDnsGateway().Internal().DnsService().ListHostedZones(req, portalUserId)
	if sdkErr != nil {
		t.Fatalf("Failed to list hosted zones: %v", sdkErr.GetError())
	}

	if resp == nil {
		t.Fatal("Response should not be nil")
	}

	t.Logf("Successfully listed hosted zones. Count: %d", len(resp.ListData))
	for _, zone := range resp.ListData {
		t.Logf("Hosted Zone: ID=%s, Name=%s, Status=%s", zone.HostedZoneId, zone.DomainName, zone.Status)
	}
}
