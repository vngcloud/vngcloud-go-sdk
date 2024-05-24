package network

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsnetworkSvcV2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/network/v2"
)

func NewNetworkServiceV2(psvcClient lsclient.IServiceClient) INetworkServiceV2 {
	return &lsnetworkSvcV2.NetworkServiceV2{
		VserverClient: psvcClient,
	}
}
