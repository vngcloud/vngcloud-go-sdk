package client

import (
	lsdkErr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

type IServiceClient interface {
	WithEndpoint(pendpoint string) IServiceClient
	WithName(pname string) IServiceClient
	WithMoreHeaders(pmoreHeaders map[string]string) IServiceClient
	WithKVheader(pkey string, pvalue string) IServiceClient
	WithClient(pclient IHttpClient) IServiceClient
	ServiceURL(pparts ...string) string

	Post(purl string, preq IRequest) lsdkErr.ISdkError
}

type ISdkAuthentication interface {
	WithAccessToken(paccessToken string) ISdkAuthentication
}
