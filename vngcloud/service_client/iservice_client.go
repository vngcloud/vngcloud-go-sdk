package service_client

import (
	lshttp "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/http"
	lsdkErr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

type IServiceClient interface {
	WithEndpoint(pendpoint string) IServiceClient
	WithName(pname string) IServiceClient
	WithMoreHeaders(pmoreHeaders map[string]string) IServiceClient
	WithKVheader(pkey string, pvalue string) IServiceClient
	WithClient(pclient lshttp.IHttpClient) IServiceClient
	ServiceURL(pparts ...string) string

	Post(purl string, preq lshttp.IRequest) lsdkErr.ISdkError
}
