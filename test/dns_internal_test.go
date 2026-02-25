package test

import (
	ltesting "testing"

	lsdnsSvcPortal "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/dns/internal_system/v1"
)

func TestDnsServiceInternal_ListHostedZonesDefault(t *ltesting.T) {
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
		t.Logf("Hosted Zone: %+v", zone)
	}
}

func TestDnsServiceInternal_ListRecordsDefault(t *ltesting.T) {
	vngcloud := validSdkConfig()
	portalUserId := "53461"
	hostedZoneId := "hosted-zone-5ba95110-e8d1-40f5-92b2-50eefa33bde2"

	req := lsdnsSvcPortal.NewListRecordsRequest(hostedZoneId)

	resp, sdkErr := vngcloud.VDnsGateway().Internal().DnsService().ListRecords(req, portalUserId)
	if sdkErr != nil {
		t.Fatalf("Failed to list record: %v", sdkErr.GetError())
	}

	if resp == nil {
		t.Fatal("Response should not be nil")
	}

	t.Logf("Successfully listed record. Count: %d", len(resp.ListData))
	for _, record := range resp.ListData {
		t.Logf("Record: %+v", record)
	}
}

func TestDnsServiceInternal_DeleteRecord(t *ltesting.T) {
	vngcloud := validSdkConfig()
	portalUserId := "53461"
	hostedZoneId := "hosted-zone-a6acbf48-9f7b-455d-b5be-efa9081722f9"
	recordId := "record-test-id"

	req := lsdnsSvcPortal.NewDeleteRecordRequest(hostedZoneId, recordId)

	sdkErr := vngcloud.VDnsGateway().Internal().DnsService().DeleteRecord(req, portalUserId)
	if sdkErr != nil {
		t.Fatalf("Failed to delete record: %v", sdkErr.GetError())
	}

	t.Logf("Successfully deleted record with ID: %s", recordId)
}

func TestDnsServiceInternal_DeleteHostedZone(t *ltesting.T) {
	vngcloud := validSdkConfig()
	portalUserId := "53461"
	hostedZoneId := "hosted-zone-243d64ba-c955-4aa1-8640-95d3463446d8"

	req := lsdnsSvcPortal.NewDeleteHostedZoneRequest(hostedZoneId)

	sdkErr := vngcloud.VDnsGateway().Internal().DnsService().DeleteHostedZone(req, portalUserId)
	if sdkErr != nil {
		t.Fatalf("Failed to delete hosted zone: %v", sdkErr.GetError())
	}

	t.Logf("Successfully deleted hosted zone with ID: %s", hostedZoneId)
}

func TestDnsServiceInternal_ListHostedZonesByVpc(t *ltesting.T) {
	vngcloud := validSdkConfig()
	portalUserId := "53461"
	targetVpcId := "net-5ed4bdc1-99d9-4d20-aea8-ce4049d9261d" // Replace with actual VPC ID to test

	// List all hosted zones
	req := lsdnsSvcPortal.NewListHostedZonesRequest()
	resp, sdkErr := vngcloud.VDnsGateway().Internal().DnsService().ListHostedZones(req, portalUserId)
	if sdkErr != nil {
		t.Fatalf("Failed to list hosted zones: %v", sdkErr.GetError())
	}

	if resp == nil {
		t.Fatal("Response should not be nil")
	}

	t.Logf("Total hosted zones found: %d", len(resp.ListData))

	// Filter hosted zones that contain the target VPC ID
	var matchingZones []map[string]interface{}
	for _, zone := range resp.ListData {
		for _, vpcId := range zone.AssocVpcIds {
			if vpcId == targetVpcId {
				matchingZones = append(matchingZones, map[string]interface{}{
					"hostedZoneId": zone.HostedZoneId,
					"domainName":   zone.DomainName,
					"type":         zone.Type,
					"status":       zone.Status,
					"assocVpcIds":  zone.AssocVpcIds,
				})
				break
			}
		}
	}

	// Print results
	t.Logf("Found %d hosted zones associated with VPC ID: %s", len(matchingZones), targetVpcId)
	for i, zone := range matchingZones {
		t.Logf("Zone %d:", i+1)
		t.Logf("  - Hosted Zone ID: %s", zone["hostedZoneId"])
		t.Logf("  - Domain Name: %s", zone["domainName"])
		t.Logf("  - Type: %s", zone["type"])
		t.Logf("  - Status: %s", zone["status"])
		t.Logf("  - Associated VPCs: %v", zone["assocVpcIds"])
	}

	if len(matchingZones) == 0 {
		t.Logf("No hosted zones found for VPC ID: %s", targetVpcId)
	}
}

func TestDnsServiceInternal_ListAndDeleteAllRecords(t *ltesting.T) {
	vngcloud := validSdkConfig()
	portalUserId := "53461"
	hostedZoneId := "hosted-zone-243d64ba-c955-4aa1-8640-95d3463446d8"

	// First, list all records in the hosted zone
	listReq := lsdnsSvcPortal.NewListRecordsRequest(hostedZoneId)
	listResp, sdkErr := vngcloud.VDnsGateway().Internal().DnsService().ListRecords(listReq, portalUserId)
	if sdkErr != nil {
		t.Fatalf("Failed to list records: %v", sdkErr.GetError())
	}

	if listResp == nil {
		t.Fatal("List response should not be nil")
	}

	t.Logf("Found %d records in hosted zone %s", len(listResp.ListData), hostedZoneId)

	// Delete each record (skip NS records as they cannot be deleted)
	allowedTypes := map[string]bool{
		"A":     true,
		"CNAME": true,
		"MX":    true,
		"PTR":   true,
		"TXT":   true,
		"SRV":   true,
	}

	for _, record := range listResp.ListData {
		t.Logf("Found record: ID=%s, SubDomain=%s, Type=%s", record.RecordId, record.SubDomain, record.Type)

		// Skip NS records as they cannot be deleted
		if !allowedTypes[record.Type] {
			t.Logf("Skipping record %s of type %s (cannot be deleted)", record.RecordId, record.Type)
			continue
		}

		deleteReq := lsdnsSvcPortal.NewDeleteRecordRequest(hostedZoneId, record.RecordId)
		sdkErr := vngcloud.VDnsGateway().Internal().DnsService().DeleteRecord(deleteReq, portalUserId)
		if sdkErr != nil {
			t.Logf("Failed to delete record %s: %v", record.RecordId, sdkErr.GetError())
			continue
		}

		t.Logf("Successfully deleted record: %s", record.RecordId)
	}

	// req := lsdnsSvcPortal.NewDeleteHostedZoneRequest(hostedZoneId)

	// sdkErr = vngcloud.VDnsGateway().Internal().DnsService().DeleteHostedZone(req, portalUserId)
	// if sdkErr != nil {
	// 	t.Fatalf("Failed to delete hosted zone: %v", sdkErr.GetError())
	// }

	t.Logf("Completed deletion of all records in hosted zone %s", hostedZoneId)
}
