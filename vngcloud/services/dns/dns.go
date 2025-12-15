package dns

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsdnsSvcV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/dns/v1"
)

func NewVDnsServiceV1(psvcClient lsclient.IServiceClient) IVDnsServiceV1 {
	return &lsdnsSvcV1.VDnsServiceV1{
		DnsClient: psvcClient,
	}
}
