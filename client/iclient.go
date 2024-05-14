package client

import ltime "time"

type IClient interface {
	WithHttpClient(httpClient IHttpClient) IClient
}

type IHttpClient interface {
	WithRetryCount(pretryCount int) IHttpClient
	WithDelay(pdelay ltime.Duration) IHttpClient
	WithSleep(psleep ltime.Duration) IHttpClient
}

type IServiceClient interface {
	WithEndpoint(pendpoint string) IServiceClient
	WithName(pname string) IServiceClient
	WithMoreHeaders(pmoreHeaders map[string]string) IServiceClient
	WithKVheaders(pkey string, pvalue string) IServiceClient
}
