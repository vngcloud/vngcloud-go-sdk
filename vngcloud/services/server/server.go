package server

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsserverSvcV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/server/v1"
)

func NewServerServiceInternalV1(psvcClient lsclient.IServiceClient) IServerServiceInternalV1 {
	return &lsserverSvcV1.ServerServiceInternalV1{
		VServerClient: psvcClient,
	}
}
