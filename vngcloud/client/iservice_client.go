package client

import (
	lreq "github.com/imroc/req/v3"

	lsdkErr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

type IServiceClient interface {
	WithEndpoint(pendpoint string) IServiceClient
	WithName(pname string) IServiceClient
	WithProjectId(pprojectId string) IServiceClient
	WithMoreHeaders(pmoreHeaders map[string]string) IServiceClient
	WithKVheader(pkey string, pvalue string) IServiceClient
	WithClient(pclient IHttpClient) IServiceClient
	ServiceURL(pparts ...string) string

	Post(purl string, preq IRequest) (*lreq.Response, lsdkErr.ISdkError)
	Get(purl string, preq IRequest) (*lreq.Response, lsdkErr.ISdkError)
}

type ISdkAuthentication interface {
	WithAccessToken(paccessToken string) ISdkAuthentication
	WithExpiresAt(pexpiresAt int64) ISdkAuthentication
	UpdateAuth(pauth ISdkAuthentication)
	NeedReauth() bool
	GetAccessToken() string
	GetExpiresAt() int64
}
