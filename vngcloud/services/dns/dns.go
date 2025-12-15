package dns

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsdnsSvcInternal "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/dns/internal_system/v1"
	lsdnsSvcV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/dns/v1"
)

func NewVDnsServiceV1(psvcClient lsclient.IServiceClient) IVDnsServiceV1 {
	return &lsdnsSvcV1.VDnsServiceV1{
		DnsClient: psvcClient,
	}
}

func NewVDnsServiceInternal(psvcClient lsclient.IServiceClient) IVDnsServiceInternal {
	return &lsdnsSvcInternal.VDnsServiceInternal{
		DnsClient: psvcClient,
	}
}
