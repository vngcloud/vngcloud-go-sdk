package v1

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
)

type ServerServiceInternalV1 struct {
	VServerClient lsclient.IServiceClient
}
